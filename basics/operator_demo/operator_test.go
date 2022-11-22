package operator_demo

import "testing"

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Readable, Writeable, Executable)
}

func TestStringCompare(t *testing.T) {
	s1 := "123"
	s2 := "123"
	t.Log(s1 == s2)

	s3 := s2
	s3 += "4"
	s1 += "4"
	t.Log(s1 == s3) // go字符串可以直接用==比较
}

func TestSpecialOperaotr(t *testing.T) {
	a := 3
	b := 1
	c := a &^ b // 与非
	t.Log(c)
}

func TestBinary(t *testing.T) {
	b := 0700      // 八进制，以"0"为前缀
	c1 := 0xaabbcc // 十六进制，以"0x"为前缀
	c2 := 0xddeeff // 十六进制，以"0X"为前缀
	t.Log(c1, c2)

	// Go 又增加了对二进制字面值的支持和两种八进制字面值的形式
	d1 := 0b10000001 // 二进制，以"0b"为前缀
	d2 := 0b10000001 // 二进制，以"0B"为前缀
	e1 := 0o700      // 八进制，以"0o"为前缀
	e2 := 0o700      // 八进制，以"0O"为s前缀
	t.Log(b == e1, e1 == e2)

	// 为提升字面值的可读性，Go 1.13 版本还支持在字面值中增加数字分隔符“_”，分隔符可以用来将数字分组以提高可读性
	d3 := 0b1000_0001
	t.Log(d1 == d2, d2 == d3)
}
