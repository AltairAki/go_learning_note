package benchmark

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"h", "e", "l", "l", "o"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal("hello", ret)
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"h", "e", "l", "l", "o"}
	var buf bytes.Buffer
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	assert.Equal("hello", buf.String())
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"h", "e", "l", "l", "o"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem // string的+操作都会返回一个新的string都会有一个新的内存分配
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"h", "e", "l", "l", "o"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
