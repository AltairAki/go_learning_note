package fibonacci

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	// var a int = 1
	// var b int = 1
	// var (
	// 	a int = 1
	// 	b int = 1
	// )
	a := 1 //自动类型推断
	b := 1
	fmt.Print(a)
	for i := 0; i < 10; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = tmp
	a, b = b, a
	t.Log(a, b)
}

func TestConst(t *testing.T) {
	a := 1
	t.Log(a)
}
