package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

/*
flag.Type(flag 名, 默认值, 帮助信息) *Type
Type 可以是 Int、String、Bool 等，返回值为一个相应类型的指针，例如我们要定义姓名、年龄、婚否三个命令行参数，我们可以按如下方式定义：
name := flag.String("name", "张三", "姓名")
age := flag.Int("age", 18, "年龄")
married := flag.Bool("married", false, "婚否")
delay := flag.Duration("d", 0, "时间间隔")
*/
var cpuprofile = flag.String("cpuprofile", "D:\\go\\tmp\\cpuprofile\\test.txt", "write cpu profile to file")

/*
* flag.TypeVar(Type 指针, flag 名, 默认值, 帮助信息)
* TypeVar 可以是 IntVar、StringVar、BoolVar 等，其功能为将 flag 绑定到一个变量上
var name string
var age int
var married bool
var delay time.Duration
flag.StringVar(&name, "name", "张三", "姓名")
flag.IntVar(&age, "age", 18, "年龄")
flag.BoolVar(&married, "married", false, "婚否")
flag.DurationVar(&delay, "d", 0, "时间间隔")
*/
var good bool
var bad bool

func init() {
	log.Println("init:", good, bad)
	flag.BoolVar(&good, "good", true, "好不好")
	flag.BoolVar(&bad, "bad", true, "坏不坏")
}

/*
* CPU Profiling: 在代码中添加 CPUProfile 代码，runtime/pprof 包提供支持, 用于性能分析
* go run .\profiling.go -good -bad
* go run .\profiling.go -good
* go run .\profiling.go
 */
func main() {
	// 通过调用 flag.Parse() 来对命令行参数进行解析
	flag.Parse()
	log.Println(*cpuprofile)
	log.Println(flag.Args())  // Args returns the non-flag command-line arguments, 对于后面没有-符号制定的参数，flag统统归为non-flags，可以使用flag.Args()使用获取
	log.Println(flag.NArg())  // Args的数量
	log.Println(flag.NFlag()) // 返回使用的命令行参数个数
	log.Println("main", good, bad)
	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	var result int
	for i := 0; i < 100000000; i++ {
		result += i
	}
	log.Println("result:", result)
}
