package concurrency

import (
	"fmt"
	"os"
	sg "os/signal"
	"reflect"
	"sync"
	"syscall"
	"testing"
	"time"
)

// go test -v .\goroutine_test.go -run TestCachedGoroutine
// 1. 有缓冲channel，channel最多只能有C个元素，超过了入队就会阻塞
func TestCachedGoroutine(t *testing.T) {
	start := time.Now()
	ch := make(chan int, 2)
	var wg sync.WaitGroup
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(num int) {
			t.Logf("开始发送第%d个元素\n", num)
			ch <- num
		}(i)
	}

	for i := 0; i < 6; i++ {
		go func(num int) {
			o := <-ch
			t.Logf("收到的第%d个元素为%d\n", num, o)
			wg.Done()
		}(i)
	}
	wg.Wait()
	elapse := time.Since(start)
	t.Log("耗时时间为：", elapse)
}

// 2. 无缓冲channel的接收先行发生于发送完成
var cht = make(chan int)
var a string

func setVal() {
	a = "hello golang"
	cht <- 9
}

func TestNoCacheChannel(t *testing.T) {
	go setVal()
	<-cht
	// 无缓冲channel的接收先行发生于发送完成，因此能正确打印出hello golang
	fmt.Println(a)
}

// 3. 对channel的关闭先行发生于接收到值，因为channel已经被关闭了
func setVal2() {
	a = "hello golang2"
	close(cht)
}

func TestCloseChannel(t *testing.T) {
	go setVal2()
	<-cht
	fmt.Println(a)
}

/**
* 4. 当多个协程同时运行时，可通过 select 轮询多个通道，如多个通道就绪则随机选择一个
**/
func TestSelectChannelCase(t *testing.T) {
	ch := make(chan string, 1)
	ch2 := make(chan string, 1)
	go func(chan string) {
		time.Sleep(101 * time.Millisecond)
		ch <- "data...data...data"
	}(ch)

	go func(chan string) {
		time.Sleep(99 * time.Millisecond)
		ch2 <- "data2...data2...data2"
	}(ch)

	select {
	case msg := <-ch:
		t.Log("从ch读取到数据：", msg)
	case msg := <-ch2:
		t.Log("从ch2读取到数据：", msg)
	case <-time.After(100 * time.Millisecond):
		t.Log("已超时")
		// default:
		// 	t.Log("什么也没找到")
	}
}

func TestSelectHeartBeat(t *testing.T) {
	ch := make(chan string)
	stop := make(chan bool)
	go func(chan string) {
		time.Sleep(100 * time.Millisecond)
		ch <- "data...data...data"
	}(ch)

	go func(chan bool) {
		time.Sleep(10 * time.Second)
		stop <- true
	}(stop)

	heartbeat := time.NewTicker(3000 * time.Millisecond)
	defer heartbeat.Stop()
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				ch = nil
			} else {
				fmt.Println(msg)
			}
		case <-heartbeat.C:
			fmt.Println("heart beat...")
		case <-stop:
			fmt.Println("stop work...")
			return
		}
	}
}

func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		// 发送到信道完了以后，要关闭信道，否则接收方会阻塞；你关闭信道后，接收方仍然可以从信道中取到数据
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

func TestNilChannel(t *testing.T) {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- 5
		close(ch1)
	}()

	go func() {
		time.Sleep(time.Second * 7)
		ch2 <- 7
		close(ch2)
	}()

	var ok1, ok2 bool
	for {
		select {
		// 由于5S后 ，ch1 处于关闭状态，从这个 channel 获取数据，我们会得到这个 channel 对应类型的零值，
		// 这里就是 0。于是程序再次输出 0；程序按这个逻辑循环执行，一直输出 0 值
		case x := <-ch1:
			ok1 = true
			fmt.Println(x)
		case x := <-ch2:
			ok2 = true
			fmt.Println(x)
		}

		if ok1 && ok2 {
			break
		}
	}
	fmt.Println("program end")
}

/*
* 使用nil channel解决上述问题
 */
func TestNilChannel2(t *testing.T) {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- 5
		close(ch1)
	}()

	go func() {
		time.Sleep(time.Second * 7)
		ch2 <- 7
		close(ch2)
	}()

	for {
		select {
		case x, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				fmt.Println(x)
			}
		case x, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				fmt.Println(x)
			}
		}
		if ch1 == nil && ch2 == nil {
			break
		}
	}
	fmt.Println("program end")
}

func TestEmptyChannel(t *testing.T) {
	var sem chan struct{}
	t.Log(sem == nil, len(sem))
}

func TestProcessMutiChan(t *testing.T) {
	var ch1 = make(chan int, 10)
	var ch2 = make(chan int, 10)

	// 创建SelectCase
	var cases = createCases(ch1, ch2)

	// 执行10次select, 第一次肯定是 send case，因为此时 chan 还没有元素，recv 还不可用。等 chan 中有了数据以后，recv case 就可以被选择了。这样，你就可以处理不定数量的 chan 了
	for i := 0; i < 10; i++ {
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() { // recv case
			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
		} else { // send case
			fmt.Println("send:", cases[chosen].Dir, ok)
		}
	}
}

/**
* createCases 函数分别为每个 chan 生成了 recv case 和 send case，并返回一个 reflect.SelectCase 数组
* reflect.Select函数接收selectCase列表作为参数，直到列表中某个case发生才返回
* Go 的 select 是伪随机的，它可以在执行的 case 中随机选择一个 case，并把选择的这个 case 的索引（chosen）返回，如果没有可用的 case 返回，会返回一个 bool 类型的返回值，
* 这个返回值用来表示是否有 case 成功被选择。
* 如果是 recv case，还会返回接收的元素。
* func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)
**/
func createCases(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// 创建recv case
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	// 创建send case
	for i, ch := range chs {
		v := reflect.ValueOf(i)
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,  // direction of case
			Chan: reflect.ValueOf(ch), // channel to use (for send or receive)
			Send: v,                   // value to send (for send)
		})
	}

	return cases
}

// 信号通知，优雅退出
// go test -v ./goroutine_test.go -run TestNotify
func TestNotify(t *testing.T) {
	var closing = make(chan struct{})
	var closed = make(chan struct{})

	go func() {
		// 模拟业务处理
		for {
			select {
			case <-closing:
				return
			default:
				// ....... 业务计算
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 处理CTRL+C等中断信号
	termChan := make(chan os.Signal)
	sg.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	close(closing)
	// 执行退出之前的清理动作
	go doCleanup(closed)

	select {
	case <-closed:
	case <-time.After(time.Second):
		fmt.Println("清理超时，不等了")
	}
	fmt.Println("优雅退出")
}

func doCleanup(closed chan struct{}) {
	time.Sleep(20 * time.Second)
	// close一定发生在接受前，所以退出前doCleanup一定执行过了。
	close(closed)
}
