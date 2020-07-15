package operator_test

import (
	"testing"
)

const (
	Readable = 1 << iota
	Writeable
	Excuteable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{4, 2, 6, 8}
	// b := [...]int{1, 3, 5, 7, 9}
	c := [...]int{2, 4, 6, 8}
	d := [...]int{2, 4, 6, 8}
	// t.Log(a == b) // invalid operation: a == b (mismatched types [4]int and [5]int)
	t.Log(a == c) // false 位置不同也不相等
	t.Log(c == d) // true 长度，位置一致才相等
}

/*
&^ 按位置零 [右边的是1则结果就置零，右边是0则左边是什么结果就是什么]
1 &^ 0 -- 1
1 &^ 1 -- 0
0 &^ 1 -- 0
0 &^ 0 -- 0
*/
func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a1 := a &^ Readable
	a2 := a &^ Writeable
	a3 := a &^ Excuteable

	t.Log(a1 == Readable, a2 == Writeable, a3 == Excuteable)
}
