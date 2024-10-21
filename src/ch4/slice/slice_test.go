package slice

import "testing"

// 把slice理解为一个结构体：1.一个指针，指向低层数组 2.len, cap int
func ChangeSlice1(nums []int) {
	nums[3] = 999
}

func ChangeSlice2(nums *[]int) {
	(*nums)[0] = 929
}

func TestDemo(t *testing.T) {
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

func reverse1(i int, nums []int) {
	if i == 5 {
		return
	}
	nums = append(nums, i)
	reverse1(i+1, nums)
}
func reverse2(i int, nums *[]int) {
	if i == 5 {
		return
	}
	*nums = append(*nums, i)
	reverse2(i+1, nums)
}
func sliceAdd(nums []int) {
	// 值传递
	for i := 0; i < 5; i++ {
		nums = append(nums, i)
	}
}
func sliceAddP(nums *[]int) {
	for i := 0; i < 5; i++ {
		*nums = append(*nums, i)
	}
}

func change(nums []int) {
	nums[0] = 10
}
func TestDemo1(t *testing.T) {
	nums := make([]int, 0, 20)
	reverse1(0, nums)
	t.Log(nums)
}
func TestDemo2(t *testing.T) {
	nums := make([]int, 0, 20)
	reverse2(0, &nums)
	t.Log(nums)
}
func TestDemo3(t *testing.T) {
	nums := make([]int, 3, 20)
	sliceAdd(nums)
	t.Log(nums)
}
func TestDemo4(t *testing.T) {
	nums := make([]int, 3, 20)
	sliceAddP(&nums)
	t.Log(nums)
}
func TestDemo5(t *testing.T) {
	nums := make([]int, 3, 20)
	change(nums)
	t.Log(nums)
}

func TestAppend(t *testing.T) {
	nums := []int{1, 2, 3}
	temp := append(nums, 5)
	t.Log(nums)
	t.Log(temp)
}

func TestMake(t *testing.T) {
	s := []int{}
	ss := make([]int, 0, 100)
	t.Log(len(s), cap(s))
	t.Log(len(ss), cap(ss))
	ssp := &ss
	t.Logf("%T %v", s, s)
	t.Logf("%T %v", ss, ss)
	t.Logf("%T %v", ssp, ssp)
}
