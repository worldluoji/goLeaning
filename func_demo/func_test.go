package func_demo

import "testing"

func add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func TestUnknownCountParams(t *testing.T) {
	t.Log(add(1, 2, 3, 4, 5))
}

func forWithDefer(t *testing.T) {
	for i := 0; i < 3; i++ {
		defer func() { t.Log("第几个数？ ", i) }()
	}
}

func TestDefer(t *testing.T) {
	forWithDefer(t)
}

func foo() func() int {
	a := 1
	return func() int {
		a++
		return a
	}
}

// 包函数共享了 foo 函数的本地变量 a
func TestFoo(t *testing.T) {
	f := foo()
	f()
	f()
	f()
	t.Log(f())
}
