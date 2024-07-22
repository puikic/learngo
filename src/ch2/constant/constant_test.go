package constant

import "testing"

const (
	//iota从0开始
	Mon = iota + 1
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConst1(t *testing.T) {
	t.Log(Fri, Sun)
}
func TestConst2(t *testing.T) {
	a := 7 //0111
	//Readable	0001
	//W	0010
	t.Log(a&Readable | Writable)
	//0011
}
