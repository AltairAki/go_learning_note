package array

import "testing"

func TestArrayInit(t *testing.T) {
	var a [3]int
	t.Log(a) // [0 0 0]
	a[0] = 1
	t.Log(a) // [1 0 0]

	b := [3]int{1, 2, 3}
	t.Log(b)

	c := [2][3]int{{1, 2, 3}, {2, 3, 4}}
	t.Log(c) // [[1 2 3] [2 3 4]]

	d := [...]int{1, 3, 4, 5, 7, 7}
	t.Log(d)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 3, 4, 7, 9}

	// for i := 0; i < len(arr); i++ {
	// 	t.Log(arr[i])
	// }

	// for index, value := range arr {
	// 	t.Log(index, value)
	// }

	//用_占位
	for _, value := range arr {
		t.Log(value)
	}
}

/*
数组截取
a[开始索引(包含):结束索引(不包含)]
*/
func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	// arrSec := arr[:6] // invalid slice index 6 (out of bounds for 5-element array)
	arrSec1 := arr[:3] // [1 2 3]
	arrSec2 := arr[3:] // [4 5]
	t.Log(arrSec1, arrSec2)
}
