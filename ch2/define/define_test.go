package define

import (
	"fmt"
	"testing"
)

/*定义变量*/
// 定义一个名称为"variableName",类型为"int"的变量
var variableName int

//定义三个类型都是"int"的三个变量
var vname1, vname2, vname3 int

//定义并初始化
var i int = 1

//同时初始化多个变量
var v1, v2, v3 int = 1, 2, 3

//可以忽略类型声明
var vv1, vv2, vv3 = 1, 2, 3

//编译器会根据初始化的值自动推导出相应类型,这种形式叫简短声明。只能用在函数内部，在函数外部使用则无法通关编译，所以一般用var方式定义全局变量。
// vvv1,vvv2,vvv3 := 1,2,3 //编译报错

// _(下划线)是个特殊的变量名，任何赋予它的值都会被丢弃。
func TestX(t *testing.T) {
	_, b := 31, 32
	fmt.Println(b)

	var p *int = &b
	fmt.Println(*p)
}

// GO对已声明但未使用的变量会在编译阶段报错。
func t() {
	v1, v2, v3 := 1, 2, 3
	fmt.Println(v1, v2, v3)
}
