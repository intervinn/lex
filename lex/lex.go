package lex

import (
	"fmt"

	"github.com/intervinn/mx/node"
)

func Type[T any](v any) (T, bool) {
	res, ok := v.(T)
	return res, ok
}

var (
	line  string
	tmp   string
	char  rune
	index int
	eof   int
	pos   int
	cmode int

	captured []node.BaseNodeInterface
)

func next() rune {
	return rune(line[index+1])
}

func clear() {
	captured = captured[:0]
}

func canCapture(k int) bool {
	return cmode == k || cmode == -1
}

func capture(k int) {
	cmode = k
}

func release() {
	cmode = -1
	tmp = ""
}

// Parses line string and returns the given expression.
func Line(l string) node.BaseNodeInterface {
	eof = len(l) - 1
	line = l
	clear()
	release()

	for index, char = range line {
		analyzeLiteral()
		analyzeToken()
	}

	fmt.Printf("%#v\n", captured)
	res := analyzeStatement()
	if res == nil {
		panic("undefined expression")
	}

	return res
}
