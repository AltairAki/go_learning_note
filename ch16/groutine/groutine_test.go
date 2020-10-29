package groutine_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 20; i++ {
		/*
			i的值被Test的协程及启动的其他协程共享了
			应该用锁的机制完成
		*/
		// go func() {
		// 	fmt.Printf("%x %d\n", &i, i)
		// }()

		/*go 的方法调用传递的时候都是值传递*/
		go func(i int) { //2. i的值在这里复制了一份，每个i的地址是不一样的
			fmt.Printf("%x %d\n", &i, i)
		}(i) // 1.传递i的同时
	}
	time.Sleep(time.Millisecond * 50)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func TestGoroutine2(t *testing.T) {
	go say("world") // 开启一个新的 Groutine 执行
	say("hello")    // 当前 Groutine 执行
}

/*
	hello
	world
	hello
	world
	hello
	world
	hello
	world
	hello
*/

/*
	上面多个goroutine运行在同一个进程里面，共享内存数据，不过设计上我们要遵循：不要通过共享来通信，而要通过通信来共享。

	runtime.Gosched()表示让CPU把时间片让给别人，下次某个时候继续恢复执行该goroutine。

	默认情况下，调度器仅使用单线程，也就是说只实现了并发。想要发挥多核处理器的并行，需要在我们的程序中显示调用runtime.GOMAXPROCS(n)告诉调度器同时使用多个线程。GOMAXPROCS设置了同时运行逻辑代码的系统线程的最大数量，并返回之前的设置。如果n<1，不会改变当前设置。以后Go语言的新版本中调度得到改进后，这将被移除。
*/
