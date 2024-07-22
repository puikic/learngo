package _func

import (
	"fmt"
	"testing"
	"time"
)

// 所有参数都是值传递:slice，map，channel也是
func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent:", time.Since(start))
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second)
	return op
}

func TestTimeFn(t *testing.T) {
	tmp := timeSpent(slowFunc)
	t.Log(tmp(111))
	//tmp(10)
}
