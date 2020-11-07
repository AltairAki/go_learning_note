package ch26

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type MyInt int

func TestUnsafeConvert(t *testing.T) {
	// 不安全的类型转换
	i := 10
	// f := *(*float64)(&i) // 编译时会报 cannot convert &i (type *int) to type *float64
	f := *(*float64)(unsafe.Pointer(&i))

	fmt.Println(unsafe.Pointer(&i)) // 0xc0000122f8
	fmt.Println(f)                  // 5e-323

	// 安全的类型转换
	i2 := *(*int)(unsafe.Pointer(&i))
	fmt.Println(i2) // 10

	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	fmt.Println(b) // [1 2 3 4]
}

/*
指针的原子操作

并发读写一个共享缓存的时候，为了读写的线程安全

在写的时候，写到另外一块地方，完全写完之后，用一个原子操作把我们读的那快地方和写的地方，重新指向一下。Buffer指向新写好的地方，这样再读的时候就会读新写好的内容，所以切换的时候要一个线程安全的特性，这时候可以用atomic 来做。
*/
func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println(data, *(*[]int)(data))
	}

	var wg sync.WaitGroup
	writeDataFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
	}
	// wg.Wait()
}
