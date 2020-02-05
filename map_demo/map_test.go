package map_demo

import "testing"

func TestMapCase1(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		t.Log(k, v)
	}
}

func TestMapCase2(t *testing.T) {
	m := make(map[string]int, 10)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	delete(m, "a")
	for k, v := range m {
		t.Log(k, v)
	}
	if v, ok := m["a"]; ok {
		t.Log("a", "存在,value=", v)
	} else {
		t.Log("a", "不存在")
	}
}

func TestMapValueFunc(t *testing.T) {
	m := map[int]func(num int) int{}
	m[0] = func(num int) int {
		return num * num
	}

	m[1] = func(num int) int {
		return num * num * num
	}
	t.Log(m[0](3), m[1](3))
}

/*
* Go语言中没有set，但是可以直接用value为bool类型的map来模拟即可
 */
func TestMapForSet(t *testing.T) {
	m := map[string]bool{}
	m["chengdu"] = true
	m["shanghai"] = true
	m["suzhou"] = true
	delete(m, "shanghai")
	if _, ok := m["suzhou"]; ok {
		t.Log("suzhou存在于set之中")
	}
}
