package interface_test

/*接口是定义对象之间交互的协议*/
/*
Go 接口 与其他主要编程语言的差异
1.接口为非入侵性，实现不依赖接口定义
2.所以接口的定义可以包含在接口使用者包内
*/
//Duck Type 式接口实现
import (
	"fmt"
	"strconv"
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

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) String() string {
	return h.name + " - " + strconv.Itoa(h.age) + "years - 📞" + h.phone
}

func (h Human) SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la ...", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi,I am %s,I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

// Interface Men 被 Human,student 和Employee 实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func TestDetail(t *testing.T) {
	mark := Student{Human{"Mark", 25, "029-222-YYYY"}, "MIT", 0.00}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc", 10.00}
	tom := Employee{Human{"Tom", 29, "111-888-XYZX"}, "Things Ltd", 100}

	//定义Men类型的变量i
	var i Men
	i = mark
	fmt.Println("This is mark ,a stduent:")
	i.SayHi()
	i.Sing("November rain")

	i = sam

	fmt.Println("This is mark ,an employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义 slice Men
	x := make([]Men, 3)
	x[0], x[1], x[2] = mark, sam, tom
	for _, v := range x {
		v.SayHi()
	}
}

func TestInterfaceParam(t *testing.T) {
	Bob := Human{"Bob", 29, "000-777-XXXX"}
	fmt.Println("This human is :", Bob)
}
