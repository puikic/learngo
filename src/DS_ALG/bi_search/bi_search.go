package bi_search

import "golang.org/x/exp/constraints"

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
