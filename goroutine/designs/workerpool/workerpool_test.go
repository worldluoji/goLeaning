package workerpool

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

// Task 是一个对用户提交的请求的抽象，它的本质就是一个函数类型
type Task func()

const (
	maxCapacity     = 6
	defaultCapacity = 6
)

type Pool struct {
	capacity int // workerpool大小

	active chan struct{} // 对应上图中的active channel
	tasks  chan Task     // 对应上图中的task channel

	wg   sync.WaitGroup // 用于在pool销毁时等待所有worker退出
	quit chan struct{}  // 用于通知各个worker退出的信号channel
}

/*
* workerpool.New：用于创建一个 pool 类型实例，并将 pool 池的 worker 管理机制运行起来；
* workerpool.Free：用于销毁一个 pool 池，停掉所有 pool 池中的 worker；
* Pool.Schedule：这是 Pool 类型的一个导出方法，workerpool 包的用户通过该方法向 pool 池提交待执行的任务（Task）。
 */

func (p *Pool) newWorker(i int) {
	p.wg.Add(1)
	// 由于每个 worker 运行于一个独立的 Goroutine 中，newWorker 方法通过 go 关键字创建了一个新的 Goroutine 作为 worker
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]: recover panic[%s] and exit\n", i, err)
				<-p.active
			}
			p.wg.Done()
		}()

		fmt.Printf("worker[%03d]: start\n", i)

		// 同样是一个死循环，代表这个worker在一直运行着
		for {
			select {
			case <-p.quit:
				// worker收到退出信号后，从active信道中退出，停止运行
				fmt.Printf("worker[%03d]: exit\n", i)
				<-p.active
				return
			case t := <-p.tasks:
				// 如果有任务进入，就执行
				fmt.Printf("worker[%03d]: receive a task\n", i)
				t()
			}
		}
	}()
}

// 创建一个新的 Goroutine，用于对 workerpool 进行管理
// run 方法内是一个无限循环，循环体中使用 select 监视 Pool 类型实例的两个 channel：quit 和 active
// 这种在 for 中使用 select 监视多个 channel 的实现，在 Go 代码中十分常见，是一种惯用法
func (p *Pool) run() {
	idx := 0

	for {
		select {
		case <-p.quit:
			// 当接收到来自 quit channel 的退出“信号”时，这个 Goroutine 就会结束运行
			return
		case p.active <- struct{}{}:
			// create a new worker
			// 而当 active channel 可写时，run 方法就会创建一个新的 worker Goroutine
			idx++
			p.newWorker(idx)
		}
	}
}

var ErrWorkerPoolFreed = errors.New("workerpool freed") // workerpool已终止运行

func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- t:
		// 把任务放入tasks中，一旦 p.tasks 可写，提交的 Task 就会被写入 tasks channel，以供 pool 中的 worker 处理
		// worker会取竞争这个任务，一个任务只有一个worker会执行
		return nil
	}
}

func (p *Pool) Free() {
	close(p.quit) // 关闭quit信道，所有接收者都能从中（<-p.quit）读取到。于是就都退出了。
	p.wg.Wait()
	fmt.Printf("workerpool freed\n")
}

func New(capacity int) *Pool {
	if capacity <= 0 {
		capacity = defaultCapacity
	}
	if capacity > maxCapacity {
		capacity = maxCapacity
	}

	p := &Pool{
		capacity: capacity,
		tasks:    make(chan Task),
		quit:     make(chan struct{}),
		active:   make(chan struct{}, capacity),
	}

	fmt.Printf("workerpool start\n")

	go p.run()

	return p
}

func TestWorkerPool(t *testing.T) {
	p := New(3)
	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(time.Second * 3)
		})
		if err != nil {
			println("task: ", i, "err:", err)
		}
	}
	p.Free()
}
