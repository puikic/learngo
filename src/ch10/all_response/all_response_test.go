package all_response

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func AllResponse() string {
	num := 10
	ch := make(chan string, num)
	for i := 0; i < num; i++ {
		go func(i int) {
			ret := RunTask(i)
			ch <- ret.(string)
		}(i)
	}
	finalRet := ""
	for j := 0; j < num; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}
func RunTask(i int) interface{} {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is %d", i)
}
func TestAllRes(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	t.Log("After:", runtime.NumGoroutine())

}
