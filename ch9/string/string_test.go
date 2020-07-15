package string

import "testing"

/*
	1.string 是数据类型，不是引用或指针类型
	2.string 是只读的byte slice, len 函数返回的是 string 所包含的byte数
	3.string 的 byte 数组可以存放任何数据
*/

func TestString(t *testing.T) {
	var s string
	// t.Log(s)
	// s = "hello"
	// t.Log(len(s)) // 5

	// // s[1] = '2' // cannot assign to s[1] // string 是不可变的byte slice
	// s = "\xE4\xB8\xA5"
	// t.Log(s)      // 严
	// t.Log(len(s)) // 3  byte数

	s = "中"
	c := []rune(s)               // rune 将字符串转换为一个rune的切片
	t.Log(len(c))                // 1
	t.Logf("中 Unicode %x", c[0]) // 4e2d
	t.Logf("中 UTF8 %x", s)       // 输出中文“中”的三个byte  // e4b8ad
}

//strings包 https://golang.org/pkg/strings/
//strconv包 https://golang.org/pkg/strconv/

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}

/*
	TestStringToRune: string_test.go:35: 中 4e2d
    TestStringToRune: string_test.go:35: 华 534e
    TestStringToRune: string_test.go:35: 人 4eba
    TestStringToRune: string_test.go:35: 民 6c11
    TestStringToRune: string_test.go:35: 共 5171
    TestStringToRune: string_test.go:35: 和 548c
    TestStringToRune: string_test.go:35: 国 56fd
*/
