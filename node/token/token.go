package token

import (
	"github.com/intervinn/mx/kind"
	"github.com/intervinn/mx/node"
)

type BaseToken struct {
	node.BaseNode
}

func NewToken(tkind int) {
	switch tkind {
	case kind.PlusToken:
		break
	}
	return
}
