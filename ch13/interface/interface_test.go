package interface_test

/*接口是定义对象之间交互的协议*/
/*
Go 接口 与其他主要编程语言的差异
1.接口为非入侵性，实现不依赖接口定义
2.所以接口的定义可以包含在接口使用者包内
*/
//Duck Type 式接口实现
import (
	"testing"
)

//接口定义
type Programmer interface {
	WriteHelloWorld() string
}

//接口实现
type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World!\"})"
}

func TestInterface(t *testing.T) {
	var p Programmer      // 接口变量
	p = new(GoProgrammer) //GoProgrammer就是p的类型，数据就是&GoProgrammer{}
	t.Log(p.WriteHelloWorld())
}
