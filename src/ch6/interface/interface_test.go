package _interface

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}
type GoProgrammer struct {
	Age  int
	Name string
}

func (p *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}
func TestInterface(t *testing.T) {
	var p Programmer
	//p = new(GoProgrammer)
	//p = GoProgrammer{} //错误写法，指针接收器的方法只属于*T
	p = &GoProgrammer{}
	//断言写法
	p.(*GoProgrammer).Age = 10
	t.Logf("%T\n %d\n", p, p.(*GoProgrammer).Age)
	t.Log(p.WriteHelloWorld())
}
