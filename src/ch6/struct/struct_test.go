package _struct

import (
	"fmt"
	"testing"
)

type Employee struct {
	Name string
	Id   string
	Age  int
}

func (e Employee) String() string {
	fmt.Printf("address : %x\n", &e.Name)
	return fmt.Sprintf("===Name:%s, Id:%s, Age:%d===", e.Name, e.Id, e.Age)
}

//	func (e *Employee) String() string {
//		fmt.Printf("address : %x\n", &e.Name)
//		return fmt.Sprintf("===Name:%s, Id:%s, Age:%d===", e.Name, e.Id, e.Age)
//	}
func TestStrcut(t *testing.T) {
	e := Employee{"cpq", "123888", 23}
	fmt.Printf("address : %x\n", &e.Name)
	t.Log(e.String())
}
