package token

import (
	"fmt"

	"github.com/intervinn/mx/kind"
	"github.com/intervinn/mx/node"
)

type BaseToken struct {
	*node.BaseNode
}

func MatchKind(s string) int {
	switch s {
	case "+":
		return kind.PlusToken
	case "-":
		return kind.MinusToken
	case "/":
		return kind.SlashToken
	case "*":
		return kind.AsteriskToken
	}

	return -1
}

func NewToken(pos, end int, text string) (node.BaseNodeInterface, error) {
	k := MatchKind(text)
	if k == -1 {
		return nil, fmt.Errorf("unknown token: %s", text)
	}

	base := BaseToken{
		&node.BaseNode{
			Pos: pos,
			End: end,

			Kind: k,
			Text: text,
		},
	}

	var res node.BaseNodeInterface

	switch k {
	case kind.AsteriskToken:
		res = &AsteriskToken{base}
	case kind.PlusToken:
		res = &PlusToken{base}
	case kind.MinusToken:
		res = &MinusToken{base}
	case kind.SlashToken:
		res = &SlashToken{base}
	}

	return res, nil
}
