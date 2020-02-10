package design_pattern
import (
	"sync"
	"sync/atomic"
	"testing"
	"fmt"
)

var (
	once sync.Once
	instance *IdGenerator
)

type IdGenerator struct {
	ID int64
}

func(idGenerator *IdGenerator) GetAndIncrement() int64 {
	atomic.AddInt64(&idGenerator.ID, 1)
	return idGenerator.ID
} 

func GetInstance() *IdGenerator {
	once.Do (func() {
		fmt.Println("创建单例对象...")
		instance = &IdGenerator {
			ID:0,
		}
	})
	return instance
}

func TestSingleton(t *testing.T) {
	idGenerator1 := GetInstance()
	idGenerator2 := GetInstance()
	t.Log("它们是相同的对象吗？",idGenerator1 == idGenerator2)
}

