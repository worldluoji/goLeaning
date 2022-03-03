package array

import (
	"fmt"
	"sort"
	"testing"
)

// $ go test -v array_test.go
// === RUN   TestArrayDeclaration
// --- PASS: TestArrayDeclaration (0.00s)
//     array_test.go:9: [1 2 3]
//     array_test.go:10: [1 2 5 0 0]
//     array_test.go:11: [5 6 7]
// PASS
// ok      command-line-arguments  0.703s
func TestArrayDeclaration(t *testing.T) {
	a1 := [...]int{1, 2, 3}
	a2 := [5]int{1, 2, 5}
	a3 := []int{5, 6, 7} // 未指定长度，是一个切片
	t.Logf("%T", a1)
	t.Log(a1)
	t.Logf("%T", a2)
	t.Log(a2)
	t.Log(a3)
	t.Logf("%T", a3)
}

func TestSliceAppend(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	a = append(a[:3], a[4:]...) // 删除第三个元素
	t.Log("删除第三个元素后的数组：", a)
	a = append(a, 7)
	a = append(a, []int{8, 9, 10}...)
	t.Log("在末尾添加元素后后的数组：", a)
	a = append(a, 0)
	copy(a[4:], a[3:]) // a[4:]用a[3:]赋值
	t.Log("在第三个位置添加元素后的数组：", a)
	a[3] = 4
	t.Log("在第三个位置添加元素后的数组：", a)
}

func TestSliceAppend2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := append(a, 7)
	fmt.Println(a)
	fmt.Println(b)
}

func add(a []int, val int) {
	a = append(a, val)
}

func TestSliceAppend3(t *testing.T) {
	a := make([]int, 3)
	a = append(a, 3, 4, 5)
	t.Log(a)
	add(a, 8)
	t.Log(a)
}

func TestSortIntSlice(t *testing.T) {
	sl := []int{3, 1, 2, 5, 6, 4}
	sort.Sort(sort.IntSlice(sl))
	t.Log("正序排序：", sl)
	sort.Sort(sort.Reverse(sort.IntSlice(sl)))
	t.Log("逆序排序：", sl)
	sort.Ints(sl)
	t.Log("正序排序：", sl)
}

func TestSortStringSlice(t *testing.T) {
	sl := []string{"a", "e", "b", "d", "c", "p"}
	sort.Strings(sl)
	t.Log("正序排序：", sl)
	sort.Sort(sort.Reverse(sort.StringSlice(sl)))
	t.Log("逆序排序：", sl)
}

func changeSlice(a []int, i int, newVal int) {
	a[i] = newVal
}

func changeArray(a [3]int, i int, newVal int) {
	a[i] = newVal
}

/*
	数组传入函数是副本，切片则是对应的地址
*/
func TestChangeArray(t *testing.T) {
	a := []int{1, 2, 3}
	changeSlice(a, 0, 6)
	fmt.Println(a)

	b := [3]int{1, 2, 3}
	changeArray(b, 0, 6)
	fmt.Println(b)

	fmt.Println(b[3:])
}

func TestArrayEqual(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	fmt.Println(a == b)
	fmt.Println(&a == &b)

	c := []int{1, 2, 3}
	d := []int{1, 2, 3}
	// fmt.Println(c == d) error
	fmt.Println(&c == &d)
}
