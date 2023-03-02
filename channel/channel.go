package stpchannel

import "sync/atomic"

type SharedChannel[T any] struct {
	c  chan T
	uc *int64
}

func NewSharedChannel[T any]() *SharedChannel[T] {
	sc := &SharedChannel[T]{
		c:  make(chan T),
		uc: new(int64),
	}
	*sc.uc = 1
	return sc
}

func NewSharedBufferChannel[T any](b int64) *SharedChannel[T] {
	sc := &SharedChannel[T]{
		c:  make(chan T, b),
		uc: new(int64),
	}
	*sc.uc = 1
	return sc
}

func (sc *SharedChannel[T]) Share() *SharedChannel[T] {
	atomic.AddInt64(sc.uc, 1)
	return &SharedChannel[T]{c: sc.c, uc: sc.uc}
}

func (s *SharedChannel[T]) Get() chan T {
	return s.c
}

func (s *SharedChannel[T]) UseCount() int64 {
	return atomic.LoadInt64(s.uc)
}

func (sc *SharedChannel[T]) Close() {
	if atomic.AddInt64(sc.uc, -1) == 0 {
		close(sc.c)
	}
}
