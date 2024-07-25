package _interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	//断言 p.(xxx)
	//if i, ok := p.(int); ok {
	//	fmt.Println("int", i)
	//	return
	//}
	//if i, ok := p.(string); ok {
	//	fmt.Println(i)
	//	return
	//}
	//p.(type) 并不是一个单独的表达式或语法结构，而是用于 switch 语句中的类型选择器（Type Switch）

	switch t := p.(type) {
	case int:
		fmt.Println("int", t)
	case string:
		fmt.Println("string", t)
	default:
		fmt.Println("i dont know!")
	}

}
func TestEmptyInterface(t *testing.T) {

	DoSomething(1)
	DoSomething("2")
	DoSomething(true)

}
