package token

import (
	"github.com/intervinn/mx/node"
)

type BaseToken struct {
	node.BaseNode
	Text string
}

func NewToken(pos, end int, text string, tkind int) *BaseToken {
	return &BaseToken{
		BaseNode: node.BaseNode{
			Pos: pos,
			End: end,

			Kind: tkind,
		},
		Text: text,
	}
}

func (b *BaseToken) GetWidth() int {
	return len(b.Text)
}

func (b *BaseToken) GetText() string {
	return b.Text
}
