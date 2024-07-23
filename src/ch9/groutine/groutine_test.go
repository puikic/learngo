package groutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine1(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1000 * time.Millisecond)
}
func TestGroutine2(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(500 * time.Millisecond)
}
