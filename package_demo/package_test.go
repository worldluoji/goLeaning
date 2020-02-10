package package_demo

import (
	"fmt"
	"testing"
	"github.com/luoji_demo/package_demo/hello" // 自定义包，要从src下面那一层开始导入
	cm "github.com/easierway/concurrent_map"
) 


func init() {
	fmt.Println("package init1...")
}

func init() {
	fmt.Println("package init2...")
}

func TestPackageImportCase1(t *testing.T) {
	hello.Hello()
	// 只有首字母大写的方法才能被调用
}


func TestPackageImportCase2(t *testing.T) {
	m := cm.CreateConcurrentMap(10)
	m.Set(cm.StrKey("a"), 100)
	t.Log(m.Get(cm.StrKey("a")))
}