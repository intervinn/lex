package expression

import (
	"github.com/intervinn/mx/kind"
	"github.com/intervinn/mx/node"
)

type BinaryExpression struct {
	node.BaseNode
	Left          node.BaseNodeInterface
	Right         node.BaseNodeInterface
	OperatorToken node.BaseNodeInterface
}

func NewBinaryExpression(pos, end int, left, op, right node.BaseNodeInterface) *BinaryExpression {
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
