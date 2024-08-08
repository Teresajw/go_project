package personal

import (
	"fmt"
	"testing"
)

const PK = 2

var ch = make(chan struct{}, 100)
var chProductionFlish = make(chan struct{})
var chConsumerFlish = make(chan struct{})

func product(ch chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
	}
	chProductionFlish <- struct{}{}
}

func consumer(ch chan struct{}) {
	for {
		if v, ok := <-ch; ok {
			fmt.Println(v)
		} else {
			break
		}
	}
	chConsumerFlish <- struct{}{}
}

func Test_channel(t *testing.T) {
	for i := 0; i < PK; i++ {
		go product(ch)
	}

	go consumer(ch)

	for i := 0; i < PK; i++ {
		<-chProductionFlish
	}

	close(ch)

	<-chConsumerFlish
}
