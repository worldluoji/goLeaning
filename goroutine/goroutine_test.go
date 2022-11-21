package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func hello(msg string) {
	fmt.Println("Hello " + msg)
}

// go test -v .\goroutine_test.go -run TestGoroutineCase1
func TestGoroutineCase1(t *testing.T) {
	//在新的协程中执行hello方法
	go hello("World")
	//等待100毫秒让协程执行结束, 否则可能退出后，goroutine还没有结束
	time.Sleep(100 * time.Millisecond)
}

// 在容量为C的channel上的第k个接收先行发生于从这个channel上的第k+C次发送完成。
// 可以理解为， channel最多只能有C个元素，超过了入队就会阻塞
func TestGoroutineCase2(t *testing.T) {
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

var cht = make(chan int)
var a string

func setVal() {
	a = "hello golang"
	cht <- 9
}

func TestGoroutineCase3(t *testing.T) {
	go setVal()
	<-cht
	// 无缓冲channel的接收先行发生于发送完成，因此能正确打印出hello golang
	fmt.Println(a)
}

func setVal2() {
	a = "hello golang2"
	close(cht)
}

// 对channel的关闭先行发生于接收到值，因为channel已经被关闭了
func TestGoroutineCase4(t *testing.T) {
	go setVal2()
	<-cht
	fmt.Println(a)
}

/*
* 当多个协程同时运行时，可通过 select 轮询多个通道
• 如果所有通道都阻塞则等待，如定义了 default 则执行 default
• 如多个通道就绪则随机选择一个
*/
func TestSelectChannelCase(t *testing.T) {
	ch := make(chan string)
	ch2 := make(chan string)
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
