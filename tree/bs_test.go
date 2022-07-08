package stptree

import (
	"fmt"
	"testing"
)

type testBSTValue struct {
	i int
	s string
}

func (v testBSTValue) lessThan(rv testBSTValue) bool {
	return v.i < rv.i && v.s < rv.s
}

func (v testBSTValue) greaterThan(rv testBSTValue) bool {
	return v.i > rv.i && v.s > rv.s
}

func Test_BSTree(t *testing.T) {
	bst := NewBSTree[testBSTValue]()
	for index := 0; index != 5; index++ {
		bst.Set(testBSTValue{i: 5 - index, s: fmt.Sprintf("%v", index)})
	}
}
