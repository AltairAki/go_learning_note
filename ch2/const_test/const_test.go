package const_test

import (
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable   = 1 << iota //001
	Writeable              //010
	Excuteable             //100
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConst2(t *testing.T) {
	a := 7 //0111
	t.Log(Readable, Writeable, Excuteable)
	t.Log(a&Readable, a&Writeable, a&Excuteable)
}
