package constraints

import "testing"

func TestSlice(t *testing.T) {
	// Slice[T]
	type Slice[T int | float32 | float64] []T
	var a Slice[int] = []int{1, 2, 3}
	t.Log(a)
	var b Slice[float32] = []float32{1.0, 2.45, 3}
	t.Log(b)
}

func TestMap(t *testing.T) {
	type MyMap[KEY float32 | float64, VALUE string | int] map[KEY]VALUE
	var m1 MyMap[float64, string] = map[float64]string{1.38: "1.388", 2.38: "2.388"}
	t.Log(m1[1.38])
	t.Log(m1)
}

func TestStruct(t *testing.T) {
	type MyStruct[T int | string] struct {
		Id   T
		Name string
	}

	var ms MyStruct[string] = struct {
		Id   string
		Name string
	}{Id: "string_cpq", Name: "name_cpq"}
	t.Log(ms.Id)

	// 范型借口
	type PrintData[T int | string | float64] interface {
		Print(data T)
	}

	// 范型通道
	type MyChan[T int | string] chan T
}
