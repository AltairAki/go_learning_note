package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
	两个类型都是int，第一个变量的类型可以省略，默认为离它最近的类型 返回值仅有一个时，()也就不需要了
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestFunc(t *testing.T) {
	max := max(10, 11)
	t.Log(max)
}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

/*
	官方建议可导出的函数，最好命名返回值，因为不命名返回值，虽然使代码更加简洁，但是会造成生成的文档可读性差。
*/
func ReturnMultiValues() (v1 int, v2 int) {
	v1 = rand.Intn(10)
	v2 = rand.Intn(20)
	return
}

func TestReturnMulti(t *testing.T) {
	v1, _ := ReturnMultiValues()
	t.Log(v1)
}

func timeSpend(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend:", time.Since(start).Seconds())
		return ret
	}
}

func slowFn(op int) int {
	time.Sleep(time.Second * 1)
	return op
}
func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)

	tsF := timeSpend(slowFn)
	t.Log(tsF(10))
}

/*
可变长参数
变量ops是一个int的slice。
*/
func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3))       // 6
	t.Log(sum(1, 2, 3, 4, 5)) // 15
}

/*
	传值与传指针
*/

//	当我们传递一个参数值到被调用的函数里面时，实际上是传了这个值的一份copy，当在被调用函数中修改参数值的时候，调用函数中相应实参b不会发生任何变化，因为数值变化只作用在copy上。
func add1(a int) int {
	a = a + 1
	return a
}
func TestPassValue(t *testing.T) {
	x := 3
	t.Log(x) //3
	//调用add1的时候，add1接收的参数其实是x的copy，而不是x本身
	x1 := add1(x)
	t.Log(x1) // 4
	t.Log(x)  // 3
}

//接收的参数改为指针类型，才能在函数中修改x的值。此时参数仍然是按copy传递的，只是copy的是一个指针。
func add2(a *int) int {
	*a = *a + 1
	return *a
}
func TestPassValue2(t *testing.T) {
	x := 3
	t.Log(x) //3
	//调用add1的时候，add1接收的参数其实是x的copy，而不是x本身
	x1 := add2(&x)
	t.Log(x1) // 4
	t.Log(x)  // 3
}

/*
	·传指针使得多个函数能操作同一个对象
	·传指针比较轻量级，只是传内存地址，可以用指针传递体积大的结构体。如果用参数值传递的话，在每次copy上面就会花费相对较多的系统开销（内存和时间）
	·Go语言中string,slice,map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。注意若函数需要改变slice的长度，则仍需要取地址传递指针
*/

/*
	defer 延迟函数，并不会立即执行，而是在函数返回前执行：
*/
func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Clear resources.")
	}()
	t.Log("Started")
	panic("Fatal error.") // defer仍会执行
	// fmt.Println("End")    // unreachable code
}

/*
	函数作为值、类型
	在Go语言中函数也是一种变量，通过type定义，它的类型就是所有拥有相同的参数，相同的返回值。
	type typeName func(input1 inputType1,input2 inputType2[,...])(result1,resultType1[,...])
*/
type funcInt func(int) bool //声明一个函数类型

func isOdd(i int) bool {
	if i%2 == 0 {
		return false
	}
	return true
}

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

//声明的函数类型在这个地方当做了一个参数
func filter(slice []int, f funcInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

/*
	函数当做值和类型传递在我们写通用接口的时候非常有用
*/
func TestFuncValue(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 7}
	odd := filter(slice, isOdd) //函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven) //函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)
}

/*
	main函数和init函数

	两个保留函数：init函数（能够应用于所有package）和main函数（只能应用于package main）。这两个函数在定义时不能有任何的参数和返回值。虽然一个package里面可以写任意多个init函数，但这无论是对于可读性还是以后的可维护性来说，我们都强烈建议用户在一个package中一个文件只写一个init函数。

	Go语言程序会自动调用init()和main()，所以你不需要在任何地方调用这两个函数。每个package中的init函数都是可选的，但package main就必须包含一个main函数。

	程序的初始化和执行都起始于main包。如果main包还导入了其他的包，那么在编译时就会将它们依次导入。有时一个包被多个包同时导入，那么它只会被程序的初始化和执行都起始于main包。如果main包还导入了其他的包，那么在编译时就会将它们依次导入。有时一个包被多个包同时导入，那么它只会被导入一次（例如很多包可能都会用到fmt包，但它只会被导入一次，因为没有必要导入多次）。当一个包被导入时，如果该包还导入了其他的包，那么会先将其他包导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行init函数（如果有的话），依此类推。等所有被导入的包都加载完毕了，就会开始对main包中的包级常量和变量进行初始化，然后执行main包中的init函数（如果存在的话），最后执行main函数。
*/

/*
	import

	import (
		"fmt"
	)
	fmt.Println("hello world")

	上面的fmt是Go语言的标准库，其实是去goroot下加载该模块，当然Go语言的import还支持如下两种方式来加载自己写的模块。
	1.相对路径
	import "./model"  //当前文件同一目录的model目录，但是不建议这种方式来import
	2.绝对路径
	import "shorturl/model" //加载 gopath/src/shorturl/model 模块

	还有一些特殊的import：
	1.点操作
	import (
		. "fmt"
	)
	该点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面调用的fmt.Println（"helloworld"）可以省略的写成Println（"helloworld"）。
	2.别名操作
	import (
		f "fmt"
	)
	别名操作调用包函数时前缀变成了我们的前缀，即f.Println("helloworld")。
	3. _ 操作
	import (
		"database/sql"
		_ "github.com/ziutek/mymysql/godrv"
	)
	_ 操作其实是引入该包，不直接使用包里面的函数，而是调用了该包里面的init函数。
*/
