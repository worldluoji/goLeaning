package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroupCase1(t *testing.T) {
	var wg sync.WaitGroup

	ch := make(chan int, 3)
	for i := 0; i < 6; i++ {
		go func(num int, wg *sync.WaitGroup) {
			ch <- num
			wg.Add(1)
		}(i, &wg)
	}
	for i := 0; i < 6; i++ {
		go func(i int, wg *sync.WaitGroup) {
			o := <-ch
			t.Logf("收到的第%d个元素为%d\n", i, o)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}

func SimpleProducer(ch chan<- int, count int, wg *sync.WaitGroup) {
	// wg.Add(1)
	for i := 0; i < count; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func SimpleReceiver(ch <-chan int, count int, wg *sync.WaitGroup) {
	// wg.Add(1)
	for i := 0; i < count; i++ {
		if o, ok := <-ch; ok {
			fmt.Printf("收到的第%d个元素为%d\n", i, o)
		} else {
			fmt.Println("信道关闭退出")
			break
		}
	}
	wg.Done()
}

func TestMultyReceiver(t *testing.T) {
	ch := make(chan int, 3)
	var wg sync.WaitGroup
	wg.Add(3)
	// 一个生产者，两个消费者
	go SimpleProducer(ch, 13, &wg)
	go SimpleReceiver(ch, 8, &wg)
	go SimpleReceiver(ch, 8, &wg)
	wg.Wait()
	t.Log("程序所有协程全部执行完毕")
}
