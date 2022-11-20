package pointer

import (
	"testing"
	"unsafe"
)

/* 不过 Go 中也有一种指针类型是例外，它不需要基类型，它就是 unsafe.Pointer。
unsafe.Pointer 类似于 C 语言中的 void*，用于表示一个通用指针类型，
也就是任何指针类型都可以显式转换为一个 unsafe.Pointer，而 unsafe.Pointer 也可以显式转换为任意指针类型
*/
func TestPointerDemo1(t *testing.T) {
	var p *int
	t.Log(p == nil) // true

	var a int = 13
	p = &a // 给整型指针变量p赋初值
	t.Log(*p)
}

type foo struct {
	id   string
	age  int8
	addr string
}

// 不同类型的指针类型的大小在同一个平台上是一致的。在 x86-64 平台上，地址的长度都是 8 个字节。
// func Sizeof(x ArbitraryType) uintptr , 在 Go 语言中 uintptr 类型的大小就代表了指针类型的大小。
func TestPointerDemo2(t *testing.T) {
	var p1 *int
	var p2 *bool
	var p3 *byte
	var p4 *[20]int
	var p5 *foo
	var p6 unsafe.Pointer
	t.Log(unsafe.Sizeof(p1)) // 8
	t.Log(unsafe.Sizeof(p2)) // 8
	t.Log(unsafe.Sizeof(p3)) // 8
	t.Log(unsafe.Sizeof(p4)) // 8
	t.Log(unsafe.Sizeof(p5)) // 8
	t.Log(unsafe.Sizeof(p6)) // 8
}

// 指针转换，如果我们使用 unsafe 包中类型或函数，代码的安全性就要由开发人员自己保证，也就是开发人员得明确知道自己在做啥！
func TestPointerDemo3(t *testing.T) {
	var a int = 0x12345678
	var pa *int = &a
	var pb *byte = (*byte)(unsafe.Pointer(pa)) // ok
	t.Logf("%x\n", *pb)                        // 78
}

/*
 unsafe.Pointer 与 uintptr 的相互转换，间接实现了“指针运算”。
 但即便我们可以使用 unsafe 方法实现“指针运算”，Go 编译器也不会为开发人员提供任何帮助，
 开发人员需要自己告诉编译器要加减的绝对地址偏移值
*/
func TestPointerDemo4(t *testing.T) {
	var arr = [5]int{11, 12, 13, 14, 15}
	var p *int = &arr[0]
	var i uintptr
	for i = 0; i < uintptr(len(arr)); i++ {
		p1 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i*unsafe.Sizeof(*p)))
		println(*p1)
	}
}
