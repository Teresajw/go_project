package context

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewCtx(t *testing.T) {
	ctx1 := context.Background()

	ctx2 := context.TODO()

	t.Logf("ctx1: %v, %s", ctx1, reflect.TypeOf(ctx1).String())
	t.Logf("ctx2: %v, %s", ctx2, reflect.TypeOf(ctx2).String())
}

func TestCtx_WithValue(t *testing.T) {
	ctxA := context.Background()
	ctxB := context.WithValue(ctxA, "B", "B")
	ctxC := context.WithValue(ctxB, "C", "C")

	t.Log(ctxB)

	value1, value2 := ctxC.Value("B"), ctxC.Value("C")
	t.Log(value1, value2)
}

func TestCtx_WithCancel(t *testing.T) {
	bg := context.Background()
	ctx, cancelFunc := context.WithCancel(bg)
	defer cancelFunc()
	fmt.Println(ctx.Err())
}

func TestCtx_WithTimeout(t *testing.T) {
	tim := time.Now()
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	sonCtx, sonCancel := context.WithTimeout(ctx, 10*time.Second)
	defer sonCancel()
	select {
	case <-sonCtx.Done():
		fmt.Println("ctx时间:", time.Now().Sub(tim).Seconds())
		fmt.Println("父ctx超时导致子ctx结束：", ctx.Err())
	case <-time.After(3 * time.Second):
		fmt.Println("超时退出:", time.Now().Sub(tim).Seconds())
		fmt.Println("超出设定时间ctx：", ctx.Err())
	}
}

func TestCtx_WithValue1(t *testing.T) {
	ctx1 := context.Background()
	ctx2 := context.WithValue(ctx1, "map", map[string]string{})
	ctx3 := context.WithValue(ctx2, "key", "value")

	m := ctx2.Value("map").(map[string]string)
	val3 := ctx3.Value("key").(string)

	m["hellom"] = val3

	val1 := ctx2.Value("map")
	fmt.Println(val1)

}
