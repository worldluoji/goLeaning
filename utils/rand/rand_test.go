package other_demo

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandInt(t *testing.T) {
	var se int64 = time.Now().UnixNano()
	t.Log(se)
	rand.Seed(se) // 设置随机种子，否则每次运行结果会一样, 只要se变化了，随机序列就会变
	for i := 0; i < 3; i++ {
		num := rand.Intn(100) // [0,100)
		key := fmt.Sprintf("stu%d", num)
		t.Log(key)
	}

}
