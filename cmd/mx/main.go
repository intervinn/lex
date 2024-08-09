package main

import (
	"fmt"
	"unicode"

	"github.com/intervinn/mx/node"
	"github.com/intervinn/mx/node/literal"
)

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

func Type[T any](v any) (T, bool) {
	res, ok := v.(T)
	return res, ok
}

func main() {
	line := "5 + 2"
	i := 0
	c := '0'

	eof := len(line) - 1
	tmp := ""
	pos := 0
	captured := []node.BaseNodeInterface{}

	/*
		s := &node.SourceFile{
			Statements: []node.BaseNodeInterface{},
		}*/

	Next := func() rune {
		return rune(line[i+1])
	}

	Clear := func() {
		captured = captured[:0]
	}

	for i, c = range line {

		// INT LITERAL
		if unicode.IsDigit(c) {
			if tmp == "" {
				pos = i
			}

			tmp += string(c)

			if i == eof || IsEndLiteral(Next()) {
				captured = append(
					captured,
					literal.NewIntLiteral(pos, i, tmp),
				)

				tmp = ""
			}
		}

		// END OF LINE
		if i == eof {
			fmt.Printf("%#v\n", captured)
			Clear()
		}

	}

}
