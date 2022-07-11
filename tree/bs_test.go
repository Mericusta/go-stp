package stptree

import (
	"fmt"
	"testing"

	stptype "github.com/Mericusta/go-stp/type"
)

type testBSTValue struct {
	i stptype.Int
	s stptype.String
}

func (v testBSTValue) LT(rv testBSTValue) bool {
	return v.i < rv.i && v.s < rv.s
}

func (v testBSTValue) GT(rv testBSTValue) bool {
	return v.i > rv.i && v.s > rv.s
}

func Test_BSTree(t *testing.T) {
	bst1 := NewBSTree(stptype.Int(3))
	bst2 := NewBSTree(testBSTValue{i: 3, s: "3"})
	for index := 0; index != 5; index++ {
		bst1.Set(stptype.Int(5 - index))
		bst2.Set(testBSTValue{
			i: stptype.Int(5 - index),
			s: stptype.String(fmt.Sprintf("%v", index)),
		})
	}
}
