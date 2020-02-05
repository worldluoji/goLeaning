package interface_demo
import "testing"
import "fmt"

// 定义电老鼠接口，所有电老鼠都具备这些技能
type ElectricMouse interface {
	Thuner() // 打雷技能
	Thunderbolt()  // 定义十万伏特技能
}

// 定义一只皮卡丘
type PiKaChu struct {
	Name string // 皮卡丘的名字
	Owner string // 皮卡丘的主人
}

// 定义一只雷丘， 雷丘直接服用皮卡丘技能，因为皮卡丘会的雷丘都会，所以用匿名的方式即可
type Raichu struct {
	PiKaChu
}

func(piKaChu PiKaChu) Thuner() {
	fmt.Println(piKaChu.Owner + "的皮卡丘" + piKaChu.Name + "使出打雷!!!")
}

func(piKaChu PiKaChu) Thunderbolt() {
	fmt.Println(piKaChu.Owner + "的皮卡丘" + piKaChu.Name + "使出十万伏特!!!")
}

func(raiChu Raichu) Thuner() {
	fmt.Println(raiChu.Owner + "的雷丘" + raiChu.Name + "使出打雷!!!")
}

func(raiChu Raichu) Thunderbolt() {
	fmt.Println(raiChu.Owner + "的雷丘" + raiChu.Name + "使出十万伏特!!!")
}

func TestIntrfaceCase1(t *testing.T) {
	var electricMouse ElectricMouse;
	pikaChu := &PiKaChu {
		Name: "zhangmiaomiao",
		Owner: "luoji",
	}
	electricMouse = pikaChu
	electricMouse.Thuner()
	electricMouse.Thunderbolt()

	raiChu := new (Raichu)
	raiChu.Name = "雷丘"
	raiChu.Owner = "luoji"
	electricMouse = raiChu
	electricMouse.Thuner()
	electricMouse.Thunderbolt()
}