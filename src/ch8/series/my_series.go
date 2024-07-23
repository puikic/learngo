package series

import "fmt"

func init() {
	fmt.Println("series init1")
}
func init() {
	fmt.Println("series init2")
}

// 方法首字母大写以导出
func GetFib(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}
func Square(n int) int {
	return n * n
}
