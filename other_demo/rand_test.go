package other_demo

import (
	"math/rand"
	"testing"
	"time"
	"fmt"
)

func TestRandInt(t *testing.T) {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子，否则每次运行结果会一样
	num := rand.Intn(100) // [0,100)
	key := fmt.Sprintf("stu%d", num) 
	t.Log(key)
}