package stp

// NoBlockChannel 会导致失去顺序性，失去公平性

// type NoBlockChannel[T any] struct {
// 	noCopy
// 	c    chan T
// 	ctx  context.Context
// 	d    PoolDequeue[T]
// 	full bool
// }

// func NewNoBlockChannel[T any]() *NoBlockChannel[T] {
// 	c := &NoBlockChannel[T]{
// 		c:    make(chan T),
// 		ctx:  context.Background(),
// 		d:    NewPoolChain[T](),
// 		full: false,
// 	}
// 	go c.hold()
// 	return c
// }

// func NewNoBlockChannelWithBuffer[T any](b int64) *NoBlockChannel[T] {
// 	c := &NoBlockChannel[T]{
// 		c: make(chan T, b),
// 		d: NewPoolChain[T](),
// 	}
// 	go c.hold()
// 	return c
// }

// func (c *NoBlockChannel[T]) Send(v T) {
// 	select {
// 	case c.c <- v:
// 		break
// 	default:
// 		c.d.PushHead(v)
// 	}
// }

// func (c *NoBlockChannel[T]) Recv() <-chan T {
// 	return c.c
// }

// func (c *NoBlockChannel[T]) hold() {
// 	for {
// 		v, ok := c.d.PopTail()
// 		if v != nil && ok {

// 		}
// 	}
// }

// func (c *NoBlockChannel[T]) Close() {

// }
