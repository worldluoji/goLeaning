package pkgtest

import (
	"fmt"
	"testing"

	// 自定义和第三方包，绝对路径要从src下面那一层开始导入github.com/luoji_demo/package_demo
	cm "github.com/concurrent_map"
	"github.com/luoji_demo/pkgtest/hello"
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
