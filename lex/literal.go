package lex

import (
	"unicode"

	"github.com/intervinn/mx/kind"
	"github.com/intervinn/mx/lex/predicate"
	"github.com/intervinn/mx/node/literal"
)

func analyzeLiteral() {
	analyzeIntLiteral()
}

func analyzeIntLiteral() {
	if unicode.IsDigit(char) && canCapture(kind.IntLiteral) {
		capture(kind.IntLiteral)
		if tmp == "" {
			pos = index
		}

		tmp += string(char)
		if index == eof || predicate.IsEndLiteral(next()) {
			captured = append(
				captured,
				literal.NewIntLiteral(pos, index, tmp),
			)
			release()
		}
	}
}
