package ticker

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println("ticker")
			}
		}
	}()
	time.Sleep(10 * time.Second)
}
