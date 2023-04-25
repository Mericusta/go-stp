package stp

import (
	"sync/atomic"
	"unsafe"
)

type eface struct {
	typ, val unsafe.Pointer
}

type PoolDequeue struct {
	headTail uint64
	vals     []eface
}

const (
	dequeueBits  = 32
	dequeueLimit = (1 << dequeueBits) / 4
)

type DequeueNil *struct{}

func (d *PoolDequeue) Unpack(ptrs uint64) (head, tail uint32) {
	const mask = 1<<dequeueBits - 1
	head = uint32((ptrs >> dequeueBits) & mask)
	tail = uint32(ptrs & mask)
	return
}

func (d *PoolDequeue) Pack(head, tail uint32) uint64 {
	const mask = 1<<dequeueBits - 1
	return (uint64(head) << dequeueBits) |
		uint64(tail&mask)
}

func (d *PoolDequeue) PushHead(val any) bool {
	ptrs := atomic.LoadUint64(&d.headTail)
	head, tail := d.Unpack(ptrs)
	if (tail+uint32(len(d.vals)))&(1<<dequeueBits-1) == head {
		return false
	}
	slot := &d.vals[head&uint32(len(d.vals)-1)]

	typ := atomic.LoadPointer(&slot.typ)
	if typ != nil {
		return false
	}

	if val == nil {
		val = DequeueNil(nil)
	}
	*(*any)(unsafe.Pointer(slot)) = val

	atomic.AddUint64(&d.headTail, 1<<dequeueBits)
	return true
}

func (d *PoolDequeue) PopHead() (any, bool) {
	var slot *eface
	for {
		ptrs := atomic.LoadUint64(&d.headTail)
		head, tail := d.Unpack(ptrs)
		if tail == head {
			return nil, false
		}

		head--
		ptrs2 := d.Pack(head, tail)
		if atomic.CompareAndSwapUint64(&d.headTail, ptrs, ptrs2) {
			slot = &d.vals[head&uint32(len(d.vals)-1)]
			break
		}
	}

	val := *(*any)(unsafe.Pointer(slot))
	if val == DequeueNil(nil) {
		val = nil
	}
	*slot = eface{}
	return val, true
}

func (d *PoolDequeue) PopTail() (any, bool) {
	var slot *eface
	for {
		ptrs := atomic.LoadUint64(&d.headTail)
		head, tail := d.Unpack(ptrs)
		if tail == head {
			return nil, false
		}

		ptrs2 := d.Pack(head, tail+1)
		if atomic.CompareAndSwapUint64(&d.headTail, ptrs, ptrs2) {
			slot = &d.vals[tail&uint32(len(d.vals)-1)]
			break
		}
	}

	val := *(*any)(unsafe.Pointer(slot))
	if val == DequeueNil(nil) {
		val = nil
	}

	slot.val = nil
	atomic.StorePointer(&slot.typ, nil)

	return val, true
}

type PoolChain struct {
	head *PoolChainElt
	tail *PoolChainElt
}

type PoolChainElt struct {
	PoolDequeue
	next, prev *PoolChainElt
}

func StorePoolChainElt(pp **PoolChainElt, v *PoolChainElt) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(pp)), unsafe.Pointer(v))
}

func LoadPoolChainElt(pp **PoolChainElt) *PoolChainElt {
	return (*PoolChainElt)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(pp))))
}

func (c *PoolChain) PushHead(val any) {
	d := c.head
	if d == nil {
		const initSize = 8
		d = new(PoolChainElt)
		d.vals = make([]eface, initSize)
		c.head = d
		StorePoolChainElt(&c.tail, d)
	}

	if d.PushHead(val) {
		return
	}

	newSize := len(d.vals) * 2
	if newSize >= dequeueLimit {
		newSize = dequeueLimit
	}

	d2 := &PoolChainElt{prev: d}
	d2.vals = make([]eface, newSize)
	c.head = d2
	StorePoolChainElt(&d.next, d2)
	d2.PushHead(val)
}

func (c *PoolChain) PopHead() (any, bool) {
	d := c.head
	for d != nil {
		if val, ok := d.PopHead(); ok {
			return val, ok
		}
		d = LoadPoolChainElt(&d.prev)
	}
	return nil, false
}

func (c *PoolChain) PopTail() (any, bool) {
	d := LoadPoolChainElt(&c.tail)
	if d == nil {
		return nil, false
	}

	for {
		d2 := LoadPoolChainElt(&d.next)
		if val, ok := d.PopTail(); ok {
			return val, ok
		}
		if d2 == nil {
			return nil, false
		}
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&c.tail)), unsafe.Pointer(d), unsafe.Pointer(d2)) {
			StorePoolChainElt(&d2.prev, nil)
		}
		d = d2
	}
}
