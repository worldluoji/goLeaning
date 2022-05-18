package pflag

import (
	"testing"

	"github.com/spf13/pflag"
)

// 返回的是指针
var flagvar = pflag.Int("flagname", 123, "help message for flagname")

// go test -v .\pflag_test.go -run TestPflagDemo1 -args arg1 arg2
// 可以调用pflag.Parse()来解析定义的标志。
// 解析后，可通过pflag.Args()返回所有的非选项参数，通过pflag.Arg(i)返回第 i 个非选项参数。参数下标 0 到 pflag.NArg() - 1。
func TestPflagDemo1(t *testing.T) {
	t.Log(*flagvar)
	pflag.Parse()
	t.Logf("argument number is: %v\n", pflag.NArg())
	t.Logf("argument list is: %v\n", pflag.Args())
	t.Logf("the first argument is: %v\n", pflag.Arg(0))
}
