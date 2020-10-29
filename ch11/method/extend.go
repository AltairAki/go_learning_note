package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) sayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n", h.name, h.phone)
}

func (e *Employee) sayHi() {
	fmt.Printf("Hi,I am %s,I workat %s. Call me on %s\n", e.name, e.company, e.phone)
}

func main() {
	mark := Student{Human{"Mark", 25, "029-8210151"}, "MIT"}
	sam := Student{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.sayHi()
	sam.sayHi()
}
