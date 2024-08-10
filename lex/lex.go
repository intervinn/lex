package lex

import (
	"fmt"
	"unicode"

	"github.com/intervinn/mx/kind"
	"github.com/intervinn/mx/lex/predicate"
	"github.com/intervinn/mx/node"
	"github.com/intervinn/mx/node/literal"
	"github.com/intervinn/mx/node/token"
)

func Type[T any](v any) (T, bool) {
	res, ok := v.(T)
	return res, ok
}

func Line(line string) []node.BaseNodeInterface {
	i := 0
	c := '0'
	eof := len(line) - 1
	pos := 0
	tmp := ""
	captured := []node.BaseNodeInterface{}
	cmode := -1 // which kind do we capture

	Next := func() rune {
		return rune(line[i+1])
	}

	Clear := func() {
		captured = captured[:0]
	}

	CanCapture := func(k int) bool {
		return cmode == k || cmode == -1
	}

	Release := func() {
		cmode = -1
		tmp = ""
	}

	Capture := func(k int) {
		cmode = k
	}

	for i, c = range line {

		// INT LITERAL
		if unicode.IsDigit(c) && CanCapture(kind.IntLiteral) {
			Capture(kind.IntLiteral)
			if tmp == "" {
				pos = i
			}

			tmp += string(c)

			if i == eof || predicate.IsEndLiteral(Next()) {
				captured = append(
					captured,
					literal.NewIntLiteral(pos, i, tmp),
				)
				Release()
			}
		}

		if predicate.IsToken(c) {
			if tmp == "" {
				pos = i
			}
			tmp += string(c)

			if i == eof || predicate.IsEnd(Next()) {
				tk, err := token.NewToken(pos, i, tmp)

				if err != nil {
					panic(err)
				}

				captured = append(
					captured,
					tk,
				)
				Release()
			}
		}

		// END OF LINE
		if i == eof {
			fmt.Printf("%#v\n", captured)
			Clear()
		}
	}
	return captured
}
