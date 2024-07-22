package _map

import "testing"

func TestMapForSet(t *testing.T) {
	//用Map[type]bool实现Set数据类型
	set := make(map[int]bool)
	set[1] = true
	set[2] = true
	//判断存在
	n := 3

	//删除元素
	delete(set, 2)
	t.Log(len(set), set)
	if set[n] {
		t.Logf("set[%d] is true", n)
	} else {
		t.Logf("set[%d] is false", n)
	}
}
