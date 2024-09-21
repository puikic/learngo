package demon_test

import (
	"fmt"
	"testing"
)

func ChangeArr(arr []int) {
	arr[0] = 10010
}

func TestF1(t *testing.T) {
	a := []int{1, 2, 3, 4}
	ChangeArr(a)
	fmt.Println(a[0])
}
