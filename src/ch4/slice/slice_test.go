package slice

import "testing"

func ChangeSlice1(nums []int) {
	nums[3] = 999
}

func ChangeSlice2(nums *[]int) {
	(*nums)[0] = 929
}

func TestDemo1(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ChangeSlice1(nums)
	ChangeSlice2(&nums)
	t.Log(nums)
}

func TestSlice(t *testing.T) {
	//切片共享存储结构
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ss := []int(s)
	t.Log(ss)
	t.Log(len(s), cap(s))
	s1 := s[:5]
	s2 := s[3:6]
	s3 := s[7:]
	t.Log(s1, s2, s3)
	t.Log(len(s1), cap(s1))
	t.Log(len(s2), cap(s2))
	t.Log(len(s3), cap(s3))
	s1[4] = 100
	t.Log(s, s1, s2, s3)
	//	s[10] = 20
	//error: index out of range [10] with length 10
}
