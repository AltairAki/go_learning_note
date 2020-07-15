package emptyinterface

import (
	"fmt"
	"testing"
)

func getDataType(i interface{}) {
	// if v, ok := i.(int); ok {
	// 	fmt.Println("Interger", v)
	// 	return
	// }
	// if v, ok := i.(string); ok {
	// 	fmt.Println("String", v)
	// 	return
	// }
	// fmt.Println("Unknow Type")
	switch v := i.(type) {
	case int:
		fmt.Println("Interger", v)
	case string:
		fmt.Println("String", v)
	case []int:
		fmt.Println("slice", v)
	default:
		fmt.Println("Unkonw Type")
	}

}

func TestEmpty(t *testing.T) {
	getDataType(10)
	getDataType("10")
	getDataType([]int{1, 2})
}
