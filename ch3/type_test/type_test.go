package type_test

import "testing"

/*
bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for utf8
rune // alias for int32,represents a Unicode code point
float32 float64
complex64 complex128

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
	不支持指针运算
	string是值类型，其默认的初始化值为空字符串，而不是 nil
*/
func TestPoin(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1       // invalid operation: aPtr + 1 (mismatched types *int and int)
	t.Log(a, aPtr)           // 1 0xc0000122c0
	t.Logf("%T %T", a, aPtr) //  int *int
}

func TestString(t *testing.T) {
	var str string //被初始化为空字符串
	t.Log("*" + str + "*")
	t.Log(len(str))
	if str == "" { // 与空字符串比较而非 nil
		t.Log("str is empty string")
	}
}
