package first_response

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func FirstResponse() string {
	num := 10
	ch := make(chan string, num)
	for i := 0; i < num; i++ {
		go func(i int) {
			ret := RunTask(i)
			ch <- ret.(string)
		}(i)
	}
	return <-ch
}

func RunTask(i int) interface{} {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is %d", i)
}
func TestFirstRes(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	t.Log("After:", runtime.NumGoroutine())

}
