package main

import "fmt"

type Pet struct {
	name string
	age  int
	eye  int
	leg  int
}

func main() {
	dog := Pet{"旺财", 1, 2, 2}
	fmt.Println(dog)
	fmt.Printf("%s", dog.name)
}
