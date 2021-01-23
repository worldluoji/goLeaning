package designs

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type ObjectPool struct {
	bufChan chan interface{}
}

func NewObjectPool(NumOfObj int, object interface{}) *ObjectPool {
	objectPool := ObjectPool{}
	objectPool.bufChan = make(chan interface{}, NumOfObj)
	for i := 0; i < NumOfObj; i++ {
		objectPool.bufChan <- object
	}
	return &objectPool
}

func (p *ObjectPool) GetObject(timeout time.Duration) (interface{}, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("获取对象超时...")
	}
}

func (p *ObjectPool) ReleaseObject(object interface{}) error {
	select {
	case p.bufChan <- object:
		return nil
	default:
		return errors.New("对象池已满...")
	}
}

// 定义一只喷火龙
type FireDragon struct {
	Name  string // 姓名
	Age   int    // 年龄
	Owner string // 主人
}

func (fireDragon *FireDragon) DragonRage() {
	fmt.Printf("%s的喷火龙%s使出龙之怒...\n", fireDragon.Owner, fireDragon.Name)
}

func UseFireDragon(objectPool *ObjectPool) {
	if object, err := objectPool.GetObject(100 * time.Millisecond); err == nil {
		if v, ok := object.(FireDragon); ok {
			fmt.Println("喷火龙，就决定是你了...")
			//dragon := (FireDragon)v
			// 使出技能龙之怒
			v.DragonRage()
			fmt.Println("回来吧，喷火龙")
			objectPool.ReleaseObject(v)
		} else {
			fmt.Println("喷火龙不听话...")
		}
	} else {
		fmt.Println(err)
	}
}

func TestObjectPool(t *testing.T) {
	fireDragon := FireDragon{
		Name:  "喷火玥",
		Owner: "luoji",
		Age:   3,
	}
	// 对象池，有3只喷火龙
	objectPool := NewObjectPool(3, fireDragon)
	// 取出一只喷火龙战斗
	for i := 0; i < 6; i++ {
		go UseFireDragon(objectPool)
	}

	time.Sleep(1000 * time.Millisecond)
}
