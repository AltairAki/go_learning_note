package testdefer

import (
	"fmt"
	"testing"
)

/*
	延迟(defer)语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。特别是 当你 在 进行 一些 打开 资源 的 操作 时， 遇到 错误 需要 提前 返回， 在 返回 前 你 需要 关闭 相应 的 资源， 不然 很容易 造成 资源 泄露 等 问题。
*/

func ReadWrite(v int) bool {
	fmt.Println("open file")
	if v > 2 {
		fmt.Println("file close")
		return false
	}
	if v > 3 {
		fmt.Println("file close")
		return false
	}
	fmt.Println("file close")
	return true
}

/*
	上面有很多重复代码，Go语言的defer有效解决了这个问题。使用它后，代码量减少，而且程序也变得更加优雅。在defer后指定的函数会在函数退出前调用。
*/
func ReadWrite2(v int) bool {
	fmt.Println("open file")
	defer fmt.Println("file close")
	if v > 2 {
		return false
	}
	if v > 3 {
		return false
	}
	return true
}

func TestDefer(t *testing.T) {
	ReadWrite2(3)
}

//如果有很多调用defer，那么defer是采用后进先出模式
func TestMultiDefer(t *testing.T) {
	//输出43210
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}
