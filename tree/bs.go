package stptree

import (
	stptype "github.com/Mericusta/go-stp/type"
)

type BSTree[T stptype.STPOrdered[T]] struct {
	r *bstNode[T]
}

func NewBSTree[T stptype.STPOrdered[T]](v T) *BSTree[T] {
	return &BSTree[T]{r: newBSTNode(v)}
}

func (t *BSTree[T]) Set(v T) {
	if t.r == nil {
		t.r = newBSTNode(v)
		return
	}
	n := t.r
	for {
		if n.v.LT(v) {
			if n.l == nil {
				n.l = newBSTNode(v)
				return
			} else {
				n = n.l
			}
		} else if n.v.GT(v) {
			if n.r == nil {
				n.r = newBSTNode(v)
				return
			} else {
				n = n.r
			}
		} else {
			return
		}
	}
}

type bstNode[T any] struct {
	v T
	l *bstNode[T]
	r *bstNode[T]
}

func newBSTNode[T stptype.STPOrdered[T]](v T) *bstNode[T] {
	return &bstNode[T]{v: v}
}
