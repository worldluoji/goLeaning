package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

var ip = pflag.StringP("ip", "i", "0.0.0.0", "help message for ip")

// 给变量赋值
var good bool

func init() {
	pflag.BoolVar(&good, "good", false, "help message for goodVar")
}

// 如果一个 Flag 具有 NoOptDefVal，并且该 Flag 在命令行上没有设置这个 Flag 的值，则该标志将设置为 NoOptDefVal 指定的值
func init() {
	pflag.Lookup("ip").NoOptDefVal = "127.0.0.1"
}

var flagSet *pflag.FlagSet

// 自定义 FlagSet
func init() {
	var version bool
	flagSet = pflag.NewFlagSet("test", pflag.ContinueOnError)
	flagSet.BoolVar(&version, "version", true, "Print version information and quit.")
}

// go run main.go -> 0.0.0.0
// go run main.go --ip  -> 127.0.0.1
// go run main.go --ip=1.1.1.1  -> 1.1.1.1
// go run main.go --good
func main() {
	pflag.Parse()
	fmt.Println(*ip)
	fmt.Println(good)

	if v, err := flagSet.GetBool("version"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

}
