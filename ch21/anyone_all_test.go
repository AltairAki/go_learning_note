package ch21

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

/*
	仅需任意任务完成
*/
func FirstResponse() string {
	numOfRunner := 10
	// ch := make(chan string)
	/*
		不用buffer channel 会导致阻塞的协程一直被阻塞
		TestFirstResponse: any_test.go:33: Before: 2
		TestFirstResponse: any_test.go:34: The result is from 0
		TestFirstResponse: any_test.go:36: After: 11
	*/
	// 用buffer channel 防止协程泄露，如果是非buffer那么放入消息无接收者就会被阻塞
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	// channel 的机制：往里面放消息后，在接收这个channel的接收者就会从阻塞中唤醒
	return <-ch // 这里return 就是尝试从channel中获取数据，一旦有数据就立即return
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}

/*
	全部任务完成
*/
func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log(AllResponse())
}
