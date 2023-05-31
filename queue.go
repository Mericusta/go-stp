package stp

type PoolDequeue interface {
	PushHead(val any) bool
	PopHead() (any, bool)
	PopTail() (any, bool)
}

func NewPoolDequeue(n int) PoolDequeue {
	d := &poolDequeue{
		vals: make([]eface, n),
	}
	d.headTail = d.pack(1<<dequeueBits-500, 1<<dequeueBits-500)
	return d
}

func (d *poolDequeue) PushHead(val any) bool {
	return d.pushHead(val)
}

func (d *poolDequeue) PopHead() (any, bool) {
	return d.popHead()
}

func (d *poolDequeue) PopTail() (any, bool) {
	return d.popTail()
}

func NewPoolChain() PoolDequeue {
	return new(poolChain)
}

func (c *poolChain) PushHead(val any) bool {
	c.pushHead(val)
	return true
}

func (c *poolChain) PopHead() (any, bool) {
	return c.popHead()
}

func (c *poolChain) PopTail() (any, bool) {
	return c.popTail()
}
