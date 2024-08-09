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
}
