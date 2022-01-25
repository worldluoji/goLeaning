package error_demo

import (
	"errors"
	"fmt"
	"testing"
)

var DevideZeroError = errors.New("除数不能为0")

func devide(a, b int) (int, error) {
	if b == 0 {
		return -1, DevideZeroError
	}
	return a / b, nil
}

func TestErrorCase1(t *testing.T) {
	if result, err := devide(3, 0); err == nil {
		t.Log("Result is", result)
	} else {
		t.Log("Error happend:", err)
	}
}

/*
* 服务可能会发生未知的panic，这时最好的方式就是使用recover兜底，防止因为报错直接退出
* 另外需要注意的是，panic不要乱用，可以明确处理返回错误的，就通过返回值判断。
 */
func TestRecoverCase1(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("错误已恢复...", err)
		}
	}()
	if result, err := devide(3, 0); err == nil {
		t.Log("Result is", result)
	} else {
		panic(err)
	}
}

func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		// 如果发生panic，仍然会执行defer中的东东
		f()
	}()
}

func TestRecoverCase2(t *testing.T) {
	// 这样进行一个封装，当函数中业务处理发生严重错误时，也不会把主线程搞挂，因为这样发生了panic，就会走到recover
	Go(func() {
		t.Log("Hello")
		// 模拟业务发生panic
		panic("业务处理发生严重错误")
	})
}

func TestPanic(t *testing.T) {
	defer func() {
		t.Log("end...")
	}()
	t.Log("start...")
	// panic，也会执行defer中的内容
	// panic("panic happens...")
	// t.Log("end...") unreachable
}
