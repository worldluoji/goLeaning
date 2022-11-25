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

// 类型约束，通过order声明的，类型只能是里面几种
type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// 与 使用 interface{} 相比，泛型版本的 maxGenerics 性能要好很多，但与原生版函数如 maxInt 等还有差距
func maxGenerics[T ordered](sl []T) T {
	if len(sl) == 0 {
		panic("slice is empty")
	}

	max := sl[0]
	for _, v := range sl[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

type myString string

func TestGeneric4(t *testing.T) {
	var m int = maxGenerics([]int{1, 2, -4, -6, 7, 0})
	fmt.Println(m)                                                           // 输出：7
	fmt.Println(maxGenerics([]string{"11", "22", "44", "66", "77", "10"}))   // 输出：77
	fmt.Println(maxGenerics([]float64{1.01, 2.02, 3.03, 5.05, 7.07, 0.01}))  // 输出：7.07
	fmt.Println(maxGenerics([]int8{1, 2, -4, -6, 7, 0}))                     // 输出：7
	fmt.Println(maxGenerics([]myString{"11", "22", "44", "66", "77", "10"})) // 输出：77
}
