package interface_demo

import (
	"fmt"
	"testing"
)

// Go 1.14 版本以后，Go 接口类型允许嵌入的不同接口类型的方法集合存在交集(两个Eat(string))，
// 但前提是交集中的方法不仅名字要一样，它的函数签名部分也要保持一致，也就是参数列表与返回值列表也要相同，
// 否则 Go 编译器照样会报错。
type Animal interface {
	Eat(string)
	Run()
	Sleep()
}

// 尽量定义小接口，即方法个数在 1~3 个之间的接口。Go 语言之父 Rob Pike 曾说过的“接口越大，抽象程度越弱”，这也是 Go 社区倾向定义小接口的另外一种表述。
// 定义电老鼠接口，所有电老鼠都具备这些技能
type ElectricMouse interface {
	Thuner()      // 打雷技能
	Thunderbolt() // 定义十万伏特技能
	Eat(string)
	flash() // 在 Go 接口类型的方法集合中放入首字母小写的非导出方法也是合法的,如果接口类型的方法集合中包含非导出方法，那么这个接口类型自身通常也是非导出的，它的应用范围也仅局限于包内
	Animal
}

// 定义一只皮卡丘
type PiKaChu struct {
	Name  string // 皮卡丘的名字
	Owner string // 皮卡丘的主人
}

// 定义一只雷丘， 雷丘直接复用皮卡丘技能，因为皮卡丘会的雷丘都会，所以用匿名的方式即可
type Raichu struct {
	PiKaChu
}

/*
* Go 语言中接口类型与它的实现者之间的关系是隐式的，不需要像其他语言（比如 Java）那样要求实现者显式放置“implements”进行修饰，
* 实现者只需要实现接口方法集合中的全部方法便算是遵守了契约，并立即生效了。
 */
func (piKaChu PiKaChu) Thuner() {
	fmt.Println(piKaChu.Owner + "的皮卡丘" + piKaChu.Name + "使出打雷!!!")
}

func (piKaChu PiKaChu) Thunderbolt() {
	fmt.Println(piKaChu.Owner + "的皮卡丘" + piKaChu.Name + "使出十万伏特!!!")
}

func (piKaChu PiKaChu) Eat(food string) {
	fmt.Println(piKaChu.Owner + "的皮卡丘" + piKaChu.Name + "正在吃" + food)
}

func (piKaChu PiKaChu) Run() {
	fmt.Println(piKaChu.Owner + "的皮卡丘" + piKaChu.Name + "在酷跑")
}

func (piKaChu PiKaChu) Sleep() {
	fmt.Println(piKaChu.Owner + "的皮卡丘" + piKaChu.Name + "在睡觉")
}

func (piKaChu PiKaChu) flash() {
	fmt.Println(piKaChu.Owner + "的皮卡丘" + piKaChu.Name + "闪现!!!")
}

func (raiChu Raichu) Thuner() {
	fmt.Println(raiChu.Owner + "的雷丘" + raiChu.Name + "使出打雷!!!")
}

func (raiChu Raichu) Thunderbolt() {
	fmt.Println(raiChu.Owner + "的雷丘" + raiChu.Name + "使出十万伏特!!!")
}

func TestIntrfaceCase1(t *testing.T) {
	var electricMouse ElectricMouse
	pikaChu := &PiKaChu{
		Name:  "zhangmiaomiao",
		Owner: "luoji",
	}
	electricMouse = pikaChu
	electricMouse.Thuner()
	electricMouse.Thunderbolt()
	electricMouse.Eat("fish")
	electricMouse.Sleep()
	electricMouse.flash()

	raiChu := new(Raichu)
	raiChu.Name = "雷丘"
	raiChu.Owner = "luoji"
	electricMouse = raiChu
	electricMouse.Thuner()
	electricMouse.Thunderbolt()
}
