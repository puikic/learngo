package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("pet is speaking...")
}
func (p *Pet) SpeakTo(name string) {
	p.Speak()
	fmt.Println("to " + name)
}

type Dog struct {
	//p *Pet
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("dog is speaking...")
}

//	func (d *Dog) SpeakTo(name string) {
//		d.Speak()
//		fmt.Println("to " + name)
//	}
func TestExt(t *testing.T) {
	d := new(Dog)
	//d.p.Speak()
	//d.p.SpeakTo("cpq")
	d.Speak()
	d.SpeakTo("cpq")
}
