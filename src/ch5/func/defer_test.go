package _func

import "testing"

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("test defer !")
	}()
	t.Log("start========")
	//panic("-----------------")
	//t.Log("end========")
}
