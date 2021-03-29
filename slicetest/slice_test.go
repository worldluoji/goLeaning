package slicetest

import (
	"container/list"
	"fmt"
	"reflect"
	"sort"
	"testing"
)

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

func TestArrayToSliceEqual(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println(reflect.DeepEqual(s1, s2))
}

func TestArrayToSliceAppend(t *testing.T) {
	s1 := []int{1}
	s2 := append(s1, 4)
	s3 := append(s1, 5)
	s2[0] = 8
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
}

func TestGoList(t *testing.T) {
	l := list.New()
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	l.PushBack(7)
	l.PushBack(17)
	t.Log(l.Back().Value)
	t.Log(l.Front().Value)
	l.Remove(l.Back())
	t.Log(l.Back().Value)
}

func TestNSum(t *testing.T) {
	nums := []int{2, -4, -5, -2, -3, -5, 0, 4, -2}
	sort.Ints(nums)
	target := -14
	res := fourSum(&nums, target)
	fmt.Println(res)

}

func fourSum(nums *[]int, target int) [][]int {
	if len(*nums) < 4 {
		return make([][]int, 0)
	}
	sort.Ints(*nums)
	res := make([][]int, 0)
	numSelected := map[int]bool{}
	for index, num := range *nums {
		if index > 0 && num == (*nums)[index-1] {
			continue
		}
		if numSelected[num] {
			continue
		}
		t := append(append([]int{}, (*nums)[0:index]...), (*nums)[index+1:]...)
		tmp := threeSum(&t, target-num)
		for _, v := range tmp {
			if !numSelected[v[0]] && !numSelected[v[1]] && !numSelected[v[2]] {
				res = append(res, append(v, num))
			}
		}
		numSelected[num] = true
	}
	return res
}

// 找三数之和等于target的所有解，不能重复
func threeSum(nums *[]int, target int) [][]int {
	var res [][]int
	numSelected := map[int]bool{}
	for index, num := range *nums {
		if index > 0 && num == (*nums)[index-1] {
			continue
		}
		if numSelected[num] {
			continue
		}

		t := append(append([]int{}, (*nums)[0:index]...), (*nums)[index+1:]...)
		tmp := twoSum(t, target-num)
		for _, v := range tmp {
			// 关键 另外两个数都没有被选过，才可以，否则会重复。因为只要前面有一个数被选过，它就可以找到对应的解了
			if !numSelected[v[0]] && !numSelected[v[1]] {
				res = append(res, append(v, num))
				numSelected[num] = true
			}
		}
	}
	return res
}

// 找两数之和为target的所有解，不能重复
func twoSum(nums []int, target int) [][]int {
	i := 0
	j := len(nums) - 1
	var res [][]int
	for {
		for {
			if i > 0 && i <= len(nums)-1 && nums[i] == nums[i-1] {
				i++
			} else {
				break
			}
		}

		for {
			if j < len(nums)-2 && j >= 0 && nums[j] == nums[j+1] {
				j--
			} else {
				break
			}
		}

		if i >= j {
			break
		}

		total := nums[i] + nums[j]
		if total == target {
			r := []int{nums[i], nums[j]}
			res = append(res, r)
			j--
			i++
		} else if total > target {
			j--
		} else {
			i++
		}
	}
	return res
}
