# 命令行参数解析工具：Pflag 使用介绍

虽然 Go 源码中提供了一个标准库 Flag 包，用来对命令行参数进行解析，但在大型项目中应用更广泛的是另外一个包：Pflag。Pflag 提供了很多强大的特性，非常适合用来构建大型项目，一些耳熟能详的开源项目都是用 Pflag 来进行命令行参数解析的，例如：Kubernetes、Istio、Helm、Docker、Etcd 等。<br>

Pflag 主要是通过创建 Flag 和 FlagSet 来使用。

## 1. Pflag
Pflag 可以对命令行参数进行处理，一个命令行参数在 Pflag 包中会解析为一个 Flag 类型的变量。
Flag 是一个结构体:
```
type Flag struct {
    Name                string // flag长选项的名称。例如--good
    Shorthand           string // flag短选项的名称，一个缩写的字符, 例如-g
    Usage               string // flag的使用文本
    Value               Value  // flag的值
    DefValue            string // flag的默认值
    Changed             bool // 记录flag的值是否有被设置过
    NoOptDefVal         string // 当flag出现在命令行，但是没有指定选项值时的默认值
    Deprecated          string // 记录该flag是否被放弃
    Hidden              bool // 如果值为true，则从help/usage输出信息中隐藏该flag
    ShorthandDeprecated string // 如果flag的短选项被废弃，当使用flag的短选项时打印该信息
    Annotations         map[string][]string // 给flag设置注解
}
```
Flag 的值是一个 Value 类型的接口：
```
type Value interface {
    String() string // 将flag类型的值转换为string类型的值，并返回string的内容
    Set(string) error // 将string类型的值转换为flag类型的值，转换失败报错
    Type() string // 返回flag的类型，例如：string、int、ip等
}
```
通过将 Flag 的值抽象成一个 interface 接口，我们就可以自定义 Flag 的类型了。
只要实现了 Value 接口的结构体，就是一个新类型。

<br>

### Pflag 支持以下 4 种命令行参数定义方式
支持长选项、默认值和使用文本，并将标志的值存储在指针中。
```
var name = pflag.String("name", "colin", "Input Your Name")
```
支持长选项、短选项、默认值和使用文本，并将标志的值存储在指针中。
```
var name = pflag.StringP("name", "n", "colin", "Input Your Name")
```
支持长选项、默认值和使用文本，并将标志的值绑定到变量。
```
var name string
pflag.StringVar(&name, "name", "colin", "Input Your Name")
```
支持长选项、短选项、默认值和使用文本，并将标志的值绑定到变量
```
var name string
pflag.StringVarP(&name, "name", "n","colin", "Input Your Name")
```

规则：
- 函数名带Var说明是将标志的值绑定到变量，否则是将标志的值存储在指针中。
- 函数名带P说明支持短选项，否则不支持短选项。

<br>

## 2. Pflag 包的 FlagSet
FlagSet 是一些预先定义好的 Flag 的集合，几乎所有的 Pflag 操作，都需要借助 FlagSet 提供的方法来完成。
可以使用两种方法来获取并使用 FlagSet：
- 方法一，调用 NewFlagSet 创建一个 FlagSet。
```
var version bool
flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
flagSet.BoolVar(&version, "version", true, "Print version information and quit.")
```
- 方法二，使用 Pflag 包定义的全局 FlagSet：CommandLine。实际上 CommandLine 是一个包级别的变量,也是由 NewFlagSet 函数创建的。
```
import (
    "github.com/spf13/pflag"
)

pflag.BoolVarP(&version, "version", "v", true, "Print version information and quit.")
```