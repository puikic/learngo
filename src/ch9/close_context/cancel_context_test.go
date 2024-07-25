package close_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled1(c context.Context) bool {
	select {
	case <-c.Done():
		return true
	default:
		return false
	}
}
func TestCancel1(t *testing.T) {
	c, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, c context.Context) {
			for {
				if isCancelled1(c) {
					break
				}
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println(i, "Cancelled")
		}(i, c)
	}
	//cancel_1(c)
	cancel()
	time.Sleep(time.Second)
}
