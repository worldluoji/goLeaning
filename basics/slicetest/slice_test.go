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
	s1 := a1[:3:3] //full slice expression，由于减少了capacity，append仍然导致内存分配
	/*
	 * a[low : high : max]
	 * constructs a slice of the same type, and with the same length and elements
	 * as the simple slice expression a[low : high].
	 * Additionally, it controls the resulting slice’s capacity by setting it to max - low.
	 * Only the first index may be omitted; it defaults to 0. After slicing the array a
	 */
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 6)
	s1[0] = 5
	t.Log(a1, len(a1), cap(a1))
	t.Log(s1, len(s1), cap(s1))

	s2 := a1[0:0]
	t.Log(s2)
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

func Test2DimensionArray(t *testing.T) {
	const m = 2
	const n = 3
	aa := [m][n]int{} // m、n必须是常量，否则会报错
	aa[1][1] = 1
	aa[0][2] = 1
	t.Log(aa)

	row := 2
	col := 3
	// 动态二维数据
	var arr [][]int
	for x := 0; x < row; x++ { //循环为一维长度
		ar := make([]int, col) //创建一个一维切片
		arr = append(arr, ar)  //把一维切片,当作一个整体传入二维切片中
	}
	arr[1][2] = 1
	t.Log(arr)
}

/*
* 排序函数如下
1. Slice() 不稳定排序

2. SliceStable() 稳定排序

3. SliceIsSorted() 判断是否已排序
结构体数据组也用同样的方式排序
*/

func TestTwoDimArraySort(t *testing.T) {
	interval := [][]int{
		{2, 2},
		{2, 3},
		{3, 3},
		{1, 3},
		{5, 7},
		{2, 2},
		{4, 6},
	}
	sort.Slice(interval, func(i, j int) bool {
		return interval[i][0] > interval[j][0] //按照每行的第一个元素排序
	})
	t.Log(interval)

	lowToHitghsorted := sort.SliceIsSorted(interval, func(i, j int) bool {
		return interval[i][0] < interval[j][0]
	})
	t.Log("是否从小到大排序：", lowToHitghsorted)

}

func TestSortString(t *testing.T) {
	var words = []string{"wo", "worl", "world", "w", "wor"}
	sort.Strings(words)
	t.Log(words)

	var word = "dcba"
	b := []byte(word)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	word = string(b)
	t.Log(word)
}
