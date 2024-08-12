package lex

import (
	"fmt"

	"github.com/intervinn/mx/node"
	"github.com/intervinn/mx/node/expression"
	"github.com/intervinn/mx/node/literal"
	"github.com/intervinn/mx/node/token"
)

func analyzeStatement() node.BaseNodeInterface {
	if s, ok := analyzeExpressionStatement(); ok {
		return s
	}
	return nil
}

func analyzeExpressionStatement() (node.BaseNodeInterface, bool) {
	left, ok := Type[*literal.IntLiteral](captured[0])
	if !ok {
		return nil, false
	}

	op, ok := Type[*token.PlusToken](captured[1])
	if !ok {
		return nil, false
	}

	right, ok := Type[*literal.IntLiteral](captured[2])
	if !ok {
		return nil, false
	}

	b := expression.NewBinaryExpression(left.Pos, right.End, left, op, right)

	fmt.Printf("%#v", b)
	return b, true
}
