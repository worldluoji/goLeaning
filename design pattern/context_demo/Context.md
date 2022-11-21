# Context
1. Context的使用场景
Context适用于以下场景：

比如一个网络请求 Request，每个 Request 都需要开启一个 goroutine 做一些事情，
这些 goroutine 又可能会开启其它的 goroutine。
所以我们需要一种可以跟踪 goroutine 的方案，才可以达到控制它们的目的，

这就是Go语言为我们提供的 Context，称之为上下文非常贴切，它就是 goroutine 的上下文。

而 WaitGroup 是一种控制并发的方式，它的这种方式是控制多个goroutine完成之后，才能执行后面的语句。

2. Context接口
```
type Context interface {
	Deadline() (deadline time.Time, ok bool)

	Done() <-chan struct{}

	Err() error

	Value(key interface{}) interface{}
}
```
Deadline 方法是获取设置的截止时间的意思，第一个返回式是截止时间，到了这个时间点，Context 会自动发起取消请求；
第二个返回值 ok==false 时表示没有设置截止时间，如果需要取消的话，需要调用取消函数进行取消。

Done 方法返回一个只读的 chan，类型为 struct{}，我们在 goroutine 中，如果该方法返回的 chan 可以读取，
则意味着 parent context 已经发起了取消请求，我们通过 Done 方法收到这个信号后，就应该做清理操作，然后退出 goroutine，释放资源。

Err 方法返回取消的错误原因，为什么 Context 被取消。

Value 方法获取该 Context 上绑定的值，是一个键值对，所以要通过一个 Key 才可以获取对应的值，这个值一般是线程安全的。

3. 经典用法
```
func Stream(ctx context.Context, out chan<- Value) error {
    for {
        v, err := DoSomething(ctx)
        if err != nil {
            return err
        }
        select {
        case <-ctx.Done():
            return ctx.Err()
        case out <- v:
        }
    }
}
```
直到 Context 被取消（cancel()）的时候，我们就可以得到一个关闭的 chan，关闭的 chan 是可以读取的，
所以只要可以读取的时候（<-ctx.Done()），就意味着收到 Context 取消的信号了。

4. Context接口实现
Context 接口并不需要我们实现，Go 内置已经帮我们实现了 2 个，
我们代码中最开始都是以这两个内置的作为最顶层的 partent context，衍生出更多的子 Context。
```
var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Context {
	return background
}

func TODO() Context {
	return todo
}
```
一个是 Background，主要用于 main 函数、初始化以及测试代码中，作为 Context 这个树结构的最顶层的 Context，也就是根 Context。

一个是 TODO，它目前还不知道具体的使用场景，如果我们不知道该使用什么 Context 的时候，可以使用这个。

它们两个本质上都是 emptyCtx 结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的 Context。
```
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}
```
这就是 emptyCtx 实现 Context 接口的方法，可以看到，这些方法什么都没做，返回的都是 nil 或者零值。


Context 的继承衍生
有了如上的根 Context，那么是如何衍生更多的子 Context 的呢？这就要靠 context 包为我们提供的 With 系列的函数了。
```
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```
这四个 With 函数，接收的都有一个 partent 参数，就是父 Context，
我们要基于这个父 Context 创建出子 Context 的意思，这种方式可以理解为子 Context 对父 Context 的继承，
也可以理解为基于父 Context 的衍生。

通过这些函数，就创建了一颗 Context 树，树的每个节点都可以有任意多个子节点，节点层级可以有任意多个。
- WithCancel 函数，传递一个父 Context 作为参数，返回子 Context，以及一个取消函数用来取消 Context。 
- WithDeadline 函数，和 WithCancel 差不多，它会多传递一个截止时间参数，意味着到了这个时间点，会自动取消 Context，当然我们也可以不等到这个时候，可以提前通过取消函数进行取消。
- WithTimeout 和 WithDeadline 基本上一样，这个表示是超时自动取消，是多少时间后自动取消 Context 的意思。
- WithValue 函数和取消 Context 无关，它是为了生成一个绑定了一个键值对数据的 Context，
这个绑定的数据可以通过 Context.Value 方法访问到。

大家可能留意到，前三个函数都返回一个取消函数 CancelFunc，这是一个函数类型，它的定义非常简单。
```
type CancelFunc func()
```
这就是取消函数的类型，该函数可以取消一个 Context，以及这个节点 Context下所有的所有的 Context，不管有多少层级。

5.  Context 使用原则
- 不要把 Context 放在结构体中，要以参数的方式传递
- 以 Context 作为参数的函数方法，应该把 Context 作为第一个参数，放在第一位。
- 给一个函数方法传递 Context 的时候，不要传递 nil，如果不知道传递什么，就使用 context.TODO
- Context 的 Value 相关方法应该传递必须的数据，不要什么数据都使用这个传递
- Context 是线程安全的，可以放心的在多个 goroutine 中传递