package sharememory

import (
	"sync"
	"testing"
	"time"
)

/*
package sync
Mutex
RWMutex
*/

/*线程不安全操作，引发的数据不一致操作*/
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second * 1)
	t.Logf("counter = %d", counter) // counter = 4923
}

func TestCounterGroutineSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() { // 加锁的时候一般都要配合 defer 来释放锁
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second * 1)     //1秒后所有协程运行完毕
	t.Logf("counter = %d", counter) // counter = 5000
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1) // 每启一个协程WaitGroup就+1
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done() //完成协程任务，主动告诉WaitGroup结束
		}()
	}
	wg.Wait()                       //等每个WaitGroup都 done 才不会继续阻塞
	t.Logf("counter = %d", counter) // counter = 5000
}
