package _func_test

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

func a() (res int) {
	fmt.Println("a()")
	res = 1
	return
}
func TestA(t *testing.T) {
	res := a()
	fmt.Println(res)
}

func testPanic() (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			result = -1
		}
	}()
	fmt.Println("Before panic")
	panic("Something went wrong")
	fmt.Println("After panic") // 不会执行
	return 1
}

func TestPanic(t *testing.T) {
	result := testPanic()
	fmt.Println("Result:", result)
}
