package maptest

import (
	"fmt"
	"testing"
)

func TestMapCase1(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		t.Log(k, v)
	}
	fmt.Println(len(m))
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

/*
* golang中的map，的 key 可以是很多种类型，比如 bool, 数字，string, 指针, channel , 还有只包含前面几个类型的 interface types, structs, arrays
* 显然，slice， map 还有 function 是不可以了，因为这几个没法用 == 来判断
 */
func TestKeySliceMap(t *testing.T) {
	const lenOfNum = 3
	// new出来的是数组而不是切片，如果用make出来就是切片
	array1 := new([lenOfNum]int)
	slice1 := []int{4, 5, 6}
	copy(array1[:], slice1)
	fmt.Printf("type = %T, content=%v\n", array1, array1)
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	resMap := map[interface{}]bool{}
	resMap[array1] = true
	fmt.Println(resMap, len(resMap))

	result := make([][]int, len(resMap))
	i := 0
	for k := range resMap {
		a := k.(*[lenOfNum]int)
		result[i] = a[:]
	}
	fmt.Println(result)
}
