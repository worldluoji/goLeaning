package json_demo

import (
	"testing"
	"encoding/json"
)

// 喷火龙
type FireDragon struct {
	Name  string  // 姓名
	Age	  int     // 年龄
	Owner string  // 主人
}

func TestJsonByMarshal(t *testing.T) {
	fireDragon := FireDragon {
		Name : "喷火玥",
		Age : 10,
		Owner : "luoji",
	}
	data,err := json.Marshal(fireDragon)
	if err == nil {
		// data就是一个json序列化后的byte[]数组
		t.Log(data)
	} else {
		t.Log(err)
	}

	c := FireDragon{}
	if err = json.Unmarshal([]byte(data), &c);err == nil {
		t.Log(c.Name, c.Age, c.Owner)
	} else {
		t.Log(err)
	}
	
	
}