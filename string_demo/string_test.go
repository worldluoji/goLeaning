package string_demo

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringIterator(t *testing.T) {
	ss := "武汉加油加油！！！"
	t.Log(len(ss), len([]rune(ss)))
	for _, s := range ss {
		t.Logf("%[1]c %[1]x", s)
	}
}

/* 可见遍历时，会自动转化为rune, UTF8字符
--- PASS: TestCreateAndInitStructCase1 (0.00s)
    string_test.go:8: 武 6b66
    string_test.go:8: 汉 6c49
    string_test.go:8: 加 52a0
    string_test.go:8: 油 6cb9
    string_test.go:8: 加 52a0
    string_test.go:8: 油 6cb9
    string_test.go:8: ！ ff01
    string_test.go:8: ！ ff01
    string_test.go:8: ！ ff01
*/

func TestStringSplit(t *testing.T) {
	ss := "123,456,789"
	parts := strings.Split(ss, ",")
	for _, val := range parts {
		t.Log(val)
	}
}

func TestStringJoin(t *testing.T) {
	s1 := "123"
	s2 := "456"
	s3 := "789"
	sa := [3]string{s1, s2, s3}
	ss := strings.Join(sa[:], ",")
	t.Log(ss)
}

func TestStringConvert(t *testing.T) {
	a := 3
	s := strconv.Itoa(a)
	t.Log(s)
	if b, err := strconv.Atoi(s); err == nil {
		t.Log(b)
	}
}
