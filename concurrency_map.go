package stp

import (
	"sync"
)

// DEPRECATED: use sync.Map or other better implementation
// sync.Map is not used here because it does not support generic type

type CMapOption[K comparable, V any] func(*CMap[K, V])

func Updater[K comparable, V any](f func(K, V) V) CMapOption[K, V] {
	return func(m *CMap[K, V]) { m.updater = f }
}

type CMap[K comparable, V any] struct {
	noCopy
	rw      sync.RWMutex
	m       map[K]V
	updater func(K, V) V
}

func NewCMap[K comparable, V any](opts ...CMapOption[K, V]) *CMap[K, V] {
	m := &CMap[K, V]{
		rw: sync.RWMutex{},
		m:  make(map[K]V), // TODO: capacity
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
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

func (cm *CMap[K, V]) Range(f func(K, V) bool) {
	cm.rw.Lock()
	defer cm.rw.Unlock()
	for k, v := range cm.m {
		if !f(k, v) {
			return
		}
	}
}

func (cm *CMap[K, V]) Update(k K) (V, bool) {
	var v V
	if cm.updater == nil {
		return v, false
	}
	cm.rw.Lock()
	defer cm.rw.Unlock()
	v = cm.m[k]
	cm.m[k] = cm.updater(k, v)
	return v, true
}
