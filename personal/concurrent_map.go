package personal

import "sync"

type ConcurrentMap[T int | float64 | string | bool, V any] struct {
	s *sync.RWMutex
	m map[T]V
}

func NewConcurrentMap[T int | float64 | string | bool, V any]() *ConcurrentMap[T, V] {
	return &ConcurrentMap[T, V]{
		s: new(sync.RWMutex),
		m: make(map[T]V),
	}
}

func (c *ConcurrentMap[T, V]) Set(key T, value V) {
	//加锁
	c.s.Lock()
	defer c.s.Unlock() //延迟解锁
	c.m[key] = value
}

func (c *ConcurrentMap[T, V]) Get(key T) (value V, ok bool) {
	//加读锁
	c.s.RLock()
	defer c.s.RUnlock()
	value, ok = c.m[key]
	return
}

func (c *ConcurrentMap[T, V]) Delete(key T) {
	c.s.Lock()
	defer c.s.Unlock()
	delete(c.m, key)
}
