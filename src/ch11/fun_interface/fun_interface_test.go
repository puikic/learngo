package fun_interface

import (
	"fmt"
	"testing"
)

type SearchEngine struct {
	Recallers []Recaller
	Sorter    Sorter
}
type Recaller interface {
	Recall() []int
}
type Sorter interface {
	Sort([]int) []int
}

func r() []int {
	return nil
}
func s([]int) []int {
	return nil
}

type Rec struct {
}

func (Rec) Recall() []int {
	return r()
}

//	type SearchEngine2 struct {
//		Recallers []func() []int
//		Sorter    func() []int
//	}
type RecallType func() []int

func (rt RecallType) Recall() []int {
	return rt()
}
func TestDemo1(t *testing.T) {
	tmp := RecallType(r)
	fmt.Println(tmp.Recall())
	se := SearchEngine{Recallers: []Recaller{tmp}}
	fmt.Printf("%v", se)
}
func TestDemo2(t *testing.T) {
	x := make([]byte, 20)
	fmt.Println(x)
	x = append(x, 2)
	fmt.Println(len(x), cap(x), x)
}
