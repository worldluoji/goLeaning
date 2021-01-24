package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func hello(msg string) {
	fmt.Println("Hello " + msg)
}

func TestGoroutineCase1(t *testing.T) {
	//在新的协程中执行hello方法
	go hello("World")
	//等待100毫秒让协程执行结束
	time.Sleep(100 * time.Millisecond)
}

func TestGoroutineCase2(t *testing.T) {
	start := time.Now()
	ch := make(chan int, 3)
	for i := 0; i < 6; i++ {
		go func(num int) {
			ch <- num
		}(i)
	}
	for i := 0; i < 6; i++ {
		go func(i int) {
			o := <-ch
			t.Logf("收到的第%d个元素为%d\n", i, o)
		}(i)
	}
	elapse := time.Since(start)
	t.Log("耗时时间为：", elapse)
	time.Sleep(100 * time.Millisecond)
}

func TestSelectChannelCase(t *testing.T) {
	ch := make(chan string)
	go func(chan string) {
		time.Sleep(101 * time.Millisecond)
		ch <- "data...data...data"
	}(ch)
	select {
	case msg := <-ch:
		t.Log("从ch读取到数据：", msg)
	case <-time.After(100 * time.Millisecond):
		t.Log("已超时")
		// default:
		// 	t.Log("什么也没找到")
	}
}

func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func odd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 == 1 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

// pipeline实际就是用channel将结果连接起来
func TestPipeline1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	// in fact , it is a pipeline
	for r := range odd(sq(echo(nums))) {
		fmt.Println(r)
	}
}

type EchoFunc func([]int) <-chan int
type PipelineFunc func(<-chan int) <-chan int

// 利用函数式编程对Pipeline进行封装
func Pipeline(nums []int, echoFunc EchoFunc, pipelineFuncs ...PipelineFunc) <-chan int {
	ch := echoFunc(nums)
	for _, f := range pipelineFuncs {
		ch = f(ch)
	}
	return ch
}

func TestPipeline2(t *testing.T) {
	out := Pipeline([]int{1, 2, 3, 4, 5, 6}, echo, sq, odd)
	for r := range out {
		fmt.Println(r)
	}
}
