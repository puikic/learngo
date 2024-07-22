package _map

import (
	"testing"
)

func TestMap1(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m, len(m))
	m["four"] = 4
	t.Log(m, len(m))
}

func TestMap2(t *testing.T) {
	m := make(map[string]int, 3)
	t.Log(m["cpq"]) // 不存在
	t.Log(m, len(m))
	m["one"] = 1 // 增加
	t.Log(m, len(m))
	m["two"] = 2
	t.Log(m, len(m))
	m["three"] = 3
	t.Log(m, len(m))
	m["four"] = 4
	t.Log(m, len(m))
	//区分value=0时，是不存在，还是value本身就为0
	m["zero"] = 0
	if v, ok := m["cpq"]; ok {
		t.Log(v)
	} else {
		t.Log("key cpq is empty.")
	}
	if v, ok := m["zero"]; ok {
		t.Log(v)
	} else {
		t.Log("key cpq is empty.")
	}
}

func TestMap3(t *testing.T) {
	m := map[string]string{"one": "1", "two": "2", "three": "3"}
	t.Log(m, len(m))
	t.Log(m["cpq"])
	m["zero"] = ""
	t.Log(m["zero"])
	if v, ok := m["cpq"]; ok {
		t.Log(v)
	} else {
		t.Log("key cpq is empty.")
	}
}
