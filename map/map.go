package stpmap

import "fmt"

func Key[K comparable, V any](tm map[K]V) []K {
	ks, i := make([]K, len(tm)), 0
	for k := range tm {
		ks[i] = k
		i++
	}
	return ks
}

type Map[K comparable, V any] struct {
	m map[K]V
	l int
	c int
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{m: make(map[K]V)}
}

func (m *Map[K, V]) Set(k K, v V) {
	if _, has := m.m[k]; !has {
		m.l++
		m.c++
	}
	m.m[k] = v
}

func (m *Map[K, V]) Del(k K) {
	if _, has := m.m[k]; has {
		m.l--
		if (float64(m.l) * 6.5) < float64(m.c) {
			nm := make(map[K]V, len(m.m)-1)
			for k, v := range m.m {
				nm[k] = v
			}
			m.m = nm
			m.c = m.l
			fmt.Printf("reload\n")
		}
	}
	delete(m.m, k)
}

func (m Map[K, V]) Get(k K) (V, bool) {
	v, has := m.m[k]
	return v, has
}

func (m Map[K, V]) Key() []K {
	return Key(m.m)
}

func (m Map[K, V]) Range(f func(K, V) bool) {
	if f != nil {
		for k, v := range m.m {
			if !f(k, v) {
				return
			}
		}
	}
}

func (m Map[K, V]) Len() int {
	return m.l
}
