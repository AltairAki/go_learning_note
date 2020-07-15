package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpend(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend:", time.Since(start).Seconds())
		return ret
	}
}

func slowFn(op int) int {
	time.Sleep(time.Second * 1)
	return op
}
func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)

	tsF := timeSpend(slowFn)
	t.Log(tsF(10))
}

/*可变长参数*/
func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3))       // 6
	t.Log(sum(1, 2, 3, 4, 5)) // 15
}

/*
	defer 延迟函数，并不会立即执行，而是在函数返回前执行：
*/
func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Clear resources.")
	}()
	t.Log("Started")
	panic("Fatal error.") // defer仍会执行
	// fmt.Println("End")    // unreachable code
}
