package panic

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

/*
	Go语言没有其他语言那样的异常机制，不能抛出异常，而是使用panic和recover机制。一定套记住，你应当把它作为最后的手段来使用，也就是说你的代码中应当没有，或者很少panic的东西。

	panic是一个内建函数，可以中断原有的控制流程,进入一个令人恐慌的流程中。当函数F调用panic，函数F的执行被中断，但是F中的延迟函数会正常执行，然后F返回到调用它的地方。在调用的地方，F的行为就像调用了panic。这一过程继续向上，直到发生panic的goroutine中所有调用的函数返回，此时程序退出。恐慌可以直接调用panic产生，也可以由运行时错误产生，例如访问越界的数组。

	Recover是一个内建的函数，可以让进入令人恐慌的流程中的goroutine恢复过来。recover仅在延迟函数中有效。在正常的执行过程中，调用recover会返回nil，并且没有其他任何效果。如果当前的goroutine陷入恐慌，调用recover可以捕获到panic的输入值，并且恢复正常的执行。
*/
var user = os.Getenv("USER")

func init() {
	if user == "" {
		panic("no value for $USER")
	}
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
			b = true
		}
	}()
	f() //执行函数f,如果f中出现panic，那么就可以恢复过来了
	return
}

func TestPanic(t *testing.T) {
}

/*
	panic 用于不可以恢复的错误
	panic 退出前会执行 defer指定的错误

	os.Exit 退出时不会调用 defer 指定的函数
	os.Exit 退出时不输出当前调用栈信息
*/

func TestPanicVsExit(t *testing.T) {
	// defer func() {
	// 	fmt.Println("finally")
	// }()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from ", err)
			/*
				// log.Error("recovered panic",err)
				仅仅记录日志，这样的修复机制是非常危险的
				（可能是系统某些核心资源消耗完了，强制恢复以后还是不能正常工作的）
				· 形成僵尸服务进程，导致health check 失效（仅会检测进程是在还是不能提供服务的）
				· "Let it Crash!" 往往是我们恢复不确定性错误的最好方法
			*/
		}
	}()
	fmt.Println("Strat")
	panic(errors.New("Something wrong"))

	// os.Exit(-1)
}
