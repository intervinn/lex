package predicate

func IsToken(r rune) bool {
	switch r {

	case '+', '-', '=', '/',
		'*', '<', '>':
		return true
	default:
		return false
	}
}

func IsEndLiteral(r rune) bool {
	return IsToken(r) || r == '\n' || r == ' '
}

func IsEndToken(r rune) bool {
	return r == ' ' || r == '\n'
}
