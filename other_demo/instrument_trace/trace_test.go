package instrument_trace

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
)

/**
* defer 关键字后面只能接函数（或方法），这些函数被称为 deferred 函数。
* defer 将它们注册到其所在 Goroutine 中，用于存放 deferred 函数的栈数据结构中，
* 这些 deferred 函数将在执行 defer 的函数退出前，按后进先出（LIFO）的顺序被程序调度执行
* 无论是执行到函数体尾部返回，还是在某个错误处理分支显式 return，又或是出现 panic，已经存储到 deferred 函数栈中的函数，都会被调度执行。
* 所以说，deferred 函数是一个可以在任何情况下为函数进行收尾工作的好“伙伴”。
**/
func TestDeferTrace(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go foo(&wg)
	go foo(&wg)
	wg.Wait()
}

/*
* Go 会在 defer 设置 deferred 函数时对 defer 后面的表达式进行求值。
* 以 foo 函数中的defer Trace("foo")()这行代码为例，Go 会对 defer 后面的表达式Trace("foo")()进行求值。
* 由于这个表达式包含一个函数调用Trace("foo")，所以这个函数会被执行。
 */
func foo(wg *sync.WaitGroup) {
	defer Trace()()
	bar()
	wg.Done()
}

func bar() {
	defer Trace()()
}

var mu sync.Mutex
var m = make(map[uint64]int)

/*
* 我们通过 runtime.Caller 函数获得当前 Goroutine 的函数调用栈上的信息，
* runtime.Caller 的参数标识的是要获取的是哪一个栈帧的信息。
* 当参数为 0 时，返回的是 Caller 函数的调用者的函数信息，在这里就是 Trace 函数。
* 但我们需要的是 Trace 函数的调用者的信息，于是我们传入 1。
*
* Caller 函数有四个返回值：第一个返回值代表的是程序计数（pc）；
* 第二个和第三个参数代表对应函数所在的源文件名以及所在行数，这里我们暂时不需要；
* 最后一个参数代表是否能成功获取这些信息
*
* runtime.FuncForPC 返回的名称中不仅仅包含函数名，还包含了被跟踪函数所在的包名
 */
func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	gid := curGoroutineID()

	mu.Lock()
	ind := m[gid]
	// 每进一层就+1
	m[gid] = ind + 1
	mu.Unlock()
	printTrace(gid, name, "->", ind+1)

	return func() {
		mu.Lock()
		ind := m[gid]
		// 每出一层就-1
		m[gid] = ind - 1
		mu.Unlock()
		printTrace(gid, name, "<-", ind)
	}
}

var goroutineSpace = []byte("goroutine ")

//  获取goroutine ID的方法
func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

func printTrace(gid uint64, name, arrow string, indent int) {
	indents := ""
	for i := 0; i < indent; i++ {
		indents += "	"
	}
	fmt.Printf("g[%05d]%s%s[%s]\n", gid, indents, arrow, name)
}
