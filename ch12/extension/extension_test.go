package extension

import (
	"fmt"
	"testing"
)

/*继承 Golang不支持继承*/
type Pet struct {
}

func (p *Pet) speak() {
	fmt.Println("... ")
}

func (p *Pet) speakTo(host string) {
	p.speak()
	fmt.Println(host)
}

type Dog struct {
	Pet
}

func (d *Dog) speak() {
	fmt.Print("wang ")
}

/*只有完全重写了Pet的SpeakTo()才能调用到Dog的speak()*/
// func (d *Dog) speakTo(host string) {
// 	d.speak()
// 	fmt.Println(host)
// }

func TestExtension(t *testing.T) {
	// p := new(Pet)
	// p.speakTo("host")

	d := new(Dog)
	// 未重写speakTo ... host
	// 重写speakTo wang host
	d.speakTo("host")
}
