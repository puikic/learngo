package arr

import "testing"

func TestArr1(t *testing.T) {
	s := []int{}
	t.Log(len(s), cap(s))
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(s, len(s), cap(s))
	}
}
func TestArr2(t *testing.T) {
	a := [3]int{1, 2, 3}
	t.Log(a[0], a[1], a[2])
	//a[3] = 0
	//index 3 out of bounds [0:3]
}
