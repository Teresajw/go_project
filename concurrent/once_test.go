package concurrent

import (
	"sync"
	"testing"
)

type OnceTest struct {
	o sync.Once
}

func (o *OnceTest) close() {
	o.o.Do(func() {
		println("close")
	})
}

func Test_closer(t *testing.T) {
	o := &OnceTest{}
	for i := 0; i < 10; i++ {
		o.close()
	}
}
