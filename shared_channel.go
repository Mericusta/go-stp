package stp

import "sync/atomic"

type SharedChannel[T any] struct {
	noCopy
	c  chan T
	uc *int64
}

func NewSharedChannel[T any]() *SharedChannel[T] {
	c := &SharedChannel[T]{
		c:  make(chan T),
		uc: new(int64),
	}
	*c.uc = 1
	return c
}

func NewSharedChannelWithBuffer[T any](b int64) *SharedChannel[T] {
	c := &SharedChannel[T]{
		c:  make(chan T, b),
		uc: new(int64),
	}
	*c.uc = 1
	return c
}

func (c *SharedChannel[T]) Share() *SharedChannel[T] {
	atomic.AddInt64(c.uc, 1)
	return &SharedChannel[T]{c: c.c, uc: c.uc}
}

func (c *SharedChannel[T]) Get() chan T {
	return c.c
}

func (c *SharedChannel[T]) UseCount() int64 {
	return atomic.LoadInt64(c.uc)
}

func (c *SharedChannel[T]) Close() {
	if atomic.AddInt64(c.uc, -1) == 0 {
		close(c.c)
	}
}
