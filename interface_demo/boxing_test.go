package interface_demo

import (
	"fmt"
	"testing"
)

type TS struct {
	n int
	s string
}

func (TS) M1() {}
func (TS) M2() {}

type NonEmptyInterface interface {
	M1()
	M2()
}

/*
* 在 Go 语言中，将任意类型赋值给一个接口类型变量也是装箱操作。
* 有了前面对接口类型变量内部表示(nilerror_test.go)的学习，我们知道接口类型的装箱实际就是创建一个 eface 或 iface 的过程。
* 装箱是有一定性能消耗的，因此Go也在不断优化
 */
func TestBoxing(t *testing.T) {
	var ts = TS{
		n: 17,
		s: "hello, interface",
	}
	var ei interface{}
	ei = ts //装箱

	var i NonEmptyInterface
	i = ts //装箱
	fmt.Println(ei)
	fmt.Println(i)
	fmt.Println(ei == i)
}
