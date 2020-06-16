package lexer
// Credit to https://github.com/golang-collections/collections

// these must be modified at metaWrap too
const (
	leftMatMeta  = "["
	rightMatMeta = "]"
	leftFuncMeta = "("
	rightFuncMeta = ")"
	leftIdxMeta = "("
	rightIdxMeta = ")"
	leftPemdas = "("
	rightPemdas = ")"
	semiSep = ";"
	commaSep = ","
)

// pushes meta Item to stack. If there is a matching opposite
// at top of list then it drops meta at top of stack and does not push
func (l *lexer) metaWrap(item ItemType) {
	switch item {
	case ItemLeftIdxMeta, ItemLeftFuncMeta, ItemLeftMatMeta, ItemLeftPemdas:
		l.metaStack.Push(item)
	case ItemRightIdxMeta:
		if l.metaStack.Pop() != ItemLeftIdxMeta {
			panic("I found mismatched right index meta")
		}
	case ItemRightFuncMeta:
		if l.metaStack.Pop() != ItemLeftFuncMeta {
			panic("I found mismatched right function meta")
		}
	case ItemRightMatMeta:
		if l.metaStack.Pop() != ItemLeftMatMeta {
			panic("I found mismatched right matrix meta")
		}
	case ItemRightPemdas:
		if l.metaStack.Pop() != ItemLeftPemdas {
			panic("I found mismatched right group meta (pemdas)")
		}
	}
}

// returns meta Item at top of stack without modifying stack
func (l *lexer) metaCurrent() (it ItemType) {
	I := l.metaStack.Peek()
	switch I.(type) {
	case ItemType:
		return I.(ItemType)
	case nil:
		return ItemNil
	}
	panic("I should have gotten ItemType in stack!")
}

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}
)
// Create a new stack
func New() *Stack {
	return &Stack{nil,0}
}
// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}
// View the top Item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}
// Pop the top Item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}
// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value,this.top}
	this.top = n
	this.length++
}


