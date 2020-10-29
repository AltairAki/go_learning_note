package stru

import (
	"fmt"
	"testing"
)

type person struct {
	name string
	age  int
}

func TestInitStruct(t *testing.T) {
	var P person

	P.name = "Bob"
	P.age = 25
	t.Log(P)

	P1 := person{"Tom", 24}
	t.Log(P1)

	P2 := person{age: 23, name: "marry"}
	t.Log(P2)
}

/*struct匿名字段
  Go语言支持只提供类型，而不写字段名的方式，也就是匿名字段，或称为嵌入字段。当匿名字段是一个
struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。
*/
type Human struct {
	name   string
	age    int
	weight int
}
type Student struct {
	Human      // 匿名字段，那么默认 Student 就包含了 Human 的所有字段
	speciality string
}

func TestAnonymousStruct(t *testing.T) {
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	t.Log(mark.name)

	mark.speciality = "AI"
	t.Log(mark.speciality)

	// 可以看到Student访问属性age和name时，就像访问自己所有用的字段一样
	// 匿名字段实现了字段的继承
	mark.age += 2
	t.Log(mark.age)

	//Student还可以访问Human字段作为字段名
	mark.Human = Human{"Marcus", 20, 155}
	mark.Human.age = 25
	t.Log(mark)
}

/*
	匿名访问和修改字段很有用，不仅struct字段，所有的内置类型和自定义类型都可以作为匿名字段。
*/
type Skills []string

type Student2 struct {
	Human      //匿名字段,struct
	Skills     //匿名字段，自定义类型 string slice
	int        //内置类型作为匿名字段
	speciality string
}

func TestAnonymousField(t *testing.T) {
	jane := Student2{Human: Human{"Jane", 26, 100}, Skills: []string{"php"}, speciality: "computer"}
	// 访问相应字段
	fmt.Println("Her name is", jane.name)
	fmt.Println("Her age is", jane.age)
	fmt.Println("Her weight is", jane.weight)
	fmt.Println("Her speciality is", jane.speciality)
	//修改skill技能字段
	jane.Skills = []string{"golang"}
	fmt.Println("Her skills are", jane.Skills)
	fmt.Println("She acquired two new ones")
	jane.Skills = append(jane.Skills, "c", "java")
	fmt.Println("Her skills are", jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is ", jane.int)
}

/* struct 同名字段
如果Human里面有个phone字段，而student里面也有一个phone字段，那么该怎么办呢
Go通过最外层的优先访问解决了这个问题，也就是当你访问student.phone的时候，是访问student里面的字段，而不是human里面的字段。
*/
type Man struct {
	name  string
	age   int
	phone string
}

type Employee struct {
	Man
	speciality string
	phone      string
}

func TestSameNameFeild(t *testing.T) {
	// Bob := Employee{Man{"Bob", 29, "86-13909290297"}, "Designer", "029-8111111"}
	Bob := Employee{Man: Man{"Bob", 29, "86-13909290297"}, speciality: "Designer"}
	fmt.Println("Bob`s work phone is:", Bob.phone)
	fmt.Println("Bob`s personal phone is:", Bob.Man.phone)
}
