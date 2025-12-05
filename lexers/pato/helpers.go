package pato

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.advance()
	}
}

func isIdentifierChar(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isDigitOrDecimal(ch rune) bool {
	return ch == '.' || isDigit(ch)
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}

func isPrintableASCII(ch rune) bool {
	return ch >= 32 && ch <= 126
}
