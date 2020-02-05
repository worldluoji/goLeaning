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
