package pato

import (
	"bufio"
	"errors"
	"io"
	"unicode/utf8"
)

type Token uint

const (
	TokUndefined Token = iota
	TokIllegal
	TokLParen
	TokRParen
	TokNewline
	TokIDENT
	TokEOF
)

type Lexer struct {
	input  bufio.Reader
	ch     rune    // current character.
	peek   [1]rune // peek characters (utf8)
	idbuf  []byte  // stores current identifier buildup.
	err    error
	source string
	// positional indices.
	line int
	col  int
	pos  int

	ReuseLiteralBuffer bool
}

// LineCol returns the current line and column position in the source.
func (l *Lexer) LineCol() LineCol {
	return LineCol{Source: l.source, Line: l.line, Col: l.col}
}

// Pos returns the current byte offset in the source.
func (l *Lexer) Pos() Pos { return Pos(l.pos) }

// Err returns the lexer error, or nil if the error is EOF.
func (l *Lexer) Err() error {
	if l.err == io.EOF {
		return nil
	}
	return l.err
}

// IsDone returns true if the lexer has finished processing (error occurred and no current character).
func (l *Lexer) IsDone() bool {
	return l.err != nil && l.ch == 0
}

// Reset initializes the lexer with a new source name and reader.
// It preserves ReuseLiteralBuffer and internal buffers across resets.
func (l *Lexer) Reset(source string, r io.Reader) error {
	if r == nil {
		return errors.New("nil reader")
	} else if source == "" {
		return errors.New("no source name")
	}
	*l = Lexer{
		ReuseLiteralBuffer: l.ReuseLiteralBuffer,
		input:              l.input,
		line:               1,
		idbuf:              l.idbuf,
		source:             source,
	}
	l.input.Reset(r)
	if l.idbuf == nil {
		l.idbuf = make([]byte, 0, 1024)
	}
	// Fill up peek and current character.
	const buflen = len(l.peek)
	l.col = -buflen + 1 // col is 1 based.
	l.pos = -buflen
	for range len(l.peek) {
		l.advance() // fill peek buffer.
	}
	l.advance() // fill ch character.
	return l.err
}

// NextToken returns the next token, its starting byte position, and its literal value.
// Returns TokEOF at end of input, TokIllegal on errors.
func (l *Lexer) NextToken() (tok Token, start Pos, literal []byte) {
	if l.source == "" {
		l.err = errors.New("lexer uninitialized")
		return TokIllegal, 0, nil
	}
	l.skipWhitespace() // We skip early, not after tokenizing. This leads to more intuitive lexer behaviour.
	start = l.Pos()
	tok = LookupSingleChar(l.ch)
	if tok == TokIllegal {
		if l.err == io.EOF {
			tok = TokEOF
		}
		return tok, start, nil
	}
	if tok != TokIDENT {
		// Single character case.
		literal = utf8.AppendRune(l.idbuf[l.bufstart():], l.ch)
		l.advance()
		return tok, start, literal
	}
	// We have an identifier in our hands.
	literal = l.readIdentifier()
	tok = Lookup(string(literal)) // Should be optimized by compiler to not allocate.
	return tok, start, literal
}

func (l *Lexer) readIdentifier() []byte {
	start := l.bufstart()
	for isIdentifierChar(l.ch) || isDigit(l.ch) {
		l.idbuf = utf8.AppendRune(l.idbuf, l.ch)
		l.advance()
	}
	return l.idbuf[start:]
}

func (l *Lexer) advance() {
	// Advance character buffer first, so even on EOF we don't lose the last char
	currentIsNewline := l.ch == '\n'
	l.ch = l.peek[0]
	for i := range len(l.peek) - 1 {
		l.peek[i] = l.peek[i+1]
	}
	ch, sz, err := l.input.ReadRune()
	if err != nil {
		l.peek[len(l.peek)-1] = 0
		l.err = err
		return
	}
	l.col++
	l.pos += sz
	l.peek[len(l.peek)-1] = ch
	if currentIsNewline {
		l.line++
		l.col = 1
	}
}

func (l *Lexer) bufstart() int {
	if l.ReuseLiteralBuffer {
		l.idbuf = l.idbuf[:0]
		return 0
	}
	return len(l.idbuf)
}
