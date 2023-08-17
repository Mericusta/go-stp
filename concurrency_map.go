package stp

import (
	"sync"
)

type CMap[K comparable, V any] struct {
	noCopy
	rw sync.RWMutex
	m  map[K]V
}

func NewCMap[K comparable, V any]() *CMap[K, V] {
	return &CMap[K, V]{
		rw: sync.RWMutex{},
		m:  make(map[K]V), // TODO: capacity
	}
}

func (cm *CMap[K, V]) Get(k K) (V, bool) {
	cm.rw.RLock()
	defer cm.rw.RUnlock()
	v, has := cm.m[k]
	return v, has
}

func (cm *CMap[K, V]) Save(k K, v V) (int, bool) {
	cm.rw.Lock()
	defer cm.rw.Unlock()
	if _, has := cm.m[k]; has {
		return 0, true
	}
	cm.m[k] = v
	l := len(cm.m)
	return l, false
}

func (cm *CMap[K, V]) Remove(k K) (V, int) {
	cm.rw.Lock()
	defer cm.rw.Unlock()
	v := cm.m[k]
	delete(cm.m, k)
	l := len(cm.m)
	return v, l
}
