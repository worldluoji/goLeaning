package if_demo

import "testing"

func TestIfCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		if j := 3 * i; j%2 == 0 {
			t.Log(i)
		}
	}
}

// 这个例子中switch相当于if
func TestSwitchCase1(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log(i, "是一个偶数")
		default:
			t.Log(i, "是一个奇数")
		}
	}
}

func TestIfCase2(t *testing.T) {
	var f interface{}
	f = 3.25
	if v, ok := f.(float64); ok {
		t.Log(v, "是float类型")
	} else {
		t.Log(v, "是其它类型")
	}
}

func case1() int {
	println("eval case1 expr")
	return 1
}

func case2_1() int {
	println("eval case2_1 expr")
	return 0
}
func case2_2() int {
	println("eval case2_2 expr")
	return 2
}

func case3() int {
	println("eval case3 expr")
	return 3
}

func switchexpr() int {
	println("eval switch expr")
	return 2
}

/*
* 从输出结果中我们看到，Go 先对 switch expr 表达式进行求值，然后再按 case 语句的出现顺序，
  从上到下进行逐一求值。在带有表达式列表的 case 语句中，Go 会从左到右，对列表中的表达式进行求值，比如示例中的 case2_1 函数就执行于 case2_2 函数之前。
* 如果 switch 表达式匹配到了某个 case 表达式，那么程序就会执行这个 case 对应的代码分支，
  比如示例中的“exec case2”。这个分支后面的 case 表达式将不会再得到求值机会，比如示例不会执行 case3 函数。
  这里要注意一点，即便后面的 case 表达式求值后也能与 switch 表达式匹配上，Go 也不会继续去对这些表达式进行求值了。
* 除了这一点外，你还要注意 default 分支。无论 default 分支出现在什么位置，它都只会在所有 case 都没有匹配上的情况下才会被执行的。
*/
func TestSwitchCase2(t *testing.T) {
	switch switchexpr() {
	case case1():
		println("exec case1")
	case case2_1(), case2_2():
		println("exec case2")
	case case3():
		println("exec case3")
	default:
		println("exec default")
	}
}

// go的case不需要break退出，相反，使用fallthrough可以不退出，继续执行下一个case
func TestSwitchCase3(t *testing.T) {
	switch switchexpr() {
	case case2_2():
		println("exec case1")
		fallthrough
	case case1():
		println("exec case2")
		fallthrough
	default:
		println("exec default")
	}
}
