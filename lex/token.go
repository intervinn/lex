package lex

import (
	"github.com/intervinn/mx/lex/predicate"
	"github.com/intervinn/mx/node/token"
)

func analyzeToken() {
	if predicate.IsToken(char) {
		if tmp == "" {
			pos = index
		}
		tmp += string(char)

		if index == eof || predicate.IsEndToken(next()) {
			tk, err := token.NewToken(pos, index, tmp)

			if err != nil {
				panic(err)
			}

			captured = append(
				captured,
				tk,
			)
			release()
		}
	}
}
