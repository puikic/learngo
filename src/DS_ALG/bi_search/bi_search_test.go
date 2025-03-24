package bi_search

import (
	"testing"
)

func TestBiSearch(t *testing.T) {
	arr := []int{0, 1, 3, 5, 6, 7, 9}
	res := BinarySearch[int](arr, 6)
	t.Log(res)
}
