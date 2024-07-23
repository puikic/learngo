package error

import (
	"errors"
	"fmt"
	"testing"
)

//panic 用于不可以恢复的错误
//panic 退出前会执行 defer 指定的内容

// os.Exit 退出时不会调用 defer 指定的函数
// os.Exit 退出时不输出当前调用栈信息
func TestPanic1(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("start...")
	panic(errors.New("something went wrong"))
	//os.Exit(-1)
}
