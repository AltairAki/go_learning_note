package map_test

/*
	map：字典，它的格式为map[keyType]valueType。
	map的读取和设置与slice类似，通过key操作，只是slice的index只能是 `int`类型，而map可以是int，可以是string及所有完全定义了==与!=操作的类型。
	需要注意：
	· map是无序的，每次打印出来的map都会不一样，不能通过index获取，而必须通过key获取。
	· map的长度不是固定的，和slice一样，也是一种引用类型。
	· 内置的len函数同样适用于map，返回map拥有的key的数量。
	· map的值也很容易修改，通过 map["key"] = value 即可修改key的字典值改为value
*/
import (
	"fmt"
	"testing"
)

func TestMapInit(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}

	//map有两个返回值，第二个返回值，如果key不存在，那么ok为false,如果存在ok为true
	mOne, ok := m["one"]
	if ok {
		fmt.Println("One is in the map and its value is ", mOne)
	} else {
		fmt.Println("We have no value associated with one in the map")
	}
	t.Log(m)                   // map[one:1 two:2]
	t.Log(m["one"])            // 1
	t.Log(m["onee"])           // 0
	t.Logf("m len=%d", len(m)) // m len=2

	//因为map是引用类型，如果两个map指向同一底层，那么一个改变，另一个也相应改变。
	mm := make(map[string]string)
	mm["hello"] = "world"
	mm1 := mm
	mm1["hello"] = "kenny" // 现在mm["hello"]的值已经是 kenny了
	t.Log(mm)
	m1 := map[string]int{}
	// m["name"] = "1"    // cannot use "1" (type untyped string) as type int in assignment
	m1["one"] = 1
	t.Log(m1)                    // map[one:1]
	t.Logf("m1 len=%d", len(m1)) // m1 len=1

	m2 := make(map[string]int, 10 /*initial capacity*/)
	// 为什么不初始化len ? slice初始化len的时候会对应的len位会初始化为零值，map无法做默认的零值
	// 与slice一样会自增容量，因此初始化的时候赋一定大的容量可以减少性能消耗
	t.Log(m2)                    // map[]
	t.Logf("m2 len=%d", len(m2)) // m2 len=0
	// t.Log(cap(m2)) // invalid argument m2 (type map[string]int) for cap
}

/*
	make、new操作
	make用于内建类型（map、slice和channel）的内存分配。new用于各种类型的内存分配

	new返回指针：new(T) 分配了零值填充的T类型的内存空间，并返回其地址，即一个*T类型的值。用Go的术语来说，它返回了一个指针，指向新分配的类型T的零值。

	make(T,args)与new(T)有着不同的功能，make只能创建slice，map和channel，并返回一个有初始值(非零)的T类型，而不是*T。
	本质来讲，导致这三个类型有所不同的原因是，指向数据结构的引用在使用前必须被初始化。例如：一个slice，是一个包含	指向数据(内部array)的指针、长度和容量的三项描述符，在这些项目被初始化之前，slice为nil。对于slice、map和channel来说，make初始化了内部的数据结构，填充了适当的值。
	make返回初始化后的（非零）值。

*/

func TestMapNotExistingKey(t *testing.T) {
	m := map[int]int{}
	m[1] = 1

	t.Log(m[3]) // 0 // 在访问的key不存在时，仍会返回零值，不能通过返回 nil 来判断元素是否存在

	/*key 3 is not existing.*/
	if v, ok := m[3]; ok {
		t.Logf("key 3`s value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}

	/*key 3`s value is 2*/
	m[3] = 2
	if v, ok := m[3]; ok {
		t.Logf("key 3`s value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestMapTravel(t *testing.T) {
	m := map[int]int{1: 123, 2: 27}

	for k, v := range m {
		t.Log(k, v)
	}
}
