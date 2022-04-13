package designs

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	once     sync.Once
	instance *IDGenerator
)

type IDGenerator struct {
	ID int64
}

func (idGenerator *IDGenerator) GetAndIncrement() int64 {
	atomic.AddInt64(&idGenerator.ID, 1)
	return idGenerator.ID
}

func GetInstance() *IDGenerator {
	once.Do(func() {
		fmt.Println("创建单例对象...")
		instance = &IDGenerator{
			ID: 0,
		}
	})
	return instance
}

func TestSingleton(t *testing.T) {
	idGenerator1 := GetInstance()
	idGenerator2 := GetInstance()
	id := idGenerator2.GetAndIncrement()
	t.Log("它们是相同的对象吗？", idGenerator1 == idGenerator2)
	t.Log(id, idGenerator1.ID)
}
