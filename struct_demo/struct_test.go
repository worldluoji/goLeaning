package struct_demo

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"unsafe"
)

type User struct {
	Name    string
	Age     int
	address string
}

func TestCreateAndInitStructCase1(t *testing.T) {
	var user User = User{
		Name:    "luoji",
		Age:     29,
		address: "Chengdu",
	}
	t.Log(user.Name, user.Age, user.address)
}

func TestCreateAndInitStructCase2(t *testing.T) {
	var user *User = &User{
		Name:    "luoji",
		Age:     29,
		address: "Chengdu",
	}
	t.Log(user.Name, user.Age, user.address)
}

func TestCreateAndInitStructCase3(t *testing.T) {
	user := new(User)
	t.Log(user.Name, user.Age, user.address)
	user.Name = "luoji"
	user.Age = 29
	user.address = "Chendu"
	t.Log(user.Name, user.Age, user.address)
}

type S0 struct {
	a uint16
	b uint32
}

func TestEmptyStruct(t *testing.T) {
	var s string
	var c complex128
	var a [3]uint32
	var st S0
	t.Log(unsafe.Sizeof(s))  // 16
	t.Log(unsafe.Sizeof(c))  // 16
	t.Log(unsafe.Sizeof(a))  // 12 = 3 * （32 / 8）
	t.Log(unsafe.Sizeof(st)) // 8 = 16 / 8 + 16 / 8 + 32 / 8 和C语言字节对齐一样

	var st2 struct{}
	t.Log(unsafe.Sizeof(st2)) // 0 因为空结构体不占用任何空间，因此就不存在内存对齐的问题。因此多个空结构体类型组合成的结构体也不占用内存。
}

type MyInt int

func (n *MyInt) Add(m int) {
	*n = *n + MyInt(m)
}

type T struct {
	a int
	b int
}

type S struct {
	*MyInt
	T
	io.Reader
	s string
	n int
}

/*
* 嵌入字段类型的底层类型不能为指针类型。而且，嵌入字段的名字在结构体定义也必须是唯一的，
* 这也意味这如果两个类型的名字相同，它们无法同时作为嵌入字段放到同一个结构体定义中。
* 不过，这些约束你了解一下就可以了，一旦违反，Go 编译器会提示你的。
 */
func TestStructCombine(t *testing.T) {
	m := MyInt(17)
	r := strings.NewReader("hello, go")
	s := S{
		MyInt: &m,
		T: T{
			a: 1,
			b: 2,
		},
		Reader: r,
		s:      "demo",
	}

	var sl = make([]byte, len("hello, go"))
	s.Reader.Read(sl)       // 注意这里是s.Reader而不是s.io.Reader，Go语语法的规定
	fmt.Println(string(sl)) // hello, go
	s.MyInt.Add(5)
	fmt.Println(*(s.MyInt)) // 22
	fmt.Println(s.a)
}
