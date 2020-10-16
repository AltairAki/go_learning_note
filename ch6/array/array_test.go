package array

/*
	array就是数组，定义方式如下：
	var arr [n]type
	由于长度也是数组类型的一部分，因此[3]int与[4]int
*/
import "testing"

func TestArrayInit(t *testing.T) {
	var a [3]int
	t.Log(a) // [0 0 0]
	a[0] = 1
	t.Log(a) // [1 0 0]

	b := [3]int{1, 2, 3}
	t.Log(b)
	bb := [10]int{1, 2, 3} //声明了一个长度为10的int数组，其中前三个元素初始化为1,2,3其他默认为0
	t.Log(bb)

	d := [...]int{1, 3, 4, 5, 7, 7} //可以忽略长度采用 `...`的方式，Go语言会自动根据元素的个数来计算长度
	t.Log(d)

	// 多维数组
	c := [2][3]int{{1, 2, 3}, {2, 3, 4}}
	t.Log(c) // [[1 2 3] [2 3 4]]
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
