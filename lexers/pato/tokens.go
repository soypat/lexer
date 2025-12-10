package pato

import "fmt"

//go:generate stringer -type=Token -linecomment -output stringers.go .

type Token uint

const (
	TokUndefined Token = iota // undefined
	TokIllegal                // illegal
	// Single character tokens:
	TokNewline  // \n
	TokLParen   // (
	TokRParen   // )
	TokLBrace   // {
	TokRBrace   // }
	TokLBracket // [
	TokRBracket // ]
	TokPlus     // +
	TokMinus    // -
	TokAsterisk // *
	TokSlash    // /

	TokIntLit // <integer literal>
	TokIDENT  // <identifier>
	TokEOF    // EOF

	// Add keywords between keywordBeg and keywordEnd.
	keywordBeg
	TokIf   // if
	TokElse // else
	TokFor  // for
	keywordEnd
)

var keywordMap [1 << 4]Token

// kwhash is a perfect hash function for keywords.
// It assumes that s has at least length 2.
func kwhash(id string) uint {
	// See perfect_hash_test.go for information on how to search for a
	// perfect hash function in cases where bit shifts add too little entropy
	// and multiplication is needed.
	return (uint(id[0])<<4 ^ uint(id[1]) + uint(len(id))) & uint(len(keywordMap)-1)
}

func init() {
	// populate keywordMap
	for tok := keywordBeg + 1; tok < keywordEnd; tok++ {
		h := kwhash(tok.String())
		if keywordMap[h] != 0 {
			panic(fmt.Sprintf("imperfect hash at %0x %s collides with %s (%d/%d ok)", h, keywordMap[h].String(), tok.String(), tok-keywordBeg-1, keywordEnd-keywordBeg-1))
		}
		keywordMap[h] = tok
	}
}

func IsKeyword(s string) bool {
	if len(s) < 2 {
		return false
	}
	tok := keywordMap[kwhash(s)]
	return tok != 0 && s == tok.String()
}

func Lookup(s string) Token {
	if len(s) < 2 {
		return TokIDENT
	}
	tok := keywordMap[kwhash(s)]
	if tok != 0 && s == tok.String() {
		return tok
	}
	return TokIDENT
}

func LookupSingleChar(r rune) (tok Token) {
	switch r {
	// Single-character token switch branch.
	case 0:
		tok = TokIllegal
	case '\n':
		tok = TokNewline
	case '(':
		tok = TokLParen
	case ')':
		tok = TokRParen
	case '{':
		tok = TokLBrace
	case '}':
		tok = TokRBrace
	case '[':
		tok = TokLBracket
	case ']':
		tok = TokRBracket
	case '+':
		tok = TokPlus
	case '-':
		tok = TokMinus
	case '/':
		tok = TokSlash
	case '*':
		tok = TokAsterisk
	default:
		tok = TokIDENT
	}
	return tok
}
