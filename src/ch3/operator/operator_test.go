package operator

import "testing"

func TestOp1(t *testing.T) {
	// Go的数组是值类型，不是引用类型
	// 相同维数且含有相同个数元素的数组才可以比较
	// 每个元素都相同的才相等
	a := [...]int{1, 2, 3, 4, 5, 6}
	b := [...]int{1, 2, 3, 4, 5, 6}
	//c := [...]int{1, 2, 3, 4, 5, 6, 7}
	t.Log(a == b)
	//t.Log(a == c)
	//invalid operation: a == c (mismatched types [6]int and [7]int)
}
func TestOp2(t *testing.T) {
	//&^按位置零
	a := 3 //011
	b := 4 //100
	c := 1 //001
	t.Log(a &^ b)
	t.Log(b &^ a)
	t.Log(c &^ b) //1
	t.Log(c &^ a)
	t.Log(a &^ c) //2
}
