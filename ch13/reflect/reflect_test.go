package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	// t := reflect.TypeOf(i)
	// v := reflect.ValueOf(i)

	// tag := t.Elem().Field(0).Tag
	// name := v.Elem().Field(0).String()

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}

func TestSetReflect(t *testing.T) {
	var x float64 = 3.4

	// 错误示例
	// v := reflect.ValueOf(x)
	// v.SetFloat(7.1)

	//修改
	p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(7.1)

	fmt.Println(x)
}
