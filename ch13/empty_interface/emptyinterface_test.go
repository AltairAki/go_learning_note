package emptyinterface

import (
	"fmt"
	"strconv"
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

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "name:" + p.name + " - age:" + strconv.Itoa(p.age)
}

func TestEmpty2(t *testing.T) {
	list := make(List, 4)
	list[0] = 1       //an int
	list[1] = "Hello" //a string
	list[2] = Person{"Dennis", 65}
	list[3] = []int{1, 2, 3}

	for i, e := range list {
		// if v, ok := e.(int); ok {
		// 	fmt.Printf("list[%d] is an int and its value is %d\n", i, v)
		// } else if v, ok := e.(string); ok {
		// 	fmt.Printf("list[%d] is a string and its value is %s\n", i, v)
		// } else if v, ok := e.(Person); ok {
		// 	fmt.Printf("list[%d] is a Person and its value is %s\n", i, v)
		// } else {
		// 	fmt.Printf("list[%d] is a different type\n", i)
		// }

		switch v := e.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", i, v)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", i, v)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", i, v)
		default:
			fmt.Printf("list[%d] is a different type\n", i)
		}
	}
}
