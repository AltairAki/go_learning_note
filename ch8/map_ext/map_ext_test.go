/*
 * @Descripttion:
 * @Author: Altair
 * @Date: 2020-07-09 23:50:26
 */
package map_ext_test

import "testing"

/*简单工厂模式*/
func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2)) // 2 4 8
}

/*
	Go的内置集合中没有 Set 实现，可以 map[type]bool
	1.元素的唯一性
	2.基本操作
		1) 添加元素
		2) 判断元素是否存在
		3) 删除元素
		4) 元素个数
*/
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true

	/*1 is exesting*/
	n := 1
	if mySet[n] {
		t.Logf("%d is exesting", n)
	} else {
		t.Logf("%d is not exesting", n)
	}

	mySet[2] = true
	t.Log(len(mySet))
	delete(mySet, 1)

	if mySet[1] {
		t.Logf("%d is exesting", n)
	} else {
		t.Logf("%d is not exesting", n)
	}
}
