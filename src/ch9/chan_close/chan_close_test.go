package chan_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		defer func() {
			wg.Done()
			close(ch)
		}()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
}
func dataConsumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if data, ok := <-ch; ok {
				fmt.Println("data", data)
			} else {
				break
			}

		}
	}()
}
func TestChanClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataConsumer(ch, &wg)
	//wg.Add(1)
	//dataConsumer(ch, &wg)
	wg.Wait()
}
