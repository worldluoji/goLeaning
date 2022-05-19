# 使用 Cobra 库创建命令
如果要用 Cobra 库编码实现一个应用程序，需要首先创建一个空的 main.go 文件和一个 rootCmd 文件，之后可以根据需要添加其他命令。Cobra可以很好的和viper,pflag结合在一起。

## 依赖包
go mod tidy cobra

## 编译
go build -v .

## 运行
- windows .\cobra.exe -h
- Linux cobra -h
- cobra version
- cobra -v --region=chengdu
- cobra version --author=luoji

## 总结
- 在开发 Go 项目时，我们可以通过 Pflag 来解析命令行参数，通过 Viper 来解析配置文件，用 Cobra 来实现命令行框架。

## 内置验证函数
- NoArgs：如果存在任何非选项参数，该命令将报错。
- ArbitraryArgs：该命令将接受任何非选项参数。
- OnlyValidArgs：如果有任何非选项参数不在 Command 的 ValidArgs 字段中，该命令将报错。
- MinimumNArgs(int)：如果没有至少 N 个非选项参数，该命令将报错。
- MaximumNArgs(int)：如果有多于 N 个非选项参数，该命令将报错。
- ExactArgs(int)：如果非选项参数个数不为 N，该命令将报错。
- ExactValidArgs(int)：如果非选项参数的个数不为 N，或者非选项参数不在 Command 的 ValidArgs 字段中，该命令将报错。
- RangeArgs(min, max)：如果非选项参数的个数不在 min 和 max 之间，该命令将报错。

<br>

### 使用预定义验证函数，示例如下：
```
var cmd = &cobra.Command{
  Short: "hello",
  Args: cobra.MinimumNArgs(1), // 使用内置的验证函数
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hello, World!")
  },
}
```

### 自定义验证函数:

```
var cmd = &cobra.Command{
  Short: "hello",
  // Args: cobra.MinimumNArgs(10), // 使用内置的验证函数
  Args: func(cmd *cobra.Command, args []string) error { // 自定义验证函数
    if len(args) < 1 {
      return errors.New("requires at least one arg")
    }
    if myapp.IsValidColor(args[0]) {
      return nil
    }
    return fmt.Errorf("invalid color specified: %s", args[0])
  },
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hello, World!")
  },
}
```
## PreRun and PostRun Hooks
在运行 Run 函数时，我们可以运行一些钩子函数，比如 PersistentPreRun 和 PreRun 函数在 Run 函数之前执行，PersistentPostRun 和 PostRun 在 Run 函数之后执行。如果子命令没有指定Persistent*Run函数，则子命令将会继承父命令的Persistent*Run函数。这些函数的运行顺序如下：
1. PersistentPreRun
2. PreRun
3. Run
4. PostRun
5. PersistentPostRun

```
注意，父级的 PreRun 只会在父级命令运行时调用，子命令是不会调用的。
Cobra 还支持很多其他有用的特性，比如：自定义 Help 命令；可以自动添加--version标志，输出程序版本信息；当用户提供无效标志或无效命令时，Cobra 可以打印出 usage 信息；当我们输入的命令有误时，Cobra 会根据注册的命令，推算出可能的命令，等等。
```