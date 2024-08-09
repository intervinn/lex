package literal

import (
	"github.com/intervinn/mx/kind"
	"github.com/intervinn/mx/node"
)

type IntLiteral struct {
	node.BaseNode
	BaseLiteral
}

func NewIntLiteral(pos, end int, str string) *IntLiteral {
	return &IntLiteral{
		node.BaseNode{
			Pos: pos,
			End: end,

			Kind: kind.IntLiteral,
			Text: str,
		},

		BaseLiteral{},
	}
}
