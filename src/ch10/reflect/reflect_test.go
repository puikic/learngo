package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	Class string
	User
}

func (u User) SayName(name string) {
	fmt.Println("my name is:", name)
}
func check1(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	fmt.Println(t, v)
	fmt.Println(v.FieldByIndex([]int{1, 1}))
	fmt.Println(v.FieldByName("User"))
	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		fmt.Println(f)
	}
}
func check2(i interface{}) {
	t := reflect.TypeOf(i)
	ty := t.Kind()
	if ty != reflect.Struct {
		fmt.Println("not struct")
	} else {
		fmt.Println("struct")
	}
}
func check3(i interface{}) {
	v := reflect.ValueOf(i)
	e := v.Elem()
	e.FieldByName("Class").SetString("二年三班111")
}
func check4(i interface{}) {
	v := reflect.ValueOf(i)
	m := v.Method(0)
	m.Call([]reflect.Value{reflect.ValueOf("ccppqq")})
}
func TestRfl1(t *testing.T) {
	u := User{"cpq", 23, true}
	s := Student{"三年二班", u}
	//check1(s)
	check3(&s)
	fmt.Println(s)
}
func TestRf2(t *testing.T) {
	u := User{"cpq", 23, true}
	//check1(s)
	check4(u)
}
