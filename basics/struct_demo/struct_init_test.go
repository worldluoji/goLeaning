package struct_demo

import (
	"fmt"
	"testing"
)

type Foo struct {
	name string
	id   int
	age  int

	db interface{}
}

// FooOption 代表可选参数
type FooOption func(foo *Foo)

// WithName 代表Name为可选参数
func WithName(name string) FooOption {
	return func(foo *Foo) {
		foo.name = name
	}
}

// WithAge 代表age为可选参数
func WithAge(age int) FooOption {
	return func(foo *Foo) {
		foo.age = age
	}
}

// WithDB 代表db为可选参数
func WithDB(db interface{}) FooOption {
	return func(foo *Foo) {
		foo.db = db
	}
}

// NewFoo 代表初始化
func NewFoo(id int, options ...FooOption) *Foo {
	foo := &Foo{
		name: "default",
		id:   id,
		age:  10,
		db:   nil,
	}
	for _, option := range options {
		option(foo)
	}
	return foo
}

// Go 中的方法必须是归属于一个类型的，而 receiver 参数的类型就是这个方法归属的类型，或者说这个方法就是这个类型的一个方法
// 这里foo就是receiver, 除了 receiver 参数名字要保证唯一外，Go 语言对 receiver 参数的基类型也有约束，那就是 receiver 参数的基类型本身不能为指针类型或接口类型。
// Go 对方法声明的位置也是有约束的，Go 要求，方法声明要与 receiver 参数的基类型声明放在同一个包内, 所以，我们不能为原生类型（诸如 int、float64、map 等）添加方法，也不能跨越 Go 包为其他包的类型声明新方法。
// Go 语言中的方法的本质就是，一个以方法的 receiver 参数作为第一个参数的普通函数。
func (foo *Foo) PrintInfo() {
	fmt.Printf("name:%s,age:%d\n", foo.name, foo.age)
}

func TestInitStructWithOptions(t *testing.T) {
	// 这样，age和db就是默认值，初始化时不会造成语义上的误会，类似于builder模式
	foo := NewFoo(1, WithName("test"))
	foo.PrintInfo()

}
