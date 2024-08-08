package personal

import "sync"

type ConcurrentSlice[T any] struct {
	l *sync.RWMutex
	s []T
}

func NewConcurrentSlice[T any]() *ConcurrentSlice[T] {
	return &ConcurrentSlice[T]{
		l: &sync.RWMutex{},
		s: make([]T, 0),
	}
}

func (c *ConcurrentSlice[T]) Push(val T) {
	// 写锁
	c.l.Lock()
	defer c.l.Unlock()
	c.s = append(c.s, val)
}
