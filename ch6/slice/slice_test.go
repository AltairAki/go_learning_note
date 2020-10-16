package slice

import (
	"fmt"
	"testing"
)

/*
	在很多应用场景中，数组并不能满足我们的需求。在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go语言里面这种数据结构叫slice。slice并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。
*/
func TestSliceInit(t *testing.T) {
	var s0 []int //声明和数组一样只是少了长度

	t.Log(len(s0), cap(s0)) // 0 0 len得到实际长度  cap可以得到容量即最大长度
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0)) // 2 2

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1)) // 4 4

	s2 := make([]int, 3, 4)
	// t.Log(s2[0], s2[1], s2[3]) // runtime error: index out of range [3] with length 3
	t.Log(len(s2), cap(s2)) // 3 4
	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2)) // 4 4

	// slice可以从一个数组或一个已存在的slice中再次声明。slice通过array[i:j]来获取，其中是开始位置，j是结束位置，但不包含array[j]，它的长度是j-i
	// slice还有一些简便操作：
	ar := [5]int{1, 2, 3, 4, 5}
	s3 := ar[:3] //等价于[0:3]
	t.Log(s3)
	s4 := ar[3:] //等价于[3:len(ar)]
	t.Log(s4)
	s5 := ar[:] //等价于[0:len(ar)]
	t.Log(s5)

}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
	/*
		TestSliceGrowing: slice_test.go:27: 1 1
		TestSliceGrowing: slice_test.go:27: 2 2
		TestSliceGrowing: slice_test.go:27: 3 4
		TestSliceGrowing: slice_test.go:27: 4 4
		TestSliceGrowing: slice_test.go:27: 5 8
		TestSliceGrowing: slice_test.go:27: 6 8
		TestSliceGrowing: slice_test.go:27: 7 8
		TestSliceGrowing: slice_test.go:27: 8 8
		TestSliceGrowing: slice_test.go:27: 9 16
		TestSliceGrowing: slice_test.go:27: 10 16
	*/
}

func TestMake(t *testing.T) {
	m := make([]int, 2, 6)
	t.Log(m[2])
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) // [Apr May Jun] 3 9 //为什么是9 切片指向的是一片连续的存储空间，剩下还有9个所以长度为9

	summber := year[5:8]
	t.Log(summber, len(summber), cap(summber)) // [Jun Jul Aug] 3 7
	summber[0] = "unknow"
	t.Log(Q2)   // [Apr May unknow]  因为Q2与summber是共享同一个后端数组，一旦这片连续的存储空间无论是被谁修改，都会影响到其他所有人
	t.Log(year) // [Jan Feb Mar Apr May unknow Jul Aug Sep Oct Nov Dec]
}

/*
	数组 VS 切片
	1.容量是否可伸缩： 数组容量是不可能伸缩的；
	2.是否可以进行比较： 相同长度，相同维度的数组是可以比较；
*/
func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	t.Log(a, b)
	// if a == b { // invalid operation: a == b (slice can only be compared to nil)
	// 	t.Log("equal")
	// }
}

func TestSliceTravel(t *testing.T) {
	s := []int{1, 3, 4, 7, 9}

	// for i := 0; i < len(s); i++ {
	// 	t.Log(s[i])
	// }

	// for index, value := range s {
	// 	t.Log(index, value)
	// }

	//用_占位
	for _, value := range s {
		fmt.Print(value)
	}
}
