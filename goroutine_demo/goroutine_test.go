package goroutine_demo

import (
	"testing"
	"time"
	"fmt"
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
		} (i)
	}
	for i := 0; i < 6; i++ {
		go func(i int) {
			o := <-ch
			t.Logf("收到的第%d个元素为%d\n", i, o)
		} (i)
	}
	elapse := time.Since(start)
	t.Log("耗时时间为：", elapse)
	time.Sleep(100 * time.Millisecond)
}


func TestSelectChannelCase(t *testing.T) {
	ch := make(chan string)
	timeout := make(chan string)
	timeout <- "超时，请检查信道"
	var msg string
	select {
		case msg = <-ch:
			t.Log("从ch读取到数据")
		case msg = <-timeout:
			t.Log("从timeout读取到数据")
		default:
			msg = "什么都没找到"
	}
	t.Log(msg)
}