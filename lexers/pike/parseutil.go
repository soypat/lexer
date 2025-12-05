package pike

import "unicode"

func isSpace(r rune) bool {
	return unicode.IsSpace(r)
}

func isNumeric(r rune) bool {
	return r >= '0' && r <= '9'
}

func isNumericalStartToken(r rune) bool {
	return r == '+' || r == '-' || isNumeric(r)
}

func isOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/' || r == '^' || r == ':'
}

func isSeparator(r rune) bool {
	return r == ',' || r == ';'
}

func isNumericalToken(r rune) bool {
	return r == '.' || r == 'e' || r == 'E' || r == '+' || r == '-' || r == 'i'
}

func isASCIIAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isASCIIAlphaNumeric(r rune) bool {
	return isNumeric(r) && isASCIIAlpha(r)
}

func isFunction(s string) bool {
	_, present := funcNames[s]
	return present
}
