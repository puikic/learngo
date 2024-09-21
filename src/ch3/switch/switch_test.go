package switch_test

import "testing"

func Switch1(x int) {
	switch {	
	case x >= 18:
		println("成年人")
	case x > 12:
		println("青年人")
	default:
		println("未知")
	}
	println("测试！")
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6, 8:
			t.Log("even")
		case 1, 3, 5, 7, 9:
			t.Log("odd")
		}
	}
}

func TestSwitch1(t *testing.T) {
	Switch1(11)
}
