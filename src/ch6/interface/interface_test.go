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
	p = new(GoProgrammer)
	t.Logf("%T\n", p)
	t.Log(p.WriteHelloWorld())
}
