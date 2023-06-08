package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

/*
* 读写锁的性能随着并发量增大的情况，与前面讲解的 sync.RWMutex 一致；
* 利用原子操作的无锁并发写的性能，随着并发量增大几乎保持恒定；
* 利用原子操作的无锁并发读的性能，随着并发量增大有持续提升的趋势，并且性能是读锁的约 200 倍。
 */

var n1 int64

func addSyncByAtomic(delta int64) int64 {
	return atomic.AddInt64(&n1, delta)
}

func readSyncByAtomic() int64 {
	return atomic.LoadInt64(&n1)
}

var n2 int64
var rwmu sync.RWMutex

func addSyncByRWMutex(delta int64) {
	rwmu.Lock()
	n2 += delta
	rwmu.Unlock()
}

func readSyncByRWMutex() int64 {
	var n int64
	rwmu.RLock()
	n = n2
	rwmu.RUnlock()
	return n
}

func BenchmarkAddSyncByAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			addSyncByAtomic(1)
		}
	})
}

func BenchmarkReadSyncByAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			readSyncByAtomic()
		}
	})
}

func BenchmarkAddSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			addSyncByRWMutex(1)
		}
	})
}

func BenchmarkReadSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			readSyncByRWMutex()
		}
	})
}

type Config struct {
	NodeName string
	Addr     string
	Count    int32
}

func loadNewConfig() Config {
	return Config{
		NodeName: "Chengdu",
		Addr:     "1.2.3.4",
		Count:    rand.Int31(),
	}
}

func TestAtomicValue(t *testing.T) {
	var config atomic.Value
	config.Store(loadNewConfig())
	var cond = sync.NewCond(&sync.Mutex{})

	// 设置新的config
	go func() {
		for {
			time.Sleep(time.Duration(5+rand.Int63n(5)) * time.Second)
			config.Store(loadNewConfig())
			cond.Broadcast() // 通知等待着配置已变更
		}
	}()

	go func() {
		for {
			cond.L.Lock()
			cond.Wait()                 // 等待变更信号
			c := config.Load().(Config) // 读取新的配置
			fmt.Printf("new config: %+v\n", c)
			cond.L.Unlock()
		}
	}()

	select {}
}
