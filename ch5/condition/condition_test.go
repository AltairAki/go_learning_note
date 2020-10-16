package condition

import (
	"fmt"
	"runtime"
	"testing"
)

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1=1")
	}

	// if的条件判断语句允许声明一个变量，但是这个变量的作用域只能在该条件逻辑内。
	// t.Log(a)

	// if result, err := someFunc(); err == nil {
	// 	t.Log("have not err")
	// } else {
	// 	t.Log("have err")
	// }
}

func TestSwitch(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("OS X.")
	case "linux":
		t.Log("Linux")
	default:
		t.Logf("%s.", os)
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 1, 3:
			t.Log(i, "Odd")
		case 0, 2:
			t.Log(i, "Even")
		default:
			t.Log(i, "it is not 0-3")
		}
	}
}

func TestSwitchConditionCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log(i, "Even")
		case i%2 == 1:
			t.Log(i, "Odd")
		default:
			t.Log("unknow")
		}
	}
}

/*
	Go语言里面switch默认相当于每个case最后带有break,匹配成功后不会自动向下执行其他case,而是跳出整个switch,但是可以使用 fallthrough强制执行后面的case代码
*/
func TestFallThrought(t *testing.T) {
	i := 6
	switch i {
	case 4:
		fmt.Println("the integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("the integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("the integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("the integer was <= 7")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
