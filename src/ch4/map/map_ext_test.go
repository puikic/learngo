package _map

import "testing"

func TestMapExt(t *testing.T) {
	//go的map的value可以为func
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m)
	t.Log(m[1](2), m[2](2), m[3](2))
}
