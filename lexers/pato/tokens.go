package pato

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

	TokIDENT // <identifier>
	TokEOF   // EOF

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
	// If you get collisions on adding a keyword you'll need to
	// process more bytes of the identifier since this'll indicate
	// two keywords share the same first two bytes.
	// Best course of action is incrementing keyword map size or tuning the hash operations.
	return (uint(id[0])<<4 ^ uint(id[1]) + uint(len(id))) & uint(len(keywordMap)-1)
}

func init() {
	// populate keywordMap
	for tok := keywordBeg + 1; tok < keywordEnd; tok++ {
		h := kwhash(tok.String())
		if keywordMap[h] != 0 {
			panic("imperfect hash")
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
