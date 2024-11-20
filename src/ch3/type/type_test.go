package type_test

import (
	"fmt"
	"strconv"
	"testing"
	"unicode/utf8"
)

type MyInt int64
type s1 []int
type MyString []string

func (mi MyInt) string() string {
	return strconv.Itoa(int(mi))
}

func TestType1(t *testing.T) {
	var test s1 = []int{1, 2, 3}
	t.Log(test)
	var test2 s1
	var intSlice = []int{3, 2, 1}
	test2 = intSlice
	t.Log(test2)

	var s = []string{"cpq", "123"}
	var ms MyString
	ms = s
	t.Log(ms)

	var a int64 = 1
	var b int
	b = int(a)
	var c MyInt
	// c = a
	// cannot use a (variable of type int64) as MyInt value in assignment
	c = MyInt(a)
	t.Log(a, b, c)
}
func TestType2(t *testing.T) {
	//string_rune 是值类型，其默认的初始化值为空字符串，而不是ni
	a := "cpq我"
	var b string
	t.Logf("%T ", a)
	t.Logf("%v", b)
	t.Log(b+a+b+a, len(a), len(b))
	t.Log(utf8.RuneCountInString(a), utf8.RuneCountInString(b))
	t.Log(len("陈"))
	rune1 := []rune{}
	rune1 = append(rune1, '陈')
	rune1 = append(rune1, '哈')
	t.Log(len(rune1))
}

func TestType3(t *testing.T) {
	var i MyInt = 8
	fmt.Println(i.string())
	fmt.Printf("%T\n", i.string())
}

func TestType4(t *testing.T) {
	var a interface{} = "abc"
	switch a.(type) {
	case string:
		fmt.Println("string:", a)
	}
	m := a.(string) // 类型断言
	fmt.Println(m)
	fmt.Printf("%T\n", m)
}
