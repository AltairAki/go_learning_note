package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFunc(t *testing.T) {
	s := "a,b,c"
	parts := strings.Split(s, ",")
	t.Log(parts) // [a b c]

	for _, v := range parts {
		t.Log(v)
	}
	t.Log(strings.Join(parts, "-")) // a-b-c
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s) // str10

	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(1 + i) // 11
	}
}
