package mapreduce

import (
	"fmt"
	"testing"
)

// 通过刚刚的一些示例，你现在应该有点明白了，Map、Reduce、Filter 只是一种控制逻辑，真正的业务逻辑是以传给它们的数据和函数来定义的。
// 这是一个很经典的“业务逻辑”和“控制逻辑”分离解耦的编程模式

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray = []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func MapIntToStr(arr []int, fn func(s int) string) []string {
	var newArray = []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}
	return sum
}

func Filter(arr []int, fn func(n int) bool) []int {
	newArray := []int{}
	for _, it := range arr {
		if fn(it) {
			newArray = append(newArray, it)
		}
	}
	return newArray
}

func TestMap(t *testing.T) {
	arr := []string{"html", "js", "css"}
	newArr := MapStrToStr(arr, func(s string) string {
		return "hello " + s
	})
	for _, val := range newArr {
		fmt.Println(val)
	}
}

func TestFilter(t *testing.T) {
	arr := []int{1, 3, 5, 6, 7, 8}
	newArr := Filter(arr, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(newArr)
}

func TestReduce(t *testing.T) {
	arr := []string{"html", "js", "css"}
	res := Reduce(arr, func(s string) int {
		return len(s)
	})
	fmt.Println(res)
}
