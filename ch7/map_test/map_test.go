package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	t.Log(m)                   // map[one:1 two:2]
	t.Log(m["one"])            // 1
	t.Log(m["onee"])           // 0
	t.Logf("m len=%d", len(m)) // m len=2

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
