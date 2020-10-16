package type_test

import (
	"errors"
	"fmt"
	"testing"
)

/*
内置基础类型
bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte // utf8的别称
rune // alias for int32,represents a Unicode code point
float32 float64
complex64 complex128
*/

// Boolean
var isActive bool                 //全局变量声明
var enable, disable = true, false //忽略类型的声明
func test() {
	var available bool //一般声明
	valid := false     //简短声明
	available = true   //赋值操作
	if available {
		fmt.Println(valid)
	}
}

/*
 字符串，Go采用UTF-8字符集编码 //标识符（因为UTF-8的发明者也就是Go语言的发明者）
 字符串是用一对双引号("")或反引号(``)括起来定义，类型是string
*/
func TestString(t *testing.T) {
	var str string //被初始化为空字符串
	t.Log("*" + str + "*")
	t.Log(len(str))
	if str == "" { // 与空字符串比较而非 nil
		t.Log("str is empty string")
	}
	// Go的字符串是不可变的，直接修改编译时会报错
	s := "hello"
	// s[0] = 'c'
	c := []byte(s) //将字符串转为 []byte 类型
	c[0] = 'H'
	s2 := string(c)
	t.Log(s2)

	//也可以这样修改字符串
	s1 := "hello"
	s1 = "C" + s1[1:] //字符串不可修改但是可以进行切片操作
	t.Log(s1)

	// 通过“`” 来声明多行字符串。
	m := `hello
	world`
	t.Log(m)
}

/*
数值

//类型预定义值
math.MaxInt64
math.MaxFloat64
math.MaxUint32
*/
type MyInt int64

func TestBase(t *testing.T) {
	var a int = 11
	var b int64 = 2

	// 不允许隐式类型转换
	// b = a // cannot use a (type int) as type int8 in assignment
	// a = b // cannot use b (type int8) as type int in assignment
	b = int64(a) // 必须得显式类型转换

	// t.Log(a, b)
	// 别名和原有类型也不能进行隐式类型转换
	var c MyInt = 3
	// c = b // cannot use b (type int64) as type MyInt in assignment
	c = MyInt(b) // 必须得显式类型转换
	t.Log(c)
}

/*
	错误类型
	Go语言内置有一个error类型，专门处理错误信息，Go的package里面还专门有一个包errors来处理错误
*/
func TestError(t *testing.T) {
	err := errors.New("new  error")
	if err != nil {
		t.Log(err)
	}
}

/*
	不支持指针的运算
	string是值类型，其默认的初始化值为空字符串，而不是 nil
*/
func TestPoin(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1       // invalid operation: aPtr + 1 (mismatched types *int and int)
	t.Log(a, aPtr)           // 1 0xc0000122c0
	t.Logf("%T %T", a, aPtr) //  int *int
}
