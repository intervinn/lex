package expression

import (
	"github.com/intervinn/mx/kind"
	"github.com/intervinn/mx/node"
)

type BinaryExpression struct {
	node.BaseNode
	Left          node.BaseNode
	Right         node.BaseNode
	OperatorToken node.BaseNode
}

func NewBinaryExpression(pos, end int, left, op, right node.BaseNode) *BinaryExpression {
	return &BinaryExpression{
		BaseNode: node.BaseNode{
			Kind: kind.BinaryExpression,
			Pos:  pos,
			End:  end,
		},

		Left:          left,
		OperatorToken: op,
		Right:         right,
	}
}
