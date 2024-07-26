package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetsingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Singleton Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}
func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			obj := GetsingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
		}()
	}
	wg.Wait()
}
