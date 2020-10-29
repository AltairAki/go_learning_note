package customertype

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {

	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}

}

func SlowFn(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestSpendTime(t *testing.T) {
	sFn := timeSpent(SlowFn)
	t.Log(sFn(5))
}
