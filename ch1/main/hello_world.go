package main // 声明 main 包，表明当前是一个可执行程序
import (
	"fmt"
	"os"
)

func main() {
	// main 函数不支持传入参数
	// 通过os.Args获取命令行参数 获取到的是一个数组[二进制命令,值]
	/*
		PS D:\golearn\src\ch1\main> go run .\hello_world.go le
		[C:\Users\Administrator\AppData\Local\Temp\go-build373088054\b001\exe\hello_world.exe le]
	*/
	// fmt.Println(os.Args)

	/*
		PS D:\golearn\src\ch1\main> go run .\hello_world.go le
		hello world le
		exit status 1
	*/
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}
	// main 函数不支持任何返回值
	// 通过 os.Exit来返回状态
	os.Exit(1)
}
