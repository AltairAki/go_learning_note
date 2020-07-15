package panic

import (
	"errors"
	"fmt"
	"testing"
)

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
