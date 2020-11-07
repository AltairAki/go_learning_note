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

func TestConvert(t *testing.T) {
	i := 10
	// f := *(*float64)(&i) // 编译时会报 cannot convert &i (type *int) to type *float64
	f := *(*float64)(unsafe.Pointer(&i))

	fmt.Println(unsafe.Pointer(&i)) // 0xc0000122f8
	fmt.Println(f)                  // 5e-323

	i2 := *(*int)(unsafe.Pointer(&i))
	fmt.Println(i2) // 10

	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	fmt.Println(b) // [1 2 3 4]
}

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

}
