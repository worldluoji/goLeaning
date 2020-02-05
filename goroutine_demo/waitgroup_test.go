package goroutine_demo

import (
	"testing"
	"sync"
)

func TestWaitGroupCase1(t *testing.T) {
	var wg sync.WaitGroup

	ch := make(chan int, 3)
	for i := 0; i < 6; i++ {
		go func(num int, wg *sync.WaitGroup) {
			ch <- num
			wg.Add(1)
		} (i, &wg)
	}
	for i := 0; i < 6; i++ {
		go func(i int, wg *sync.WaitGroup) {
			o := <-ch
			t.Logf("收到的第%d个元素为%d\n", i, o)
			wg.Done()
		} (i, &wg)
	}
	wg.Wait()
}