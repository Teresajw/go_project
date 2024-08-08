package concurrent

import (
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	p := sync.Pool{
		New: func() any {
			return &User{}
		},
	}

	u1 := p.Get().(*User)
	u1.Id = 1
	u1.Name = "zhangsan"
	p.Put(u1)
	u2 := p.Get().(*User)
	t.Log(u2)
}

type User struct {
	Id   int
	Name string
}

func Test_channel(t *testing.T) {
	a := []int{1, 2, 3}
	ch := make(chan int)
	for _, v := range a {
		ch <- v
	}
}
