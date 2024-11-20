package constraints

import (
	"fmt"
	"testing"
)

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

type MySlice[T int | float32 | float64] []T

// 自定义泛型约束
type MyInt interface {
	int | int8 | int16 | int32 | int64
}
type aliasString string
type MyS []int

func (ms MySlice[T]) Sum() T {
	var sum T
	for _, v := range ms {
		sum += v
	}
	return sum
}

func add[T MyInt | float32 | float64 | ~string](a, b T) T {
	return a + b
}
func myPrint[T MyInt | float32 | float64 | ~string | ~[]int](x T) {
	fmt.Println(x)
}

func TestMethodAndFunc(t *testing.T) {
	// 方法不能在函数内定义！
	var ms MySlice[float64] = []float64{1.2, 2.3, 4.5}
	t.Log(ms.Sum())

	// 方法不能嵌套方法
	var s1 = "test"
	var s2 = "~method"
	t.Log(add(s1, s2))
	//t.Log(add[aliasString]("test", "~method")) 错误
	var ss1 aliasString = "alias"
	var ss2 aliasString = "cpq"
	//t.Log(add(ss1, ss2)) 错误
	//若想执行上面的操作，应该把add泛型约束中的string改为~string
	t.Log(add(ss1, ss2))
	t.Log(add[aliasString](ss1, ss2))

	var t1 = []int{1, 2, 999}
	var t2 = MyS{999, 888, 222}
	myPrint(t1)
	myPrint(t2)
}
