package ch18_test

/*
	csp(communicating sequential processes) 通过一个通道完成两个通讯实体之间的协调
*/

import (
	"fmt"
	"testing"
	"time"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func TestChannel(t *testing.T) {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	// go sum(a[:len(a)/2], c) // groutine1
	// go sum(a[len(a)/2:], c) // groutine2

	// x, y := <-c, <-c
	// fmt.Println(x, y, x+y) //-5 17 12

	go sum(a[:2], c)  // groutine1
	go sum(a[2:4], c) // groutine2
	go sum(a[4:], c)  // groutine3

	x, y, z := <-c, <-c, <-c
	fmt.Println(x, y, z, x+y+z) // 4 9 -1 12
}

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

/*总耗时0.15s*/
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsynService() chan string {
	// retCh := make(chan string)
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

/*异步返回结果：总耗时0.1s*/
func TestAsynService(t *testing.T) {
	retCh := AsynService() // 异步调用AsynService service是在另外一个协程中
	otherTask()
	fmt.Println(<-retCh) // 当需要结果时，取出结果
}
