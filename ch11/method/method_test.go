package method

/* method*/

import (
	"fmt"
	"math"
	"testing"
)

type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

func TestArea(t *testing.T) {
	area1 := area(Rectangle{12, 2})
	t.Log(area1)
}

/*
	method的语法如下：
	func (r ReceiverType) funcName(parameters) (results)
*/
type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func TestArea2(t *testing.T) {
	s := Square{10}
	fmt.Println("square`s area is", s.area())

	c := Circle{2}
	fmt.Println("circle`s area is", c.area())
}

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box //a slice of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	// (*b).color = c
	//当你用指针去访问相应字段时（虽然指针没有任何的字段）Go语言知道你要通过指针去获取这个值
	b.color = c
}

func (b1 BoxList) BiggestsColor() Color {
	v := 0.00
	c := Color(WHITE)
	for _, b := range b1 {
		if b.Volume() > v {
			v = b.Volume()
			c = b.color
		}
	}
	return c
}

func (b1 BoxList) paintColor(c Color) {
	for i := range b1 {
		// (&b1[i]).SetColor(c)
		// Go语言知道receiver是指针，它自动帮你转了
		b1[i].SetColor(c)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func TestBox(t *testing.T) {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 4, 4, BLACK},
		Box{2, 1, 1, BLUE},
		Box{12, 21, 2, WHITE},
		Box{3, 4, 5, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The color of the last one is ", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is ", boxes.BiggestsColor().String())
	fmt.Println("Let`s paint them in a certain color")
	boxes.paintColor(YELLOW)
	fmt.Println("The color of the second one is", boxes[1].color.String())
	fmt.Println("Obviously,now,th biggest one is", boxes.BiggestsColor().String())
}
