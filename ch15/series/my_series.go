package series

import "fmt"

/*
	package
	1. 基本复用模块单元
	以首字母大写来表明可被包外代码访问
	2. 代码的 package 可以和所在的目录不一致
	3. 同一目录里的 Go 代码的 package 要保持一致

	在环境变量里引入
	vim ~/.bash_profile
	export GOPATH="/User/altair/go:/User/altair/go_learning"
	export PATH="$HOME/user/"

	init 方法
	· 在 main 被执行前，所有依赖的 package 的 init 方法都会被执行
	· 不同包的 init 函数按照包导入的依赖关系决定执行顺序
	· 每个包可以有多个 init 函数
	· 包的每个源文件也可以有多个 init 函数，这点比较特殊
*/

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func GetFibonaciiSerie(n int) []int {
	fib := []int{1, 1}
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-2]+fib[i-1])
	}
	return fib
}

// 小写不能被包外引用
func square(i int) int {
	return i * i
}
