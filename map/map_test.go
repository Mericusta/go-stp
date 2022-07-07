package stpmap

import (
	"math"
	"testing"
)

func Test_Key(t *testing.T) {
	m := make(map[int]int)
	km := make(map[int]interface{}, math.MaxInt16)
	for i := 0; i != math.MaxInt16; i++ {
		m[i] = i
		km[i] = struct{}{}
	}
	for _, k := range Key(m) {
		if _, hasK := km[k]; !hasK {
			panic("Test_Key failed")
		}
	}
}

func Test_Map(t *testing.T) {
	m := NewMap[int, int]()
	for i := 0; i != math.MaxInt16; i++ {
		m.Set(i, i)
		if m.l != m.c || m.l != i+1 {
			panic("Test_Map failed")
		}
	}
	if m.Len() != math.MaxInt16 {
		panic("Test_Map failed")
	}
	for i := 0; i != math.MaxInt16; i++ {
		if v, has := m.Get(i); !has || v != i {
			panic("Test_Map failed")
		}
	}
	for i := 0; i != math.MaxInt16; i++ {
		m.Del(i)
		if _, has := m.Get(i); has {
			panic("Test_Map failed")
		}
		if m.Len() != math.MaxInt16-i-1 {
			panic("Test_Map failed")
		}
	}
	for i := 0; i != math.MaxInt16; i++ {
		m.Set(i, i)
		if m.l != m.c || m.l != i+1 {
			panic("Test_Map failed")
		}
	}
	m.Range(func(i1, i2 int) bool {
		if i2%2 == 0 {
			m.Del(i2)
		}
		return true
	})
	for i := 0; i != math.MaxInt16; i++ {
		if _, has := m.Get(i); (i%2 == 0 && has) || (i%2 != 0 && !has) {
			panic("Test_Map failed")
		}
	}
}
