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
}

func TestSpecialOperaotr(t *testing.T) {
	a := 3
	b := 1
	c := a &^ b
	t.Log(c)
}
