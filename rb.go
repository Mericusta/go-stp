package stp

type rbtNode[K comparable, V any] struct {
	k K
	v V
	c bool
	l *rbtNode[K, V]
	r *rbtNode[K, V]
}
