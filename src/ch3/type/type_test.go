package type__test

import (
	"testing"
)

type MyInt int64

func TestType1(t *testing.T) {
	var a int64 = 1
	var b int
	b = int(a)
	var c MyInt
	//c = a
	//cannot use a (variable of type int64) as MyInt value in assignment
	c = MyInt(a)
	t.Log(a, b, c)
}
func TestType2(t *testing.T) {
	//string_rune 是值类型，其默认的初始化值为空字符串，而不是ni
	a := "cpq我"
	var b string
	t.Logf("%T ", a)
	t.Log(b+a+b+a, len(a), len(b))
}
