package close_context

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(c chan struct{}) bool {
	select {
	case <-c:
		return true
	default:
		return false
	}
}
func cancel_1(c chan struct{}) {
	c <- struct {
	}{}
}
func cancel_2(c chan struct{}) {
	close(c)
}
func TestCancel(t *testing.T) {
	c := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, c chan struct{}) {
			for {
				if isCancelled(c) {
					break
				}
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println(i, "Cancelled")
		}(i, c)
	}
	//cancel_1(c)
	cancel_2(c)
	time.Sleep(time.Second)
}
