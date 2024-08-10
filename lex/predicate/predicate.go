package predicate

func IsToken(r rune) bool {
	switch r {
	case '=':
	case '-':
	case '+':
	case '/':
	case '*':
	case '>':
	case '<':
		return true
	default:
		return false
	}
	return false
}

func IsEnd(r rune) bool {
	return IsToken(r) || r == '\n'
}

func IsEndLiteral(r rune) bool {
	return IsEnd(r) || r == ' '
}
