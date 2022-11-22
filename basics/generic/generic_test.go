package generic

import (
	"fmt"
	"testing"

	"golang.org/x/exp/constraints"
)

func Add[T constraints.Integer](a, b T) T {
	return a + b
}

func TestGeneric(t *testing.T) {
	var m, n int = 5, 6
	println(Add(m, n)) // 自动推导, 等价于 Add[int](m, n)
	var i, j int64 = 15, 16
	println(Add(i, j)) // Add[int64](i, j)
	var c, d byte = 0x11, 0x12
	println(Add(c, d)) // Add[byte](c, d)
}

func Sort[Elem interface{ Less(y Elem) bool }](list []Elem) {
	// bubble sort
	n := len(list)
	for i := n - 1; i >= 0; i-- {
		for j := i; j > 0; j-- {
			if list[j].Less(list[j-1]) {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
}

type book struct {
	id   int
	name string
}

func (x book) Less(y book) bool {
	return x.id < y.id
}

func TestGeneric2(t *testing.T) {

	book1 := book{id: 2, name: "老人与海"}
	book2 := book{id: 1, name: "西游记"}
	fmt.Println("book1 < book2 ? ", book1.Less(book2))

	var bookshelf []book
	bookshelf = append(bookshelf, book1, book2)

	bookSort := Sort[book] // 泛型具化
	bookSort(bookshelf)    // 泛型函数调用
	fmt.Println(bookshelf)
}

// 在 Go 1.18 中，any 是 interface{}的别名，也是一个预定义标识符，使用 any 作为类型参数的约束，代表没有任何约束。
type Vector[T any] []T

func (v Vector[T]) Dump() {
	fmt.Printf("%#v\n", v)
}

func TestGeneric3(t *testing.T) {
	var iv = Vector[int]{1, 2, 3, 4}
	var sv Vector[string]
	sv = []string{"a", "b", "c", "d"}
	iv.Dump()
	sv.Dump()
}
