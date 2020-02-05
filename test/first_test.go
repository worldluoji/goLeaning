package test

import "testing"

func TestFirst(t *testing.T) {
	t.Log("My First Test")
}

// go test -v xxx_test.go 才能输出 t.Log 里的文字
