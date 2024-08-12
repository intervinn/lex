package node

type BaseNodeInterface interface {
	GetWidth() int
	GetText() string
}

type BaseNode struct {
	Pos int
	End int

	Kind int
	Text string
}

func (b *BaseNode) GetWidth() int {
	return len(b.Text)
}

func (b *BaseNode) GetText() string {
	return b.Text
}
