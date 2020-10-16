package groutine_test

import (
	"fmt"
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
