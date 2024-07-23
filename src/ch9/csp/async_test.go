package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 60)
	return "service finished"
}
func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("other task finished")
}
func AsyncService() chan string {
	//retCh := make(chan string)
	//working on something else
	//returned result:
	//other task finished
	//service finished
	//service existed
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result:")
		retCh <- ret
		fmt.Println("service existed")
	}()
	return retCh
}
func TestAsync1(t *testing.T) {
	fmt.Println(service())
	otherTask()
}
func TestAsync2(t *testing.T) {
	rerCh := AsyncService()
	otherTask()
	fmt.Println(<-rerCh)
	time.Sleep(time.Second)
}
