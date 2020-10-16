package const_test

/*
	常量是在编译阶段就确定下来的值，而程序在运行时则无法改变该值。Go中常量可定义为数值、布尔值或字符串等类型
*/
import (
	"testing"
)

const constantName = "value"

//如果需要，也可以明确指定常量类型
const Pi float32 = 3.1415926

/*
	iota 该关键字用来声明enum，默认开始值为0，每调用一次加1。
*/
const (
	Monday  = iota + 1
	Tuesday //常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说 Tuesday = iota + 1。因此Tuesday == 2
	Wednesday
)

const (
	Readable   = 1 << iota //001
	Writeable              //010
	Excuteable             //100
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConst2(t *testing.T) {
	a := 7 //0111
	t.Log(Readable, Writeable, Excuteable)
	t.Log(a&Readable, a&Writeable, a&Excuteable)
}
