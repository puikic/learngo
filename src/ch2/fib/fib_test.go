package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	//var (
	//	a = 1
	//	b = 1
	//)
	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		temp := a
		a = b
		b = temp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
