package context_demo

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	/*
	* context.Background() 返回一个空的 Context，这个空的 Context 一般用于整个 Context 树的根节点。
	* 然后我们使用 context.WithCancel(parent) 函数，创建一个可取消的子 Context，
	* 然后当作参数传给 goroutine 使用，这样就可以使用这个子 Context 跟踪这个 goroutine。
	 */
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				t.Log("监控退出，停止了...")
				return
			default:
				t.Log("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	t.Log("可以了，通知监控停止")

	/*
	* 那么是如何发送结束指令的呢？这就是示例中的 cancel 函数啦，
	* 它是我们调用context.WithCancel(parent) 函数生成子 Context 的时候返回的，
	* 第二个返回值就是这个取消函数，它是 CancelFunc 类型的。
	* 我们调用它就可以发出取消指令，然后我们的监控 goroutine 就会收到信号(<-ctx.Done())，就会返回结束。
	 */
	cancel()

	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

/*
* 示例中启动了 3 个监控 goroutine 进行不断的监控，每一个都使用了 Context 进行跟踪，
* 当我们使用 cancel 函数通知取消时，这 3 个 goroutine 都会被结束。
* 这就是 Context 的控制能力，它就像一个控制器一样，按下开关后，
* 所有基于这个 Context 或者衍生的子 Context 都会收到通知，这时就可以进行清理操作了，
* 最终释放 goroutine，这就优雅的解决了 goroutine 启动后不可控的问题。
 */
func TestContextDemo2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "Wather1")
	go watch(ctx, "Wather2")
	go watch(ctx, "Wather3")
	time.Sleep(10 * time.Second)
	t.Log("可以了，通知监控停止")
	cancel()
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

var key string = "name"

func TestContextDemo3(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	valueCtx := context.WithValue(ctx, key, "【监控1】")
	go watchWithValue(valueCtx)
	time.Sleep(10 * time.Second)
	t.Log("可以了，通知监控停止")
	cancel()
	time.Sleep(5 * time.Second)
}

func watchWithValue(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
