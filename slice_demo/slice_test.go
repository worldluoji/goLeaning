package slice_demo

import "testing"

func TestArrayToSlice1(t *testing.T) {
	a1 := [...]int{1, 2, 3}
	t.Log(a1, len(a1), cap(a1))
	s1 := a1[:]
	s1[0] = 5 // 范围内修改，共享内存的
	t.Log(a1, len(a1), cap(a1))
}

func TestArrayToSlice2(t *testing.T) {
	a1 := [...]int{1, 2, 3}
	t.Log(a1, len(a1), cap(a1))
	s1 := a1[:]
	s1 = append(s1, 6) // cap不够时，append将导致内存重新分配，不会再共享内存
	s1[0] = 5
	t.Log(a1, len(a1), cap(a1))
	t.Log(s1, len(s1), cap(s1))
}

// 重点注意一下，append()这个函数在 cap 不够用的时候，就会重新分配内存以扩大容量
// 如果够用，就不会重新分配内存了
func TestArrayToSlice3(t *testing.T) {
	a1 := make([]int, 3, 4)
	a1[0] = 1
	a1[1] = 2
	a1[2] = 3
	t.Log(a1, len(a1), cap(a1))
	s1 := a1[:]
	s1 = append(s1, 6) // cap足够，append不会导致内存重新分配
	s1[0] = 5
	t.Log(a1, len(a1), cap(a1))
	t.Log(s1, len(s1), cap(s1))
}

// full slice expression
func TestArrayToSlice4(t *testing.T) {
	a1 := make([]int, 3, 4)
	a1[0] = 1
	a1[1] = 2
	a1[2] = 3
	t.Log(a1, len(a1), cap(a1))
	s1 := a1[:3:3] //full slice expression，append仍然导致内存分配
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 6)
	s1[0] = 5
	t.Log(a1, len(a1), cap(a1))
	t.Log(s1, len(s1), cap(s1))
}
