package close_context

import (
	"fmt"
	"sync"
	"testing"
)

func TestCanDemo1(t *testing.T) {
	ch := make(chan int)
	var wg sync.WaitGroup
	// 创建多个接收者 goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				value, ok := <-ch
				if !ok {
					fmt.Printf("Receiver %d: Channel closed\n", id)
					return
				}
				fmt.Printf("Receiver %d: Received %d\n", id, value)
			}
		}(i)
	}

	// 发送者 goroutine
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			fmt.Printf("Sender: Sending %d\n", i)
			ch <- i
		}
	}()
	wg.Wait()
}
