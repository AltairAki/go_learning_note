package error

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var ErrLessThanTwo = errors.New("n should be less than 2")
var ErrLargerHundred = errors.New("n should be larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, ErrLessThanTwo
	}
	if n > 100 {
		return nil, ErrLargerHundred
	}
	fidList := []int{1, 1}
	for i := 2; i < n; i++ {
		fidList = append(fidList, fidList[i-2]+fidList[i-1])
	}
	return fidList, nil
}

func TestGetFibonacci(t *testing.T) {
	if fib, err := GetFibonacci(1); err != nil {
		if err == ErrLessThanTwo {
			fmt.Println("异常信息为:N不能小于2")
		}
		if err == ErrLargerHundred {
			fmt.Println("异常信息为:N不能大于100")
		}
	} else {
		fmt.Println(fib)
	}
}

/*循环嵌套*/
func PrintFib1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error：", err)
		}
	} else {
		fmt.Println("Error：", err)
	}
}

/*应尽早失败，避免嵌套*/
func PrintFib2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error：", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error：", err)
		return
	}
	fmt.Println(list)
}

func TestGetFib(t *testing.T) {
	PrintFib1("1")
	PrintFib2("10")
}
