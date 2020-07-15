package polymorphism

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	say() Code
}

type GoProgrammer struct {
}

func (g *GoProgrammer) say() Code {
	return "fmt.Println(\"Hello World!\")"
}

type PhpProgrammer struct {
}

func (p *PhpProgrammer) say() Code {
	return "print(\"Hello world!\")"
}

func say(p Programmer) {
	fmt.Printf("%T %v\n", p, p.say())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	phpProg := new(PhpProgrammer)
	say(goProg)  //*polymorphism.GoProgrammer fmt.Println("Hello World!")
	say(phpProg) //*polymorphism.PhpProgrammer print("Hello world!")
}
