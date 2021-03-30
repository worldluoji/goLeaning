package struct_demo

import (
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
	user.Name = "luoji"
	user.Age = 29
	user.address = "Chendu"
	t.Log(user.Name, user.Age, user.address)
}

type S struct {
	a uint16
	b uint32
}

func TestEmptyStruct(t *testing.T) {
	var s string
	var c complex128
	var a [3]uint32
	var st S
	t.Log(unsafe.Sizeof(s))  // 16
	t.Log(unsafe.Sizeof(c))  // 16
	t.Log(unsafe.Sizeof(a))  // 12 = 3 * （32 / 8）
	t.Log(unsafe.Sizeof(st)) // 8 = 16 / 8 + 16 / 8 + 32 / 8 和C语言字节对齐一样

	var st2 struct{}
	t.Log(unsafe.Sizeof(st2)) // 0 因为空结构体不占用任何空间，因此就不存在内存对齐的问题。因此多个空结构体类型组合成的结构体也不占用内存。
}
