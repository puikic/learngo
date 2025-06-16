package sequence

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println(add(1, 2))
}

func add(a, b int) int {
	return a + b
}
