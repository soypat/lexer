package pike

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type ItemType int

// tokens in code
const (
	ItemError ItemType = iota // error occurred, value is error text
	ItemNil
	ItemEOF
	ItemNumber
	ItemVar
	ItemOperator
	ItemFunc
	ItemCommaSep
	ItemColonOp
	ItemSemiSep
	ItemLeftFuncMeta
	ItemRightFuncMeta
	ItemLeftMatMeta
	ItemRightMatMeta
	ItemLeftIdxMeta
	ItemRightIdxMeta
	ItemLeftPemdas
	ItemRightPemdas
	ItemIdentifier
	ItemText // plain text
	// unused
	itemEnd
	itemVarIdx
	itemString // quoted string
	itemAnon   // Anonymous function identifier
	itemIf
	itemElse
)

type Item struct {
	typ ItemType // such as ItemNumber
	val string
}

// Identifier types. Variables and functions (keywords?)
type idType int

const (
	idError idType = iota
	idUndefined
	idFunc
	idVar
)

type identifier struct {
	typ idType // such as idFunc
	val string
}

const eof = -1

type lexer struct {
	name        string    // used only for error reports
	input       string    // the string being scanned
	start       int       // start pos of this Item
	pos         int       // current pos in input
	width       int       // width of last rune read
	items       chan Item // channel of scanned Item
	state       stateFn
	identifiers map[string]identifier
	metaStack   Stack
}

// Creates new lexer with a name for error formatting
// and gives it an input string to lex
func NewStringLexer(name, input string) *lexer {
	return lex(name, input)
}

// run lexes the input by executing state functions until
// the state is nil.
func (l *lexer) Run() {
	for state := lexStart; state != nil; {
		state = state(l)
	}
	close(l.items) // No more tokens will be delivered.
}

func (l *lexer) ItemChannel() (items <-chan Item) {
	return l.items
}

// Creates new variable identifier for lexer
// if variable already exists throws error
func (l *lexer) NewVariableID(value string) error {
	if l.getIDType(value) != idUndefined {
		return fmt.Errorf("Variable already exists")
	}
	l.idAdd(identifier{typ: idVar, val: value})
	return nil
}

// Creates new function identifier for lexer
// if function already exists throws error
func (l *lexer) NewFunctionID(value string) error {
	if l.getIDType(value) != idUndefined {
		return fmt.Errorf("Function already exists")
	}
	l.idAdd(identifier{typ: idFunc, val: value})
	return nil
}

// returns Item type enum
func (i *Item) Type() ItemType {
	return i.typ
}

// returns item's value as read in input
func (i *Item) Value() string {
	return i.val
}

// lex creates a new scanner for the input string.
func lex(name, input string) *lexer {
	l := &lexer{
		name:        name,
		input:       input,
		state:       lexStart,
		items:       make(chan Item, 2), // Two items sufficient.
		identifiers: make(map[string]identifier),
	}
	return l
}

// nextItem returns the next Item from the input.
// this for hiding the goroutine
// actually does not use goroutine but does use channel
func (l *lexer) nextItem() Item {
	for {
		select {
		case item := <-l.items:
			return item
		default:
			l.state = l.state(l)
		}
	}
	panic("not reached")
}

// emit passes an Item back to the client.
func (l *lexer) emit(t ItemType) {
	l.items <- Item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

// advances cursor for next rune's width
func (l *lexer) next() (rune rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	rune, l.width =
		utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return rune
}

// terminates lexer and returns a formatted error message to lexer.items
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	msg := fmt.Sprintf(format, args...)
	start := l.pos - 10
	if start < 0 {
		start = 0
	}
	l.items <- Item{
		ItemError,
		fmt.Sprintf("Error at char %d: '%s'\n%s", l.pos, l.input[start:l.pos+1], msg),
	}
	//panic("PANIC")
	return nil
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune.
// Can be called only once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// peek returns but does not consume
// the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// accept consumes the next rune
// if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) bool {
	var accepted bool
	for strings.IndexRune(valid, l.next()) >= 0 {
		accepted = true
	}
	l.backup()
	return accepted
}

// adds new identifier to list of known identifiers.
// returns false if identifier type does not match
// with existing identifier's type
func (l *lexer) idAdd(id identifier) bool {
	currentID, present := l.identifiers[id.val]
	if present && currentID.typ != id.typ {
		return false
	}
	l.identifiers[id.val] = id
	return true
}

func (l *lexer) getIDType(val string) idType {
	id, present := l.identifiers[val]
	if !present {
		return idUndefined
	}
	return id.typ
}

func (l *lexer) emitJunk() bool {
	if l.pos > l.start { //is token empty?
		l.emit(ItemText) // emit whatever came before
		return true
	}
	return false
}
