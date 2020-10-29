package ch20

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type SingletonObj struct {
}

var once sync.Once
var obj *SingletonObj

func GetSingletonObj() *SingletonObj {
	once.Do(func() {
		fmt.Println("Create obj.")
		// obj = &SingletonObj{}
		obj = new(SingletonObj)
	})
	return obj
}

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
