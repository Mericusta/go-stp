package stp

import (
	"fmt"
	"testing"
)

type testBSTValue struct {
	i Int
	s String
}

func (v testBSTValue) LT(rv testBSTValue) bool {
	return v.i < rv.i && v.s < rv.s
}

func (v testBSTValue) GT(rv testBSTValue) bool {
	return v.i > rv.i && v.s > rv.s
}

func Test_BSTree(t *testing.T) {
	bst1 := NewBSTree(Int(3))
	bst2 := NewBSTree(testBSTValue{i: 3, s: "3"})
	for index := 0; index != 5; index++ {
		bst1.Set(Int(5 - index))
		bst2.Set(testBSTValue{
			i: Int(5 - index),
			s: String(fmt.Sprintf("%v", index)),
		})
	}
}
