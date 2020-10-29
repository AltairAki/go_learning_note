package ch19_test

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		// ch <- 10  //panic: send on closed channel
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			// for i := 0; i < 11; i++ { // channel 在拿到10个int后已经关闭了，而这里又没有ok 所以将多返回一个零值 0
			data := <-ch
			fmt.Println(data)
		}
		wg.Done()
	}()
}

func dataReceiver2(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver2(ch, &wg)
	wg.Add(1)
	dataReceiver2(ch, &wg)

	// wg.Add(1)
	// dataReceiver(ch, &wg)
	wg.Wait()
}
