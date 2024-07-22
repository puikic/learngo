package string_rune

import (
	"strconv"
	"strings"
	"testing"
)

// string_rune 是数据类型，不是引用或指针类型
// string_rune 是只读的 byte slice，len 函数可以它所包含的 byte 数
// string_rune 的 byte 数组可以存放任何数据
func TestString(t *testing.T) {
	var s string
	s = "hello "
	t.Log(len(s))
	s = "\xE4\xB8\xA5"
	t.Log(len(s), s, s[1]) // 16 * 11 + 8 = 184
}
func TestRune1(t *testing.T) {
	var s string = "中"
	r := []rune(s)
	t.Log(len(s), len(r))
	t.Logf("中 unicode %x", r[0])
	t.Logf("中 UTF8 %x", s)
}

func TestRune2(t *testing.T) {
	s := "中华人民共和国"
	t.Log(len(s), s[0], s[2])
	for i, v := range s {
		t.Log(i)
		t.Logf("%c", v)
	}

}

func TestStringsFn1(t *testing.T) {
	s := "A,B,C,D,E,陈,中"
	parts := strings.Split(s, ",")
	for i, v := range parts {
		t.Log(i, v)
	}
	tog := strings.Join(parts, "*")
	t.Log(tog)
}

func TestStringsFn2(t *testing.T) {
	s := strconv.Itoa(20)
	t.Log("10" + s)
	//x := strconv.Atoi("20")
	//Atoi的返回值有两个，和Itoa不一样
	if i, err := strconv.Atoi("20"); err == nil {
		t.Log(10 + i)
	}
}