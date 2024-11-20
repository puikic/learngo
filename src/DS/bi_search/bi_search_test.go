package bi_search

import (
	"golang.org/x/exp/constraints"
	"testing"
)

func BinarySearch[T constraints.Ordered](arr []T, target T) int {
	begin := 0
	end := len(arr) - 1
	for begin <= end {
		mid := (begin + end) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return -1
}
func TestBiSearch(t *testing.T) {
	arr := []int{0, 1, 3, 5, 6, 7, 9}
	res := BinarySearch[int](arr, 6)
	t.Log(res)
}
