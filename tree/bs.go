package stptree

type BSTComparable[T any] interface {
	lessThan(T) bool
	greaterThan(T) bool
}

type BSTree[T BSTComparable[T]] struct {
	r *bstNode[T]
}

func NewBSTree[T BSTComparable[T]]() *BSTree[T] {
	return &BSTree[T]{}
}

func (t *BSTree[int]) Set(v int) {
	if t.r == nil {
		t.r = newBSTNode(v)
		return
	}
	n := t.r
	for {
		if n.v.lessThan(v) {
			if n.l == nil {
				n.l = newBSTNode(v)
				return
			} else {
				n = n.l
			}
		} else if n.v.greaterThan(v) {
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

type bstNode[T BSTComparable[T]] struct {
	v T
	l *bstNode[T]
	r *bstNode[T]
}

func newBSTNode[VT BSTComparable[VT]](v VT) *bstNode[VT] {
	return &bstNode[VT]{v: v}
}
