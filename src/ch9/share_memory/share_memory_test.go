package share_memory

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestShareMem1(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second)
	t.Logf("counter: %d", counter)
}
func TestShareMem2(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second)
	t.Logf("counter: %d", counter)
}

func TestShareMem3(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
				wg.Done()
			}()
			mut.Lock()
			counter++
		}()
	}
	wg.Wait()
	t.Logf("counter: %d", counter)
}

func TestShareMem4(t *testing.T) {
	var wg sync.WaitGroup
	lock := new(sync.Mutex)
	var a int32 = 0
	var b int32 = 2

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(index int) {
			lock.Lock()
			defer lock.Unlock()
			if a > b {
				fmt.Println("done")
				wg.Done()
				return
			}

			a++
			fmt.Printf("i: %d a: %d \n", index, a)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
