package loop_test

import (
	"fmt"
	"testing"
)

func TestWhileLoop(t *testing.T) {
	n := 0
	/*while(n<5)*/
	for n < 5 {
		n++
		fmt.Print(" ", n)
	}
	/*while(true)*/
	// for {
	// 	fmt.Println(n)
	// }
}

func TestGoto(t *testing.T) {
	for m := 1; m < 10; m++ {
		n := 1
	LOOP: //标签名是大小写敏感的
		if n <= m {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
			n++
			goto LOOP
		} else {
			fmt.Println("")
		}
		n++
	}
}

func TestBreak(t *testing.T) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			t.Log(i, j)
			if j < 3 {
				break // 跳出的是j的循环
			}
		}
	}
	/*输出：
	0 0
	1 0
	2 0*/
}

/*
	for配合range可以用于遍历slice和map
*/
func TestRange(t *testing.T) {
	m := map[string]string{"first": "one", "secend": "two"}
	for k, v := range m {
		t.Log(k)
		t.Log(v)
	}
	for _, v := range m {
		t.Log(v)
	}
}
