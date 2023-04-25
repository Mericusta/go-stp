package stp

import "unsafe"

// 内存反转函数
func tConvertByteArrayToObject[T any](b []byte) T {
	return *(*T)(unsafe.Pointer(&b))
}

type Pool[T any] struct {
	eCount int
	eSize  int
	last   int
	b      []byte
}

func NewPool[T any](c int) *Pool[T] {
	p := &Pool[T]{}
	p.allocateMemory(c)
	return p
}

func (p *Pool[T]) allocateMemory(c int) {
	var e T
	p.eSize = int(unsafe.Sizeof(e))
	p.eCount = c
	p.b = make([]byte, p.eCount*p.eSize)
}

func (p *Pool[T]) isFree(b, e int) bool {
	for _, b := range p.b[b:e] {
		if b != 0 {
			return false
		}
	}
	return true
}

func (p *Pool[T]) scan(b, e int) int {
	for i := b; i < e; i++ {
		if !p.isFree(i*p.eSize, (i+1)*p.eSize) {
			goto NEXT
		}
		return i
	NEXT:
	}
	return -1
}

func (p *Pool[T]) Get() *T {
	i := p.scan(p.last, p.eCount)
	if i == -1 {
		i = p.scan(0, p.last)
		if i == -1 {
			return nil
		}
	}

	p.last = (i + 1) % p.eCount
	return tConvertByteArrayToObject[*T](p.b[i*p.eSize : (i+1)*p.eSize])
}
