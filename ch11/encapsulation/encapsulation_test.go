/*
 * @Descripttion:
 * @Author: Altair
 * @Date: 2020-07-09 23:50:26
 */
package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	ID   string
	Name string
	Age  int
}

/*
	Address is c0000aa520
	Address is c0000aa520
    TestStructOperations: encapsulation_test.go:45: ID:1-Name:ken-Age:20
*/
// func (e *Employee) toString() string {
// 	fmt.Printf("Address is %x \r\n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
// }

/*
	Address is c00005c520
	Address is c00005c550  // 重新复制了一份结构
    TestStructOperations: encapsulation_test.go:51: ID:1-Name:ken-Age:20
*/
func (e Employee) toString() string {
	fmt.Printf("Address is %x \r\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "ken", 20}
	t.Log(e) // {0 ken 20}

	e1 := Employee{Name: "marry", Age: 24}
	t.Log(e1) // { marry 24}

	e2 := new(Employee) //返回指针 相当于 e2 := &Employee{}
	e2.ID = "1"         //与其他语言的差异：通过实例的指针访问成员不需要使用->
	e2.Name = "Bob"
	e2.Age = 29
	t.Log(e2) // &{1 Bob 29}

	t.Logf("e is %T", e)   // e is encap.Employee
	t.Logf("e2 is %T", e2) // e2 is *encap.Employee

}

func TestStructOperations(t *testing.T) {
	e := Employee{"1", "ken", 20}
	fmt.Printf("Address is %x \r\n", unsafe.Pointer(&e.Name))
	t.Log(e.toString())
}
