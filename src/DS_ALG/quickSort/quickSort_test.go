package quickSort

import "testing"

func TestQuickSort(t *testing.T) {
	nums := []int{7, 2, 3, 5, 1, 8, 0, 2, 9, 0, 1, 5, 2, 8, 6}
	quickSort(nums)
	t.Log(nums)
}
