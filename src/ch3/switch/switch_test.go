package _switch

import "testing"

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
