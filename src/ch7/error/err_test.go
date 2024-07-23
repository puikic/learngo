package error

import (
	"errors"
	"testing"
)

var LessThanErros = errors.New("less than err")
var LargerThanErros = errors.New("larger than err")

func GetFib(n int) ([]int, error) {
	//错误处理原则：及早失败
	if n < 0 {
		return nil, LessThanErros
	} else if n > 100 {
		return nil, LargerThanErros
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
func TestFibErr(t *testing.T) {
	if v, err := GetFib(10); err == LessThanErros {
		t.Error("so small")
	} else if err == LargerThanErros {
		t.Error("so large")
	} else {
		t.Log(v)
	}
}
