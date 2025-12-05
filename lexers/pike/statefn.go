package pike

import (
	"strings"
)

const idRuneSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

type stateFn func(*lexer) stateFn

var funcNames map[string]struct{}

// This is the initial state and base state
func lexStart(l *lexer) stateFn {
	var eofFound bool
	for !eofFound {
		if strings.HasPrefix(l.input[l.pos:], leftMatMeta) {
			l.emitJunk() // Emit whatever came before matrix if anything at all
			return lexLeftMatMeta
		} else if strings.HasPrefix(l.input[l.pos:], rightMatMeta) {
			l.emitJunk()
			return lexRightMatMeta
		} else if strings.HasPrefix(l.input[l.pos:], ")") {
			return lexClosingMeta
		}
		switch r := l.peek(); {
		case r == eof:
			return lexEOF
		case isNumeric(r):
			return lexNumber
		case isASCIIAlpha(r):
			return lexAlpha
		case isOperator(r):
			return lexOperator
		case isSeparator(r):
			return lexSeparator
		case r == '(':
			return lexLeftPemdas
		}
		l.next()
	}
	panic("unreachable") // end of function on EOF. See lexEOF
}

func lexAlpha(l *lexer) stateFn {
	l.acceptRun(idRuneSet)
	idType := l.getIDType(l.input[l.start:l.pos])
	switch idType {
	case idUndefined:
		l.errorf("I found an undefined identifier '%s'", l.input[l.start:l.pos])
	case idFunc:
		l.emit(ItemFunc)
		return lexLeftFuncMeta
	case idVar:
		l.emit(ItemVar)
		l.accept(leftIdxMeta)
		if l.pos-l.start == len(leftIdxMeta) {
			l.metaWrap(ItemLeftIdxMeta)
			l.emit(ItemLeftIdxMeta)
		}
		return lexStart
	}
	return l.errorf("Unhandled id type for identifier %s", l.input[l.start:l.pos])
}

func lexLeftFuncMeta(l *lexer) stateFn {
	if strings.HasPrefix(l.input[l.pos:], leftFuncMeta) {
		l.pos += len(leftFuncMeta)
		l.metaWrap(ItemLeftFuncMeta)
		l.emit(ItemLeftFuncMeta)
		return lexStart
	}
	return l.errorf("I looked for a function opening meta and couldn't find one")
}

func lexNumber(l *lexer) stateFn {
	// Lex number, decimal, float, imaginary
	l.accept("+-")
	digits := "0123456789" // only decimal
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}
	if l.accept("eE") {
		l.accept("+-")
		floatOK := l.acceptRun("0123456789")
		if !floatOK {
			return l.errorf("I couldn't find all I needed for a float value.")
		}
	}
	l.accept("i") // if imaginary
	if isASCIIAlphaNumeric(l.peek()) {
		l.next()
		return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
	}
	l.emit(ItemNumber)
	return lexStart
}

func lexLeftMatMeta(l *lexer) stateFn {
	l.pos += len(leftMatMeta)
	l.emit(ItemLeftMatMeta)
	return lexStart // Now inside [ ].
}
func lexRightMatMeta(l *lexer) stateFn {
	l.pos += len(rightMatMeta)
	l.emit(ItemRightMatMeta)
	return lexStart // exiting [ ].
}

func lexOperator(l *lexer) stateFn {
	if l.accept("+-*/^") {
		l.emit(ItemOperator)
	} else if l.accept(":") {
		l.emit(ItemColonOp)
	} else {
		return l.errorf("I could not find operator!")
	}
	return lexStart
}

func lexSeparator(l *lexer) stateFn {
	if l.accept(",") {
		l.emit(ItemCommaSep)
	} else if l.accept(";") {
		l.emit(ItemSemiSep)
	} else {
		return l.errorf("I could not find separator!")
	}
	return lexStart
}

func lexClosingMeta(l *lexer) stateFn {
	l.pos += len(")")
	switch currentMeta := l.metaCurrent(); {
	case currentMeta == ItemLeftIdxMeta:
		l.metaWrap(ItemRightIdxMeta)
		l.emit(ItemRightIdxMeta)
	case currentMeta == ItemLeftFuncMeta:
		l.metaWrap(ItemRightFuncMeta)
		l.emit(ItemRightFuncMeta)
	case currentMeta == ItemLeftPemdas:
		l.metaWrap(ItemRightPemdas)
		l.emit(ItemRightPemdas)
	}
	return lexStart
}

func lexLeftPemdas(l *lexer) stateFn {
	if l.accept("(") {
		l.emit(ItemLeftPemdas)
		l.metaWrap(ItemLeftPemdas)
	} else {
		return l.errorf("I was expecting to find left group (pemdas)")
	}
	return lexStart
}

func lexEOF(l *lexer) stateFn {
	if l.pos > l.start {
		l.emit(ItemText)
	}
	l.emit(ItemEOF)
	return nil
}
