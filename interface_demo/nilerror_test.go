package interface_demo

import (
	"errors"
	"fmt"
	"testing"
)

type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad things happened"),
}

func bad() bool {
	return false
}

/*
* 修改方法：修改方法:
1. 把returnsError()里面p的类型改为error
2. 删除p，直接return &ErrBad或者nil
*/
func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}

func TestNilError(t *testing.T) {
	err := returnsError()
	if err != nil {
		// error occur: <nil>  为什么nil 不等于 nil ?
		fmt.Printf("error occur: %+v\n", err)
	} else {
		fmt.Println("ok")
	}
}

func TestNilInterface(t *testing.T) {
	printNilInterface()
	fmt.Println("******************************************")
	printEmptyInterface()
	fmt.Println("******************************************")
	printNonEmptyInterface()
	fmt.Println("******************************************")
	printEmptyInterfaceAndNonEmptyInterface()
}

func printNilInterface() {
	// nil接口变量
	var i interface{} // 空接口类型
	var err error     // 非空接口类型
	println(i)
	println(err)
	println("i = nil:", i == nil)
	println("err = nil:", err == nil)
	println("i = err:", i == err)
}

/*
* 从输出结果中我们可以总结一下：对于空接口类型变量，只有 _type 和 data 所指数据内容一致的情况下，
* 两个空接口类型变量之间才能划等号
 */
func printEmptyInterface() {
	var eif1 interface{} // 空接口类型
	var eif2 interface{} // 空接口类型
	var n, m int = 17, 18

	eif1 = n
	eif2 = m

	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2) // false

	eif2 = 17
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2) // true

	eif2 = int64(17)
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2) // false
}

type T int

func (t T) Error() string {
	return "bad error"
}

func printNonEmptyInterface() {
	var err1 error   // 非空接口类型
	var err2 error   // 非空接口类型
	err1 = (*T)(nil) // 针对这种赋值，println 输出的 err1 是（0x10ed120, 0x0），也就是非空接口类型变量的类型信息并不为空，数据指针为空，因此它与 nil（0x0,0x0）之间不能划等号
	println("err1:", err1)
	println("err1 = nil:", err1 == nil)

	err1 = T(5)
	err2 = T(6)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)

	err2 = fmt.Errorf("%d\n", 5)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)
}

func printEmptyInterfaceAndNonEmptyInterface() {
	var eif interface{} = T(5)
	var err error = T(5)
	println("eif:", eif)
	println("err:", err)
	// 空接口类型变量和非空接口类型变量内部表示的结构有所不同（第一个字段：_type vs. tab)，两者似乎一定不能相等。
	// 但 Go 在进行等值比较时，类型比较使用的是 eface 的 _type 和 iface 的 tab._type，因此就像我们在这个例子中看到的那样，当 eif 和 err 都被赋值为T(5)时，两者之间是划等号的
	println("eif = err:", eif == err)

	err = T(6)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err)
}

/*
eface 用于表示没有方法的空接口（empty interface）类型变量，也就是 interface{}类型的变量；
iface 用于表示其余拥有方法的接口 interface 类型变量。

// $GOROOT/src/runtime/runtime2.go

// 对于空interface来说，_type和data都需要一样才相等
type eface struct {
    _type *_type
    data  unsafe.Pointer
}

// 对于非空interface来说，data和tab都要一样才相等，tab中又包括了_type, 接口的其它信息
type iface struct {
    tab  *itab
    data unsafe.Pointer
}


// $GOROOT/src/runtime/runtime2.go
type itab struct {
    inter *interfacetype
    _type *_type
    hash  uint32 // copy of _type.hash. Used for type switches.
    _     [4]byte
    fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}

*/
