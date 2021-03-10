package error_demo

import (
	"errors"
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

func TestPanic(t *testing.T) {
	defer func() {
		t.Log("end...")
	}()
	t.Log("start...")
	panic("panic happens...")
	// t.Log("end...") unreachable
}
