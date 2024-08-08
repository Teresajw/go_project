package personal

import (
	"strconv"
	"sync"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	var wg sync.WaitGroup

	// 创建一个 ConcurrentMap 实例
	/*cm := NewConcurrentMap[string, int]()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key string, value int) {
			defer wg.Done()
			cm.Set(key, value)
		}(strconv.Itoa(i), i)
	}
	wg.Wait()
	t.Log(cm.m)*/

	// 创建一个普通map实例
	cm1 := make(map[string]int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key string, value int) {
			defer wg.Done()
			cm1[key] = value
		}(strconv.Itoa(i), i)
	}
	wg.Wait()
	t.Log(cm1)
}
