package node

type BaseNodeInterface interface {
	GetWidth() int
	GetText() string
}

type BaseNode struct {
	BaseNodeInterface
	Pos int
	End int

	Kind int
	Text string
}

func (b *BaseNode) Width() int {
	return len(b.Text)
}
