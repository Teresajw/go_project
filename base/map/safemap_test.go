package _map

import (
	"sync"
	"testing"
)

type SafeMap[K comparable, V any] struct {
	m    map[K]V
	lock sync.RWMutex
}

// 已经有key，则返回value，loaded为true
// 如果没有key，则放进去，loaded为false
// goroutine 1 => ("key1",1)
// goroutine 2 => ("key1",2)
func (s *SafeMap[K, V]) LoadOrStore(key K, newVal V) (V, bool) {
	s.lock.RLock()
	oldVal, ok := s.m[key]
	s.lock.RUnlock()
	if ok {
		return oldVal, true
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	//double check
	oldVal, ok = s.m[key]
	if ok {
		return oldVal, true
	}
	// goroutine1 先进来，那么这里就会变成key1=1
	// goroutine2 进来，那么这里就会变成key1=2,覆盖了
	s.m[key] = newVal
	return newVal, false
}

func Test_m(t *testing.T) {
	m1 := &SafeMap[string, int]{
		m: make(map[string]int, 10),
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		m1.LoadOrStore("key1", 1)
		wg.Done()
	}()
	go func() {
		m1.LoadOrStore("key1", 2)
		wg.Done()
	}()
	wg.Wait()

	t.Log(m1.m)
}
