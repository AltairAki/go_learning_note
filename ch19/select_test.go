package ch19_test

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 100)
	return "Done"
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

var ErrTimeOut = errors.New("timeout")

/*异步返回结果：总耗时0.1s*/
func TestSelectTimeout(t *testing.T) {
	select {
	case retCh := <-AsynService():
		fmt.Println(retCh)
	case <-time.After(time.Millisecond * 99):
		fmt.Println("Error：", ErrTimeOut)
		break
	}
}
