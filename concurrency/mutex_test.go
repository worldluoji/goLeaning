package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

func TestAddWithoutMutex(t *testing.T) {
	var count = 0
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				// count++ 不是一个原子操作，多个goroutine读取到count的当前值可能一样
				count++
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

func TestAddWithMutex(t *testing.T) {
	var count = 0
	// Mutex 的零值是还没有 goroutine 等待的未加锁的状态，所以你不需要额外的初始化，直接声明变量（如 var mu sync.Mutex）
	var mu sync.Mutex
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

// 在初始化嵌入的 struct 时，也不必初始化这个 Mutex 字段，不会因为没有初始化出现空指针或者是无法获取到锁的情况。
type Counter struct {
	mu    sync.Mutex
	Count uint64
}

func TestMutexEmbeded(t *testing.T) {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.mu.Lock()
				counter.Count++
				counter.mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count)
}
