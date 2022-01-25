package pkg_error

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func Diff(in1 int, in2 int) error {
	if in1 == in2 {
		return nil
	}
	return errors.New("diff error")
}

func foo() {
	fmt.Println("foo...")
	bar()
}

func bar() {
	if err := Diff(3, 4); err != nil {
		fmt.Printf("%+v", err) // 关键，用%+v 和第三方库pkg/errors就能打印出错误信息的调用栈
		fmt.Println("bar failed")
		return
	}
	fmt.Println("bar success")
}

/*
* 我推荐你在应用层使用 github.com/pkg/errors 来替换官方的 error 库。
* 因为使用 pkg/errors，我们不仅能传递出标准库 error 的错误信息，还能传递出抛出 error 的堆栈信息。
 */

func TestPkgError(t *testing.T) {
	foo()
}
