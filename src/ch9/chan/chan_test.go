package _chan

import (
	"fmt"
	"testing"
)

func TestChan1(t *testing.T) {
	c1 := make(chan int, 1)
	//c1 := make(chan int) -> 阻塞
	c1 <- 1
	fmt.Println(<-c1)
}

func TestChan2(t *testing.T) {
	c1 := make(chan int, 2)

	c1 <- 1
	c1 <- 2
	close(c1)
	for v := range c1 {
		fmt.Println(v)
	}
}

func TestChan3(t *testing.T) {
	c1 := make(chan int, 2)

	c1 <- 1
	c1 <- 3
	<-c1
	x := <-c1
	t.Log(x)
}
