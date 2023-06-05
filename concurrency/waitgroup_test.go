package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroupCase1(t *testing.T) {
	var wg sync.WaitGroup

	ch := make(chan int, 3)
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(num int) {
			ch <- num
		}(i)
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
	for i := 0; i < count; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func SimpleReceiver(ch <-chan int, count int, wg *sync.WaitGroup) {
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

type counter struct {
	c chan int
	i int
}

func NewCounter() *counter {
	cter := &counter{
		c: make(chan int),
	}
	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()
	return cter
}

func (cter *counter) Increase() int {
	return <-cter.c
}

/*
* 利用无缓冲信道，发一定在收之前，否则收阻塞的特性，实现一个counter
 */
func TestWaitGroupCase2(t *testing.T) {
	cter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := cter.Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(cter.i)
}
