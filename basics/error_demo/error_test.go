package error_demo

import (
	"errors"
	"fmt"
	"testing"
)

// 构造error的两种方式
// err := errors.New("your first demo error")
// errWithCtx = fmt.Errorf("index %d is out of bounds", i)

var ErrorDevideZero = errors.New("除数不能为0")

func devide(a, b int) (int, error) {
	if b == 0 {
		return -1, ErrorDevideZero
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

// $GOROOT/src/bufio/bufio.go
var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)

var ErrSentinel = errors.New("the underlying sentinel error")

// 如果 error 类型变量的底层错误值是一个包装错误（Wrapped Error），errors.Is 方法会沿着该包装错误所在错误链（Error Chain)，与链上所有被包装的错误（Wrapped Error）进行比较，直至找到一个匹配的错误为止
func TestErrorCase2(t *testing.T) {
	err1 := fmt.Errorf("wrap sentinel: %w", ErrSentinel)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	println(err2 == ErrSentinel) //false
	if errors.Is(err2, ErrSentinel) {
		println("err2 is ErrSentinel")
		return
	}

	println("err2 is not ErrSentinel")
}

type MyError struct {
	e string
}

func (e *MyError) Error() string {
	return e.e
}

// 从 Go 1.13 版本开始，标准库 errors 包提供了As函数给错误处理方检视错误值。
// As函数类似于通过类型断言判断一个 error 类型变量是否为特定的自定义错误类型
func TestErrorCase3(t *testing.T) {
	var err = &MyError{"MyError error demo"}
	err1 := fmt.Errorf("wrap err: %w", err)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	var e *MyError
	if errors.As(err2, &e) {
		println("MyError is on the chain of err2")
		// errors.As函数沿着 err2 所在错误链向下找到了被包装到最深处的错误值，并将 err2 与其类型 * MyError成功匹配。匹配成功后，errors.As 会将匹配到的错误值存储到 As 函数的第二个参数中，这也是为什么println(e == err)输出 true 的原因
		println(e == err)
		return
	}
	println("MyError is not on the chain of err2")
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
