package designs

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"
)

/*
这个模式是将算法与操作对象的结构分离的一种方法。
这种分离的实际结果是能够在不修改结构的情况下向现有对象结构添加新操作，是遵循开放 / 封闭原则的一种方法。
其实，这段代码的目的就是想解耦数据结构和算法。虽然使用 Strategy 模式也是可以完成的，而且会比较干净.
但是在有些情况下，多个 Visitor 是来访问一个数据结构的不同部分，这种情况下，数据结构有点像一个数据库，而各个 Visitor 会成为一个个的小应用。
kubectl就是这种情况。
*/

type Visitor func(shape Shape)

type Shape interface {
	accept(Visitor)
}

type Circle struct {
	Radius int
}

type RecTangle struct {
	Length, Width int
}

// 实现了accept方法，则它（Circle）就是一个Shape
func (c Circle) accept(v Visitor) {
	v(c)
}

func (r RecTangle) accept(v Visitor) {
	v(r)
}

func JSONVisitor(shape Shape) {
	byte, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}

func XMLVisitor(shape Shape) {
	byte, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}

func TestVisitor(t *testing.T) {
	c := Circle{
		Radius: 3,
	}

	r := RecTangle{
		Length: 3,
		Width:  6,
	}

	shapes := []Shape{c, r}

	for _, shape := range shapes {
		shape.accept(JSONVisitor)
		shape.accept(XMLVisitor)
	}
}
