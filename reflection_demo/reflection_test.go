package reflection_demo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var a int64 = 128
	t.Log(reflect.TypeOf(a), reflect.ValueOf(a))
	t.Log(reflect.ValueOf(a).Type())
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Int64, reflect.Int32, reflect.Int:
		fmt.Println("这是一个整数")
	case reflect.Float64, reflect.Float32:
		fmt.Println("这是一个浮点数")
	case reflect.String:
		fmt.Println("这是一个字符串")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestCheckType(t *testing.T) {
	var a int = 63
	var f float64 = 3.158
	var s string = "zzy"
	CheckType(a)
	CheckType(f)
	CheckType(s)
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestDeepEqual(t *testing.T) {
	m1 := map[int]string{1: "aaa", 2: "bbb", 3: "ccc"}
	m2 := map[int]string{1: "aaa", 2: "bbb", 3: "ccc"}
	//t.Log(m1 == m2)  map只能和nil做比较
	t.Log(reflect.DeepEqual(m1, m2))

	// 数组和列表元素和顺序都一样，才先等
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	t.Log("c1 ==c2?", c1 == c2)
	t.Log("c1 deep equal c2?", reflect.DeepEqual(c1, c2))
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"377766", "luoji", 29}
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
	nameField, ok := reflect.TypeOf(*e).FieldByName("Name")
	t.Logf("NameField: value(%[1]v), Type(%[1]T) ", nameField)
	if ok {
		t.Log(nameField.Tag.Get("format"))
	} else {
		t.Error("Fail to get Field")
	}

	// method,ok := reflect.ValueOf(e).MethodByName("UpdateAge")
	// if ok {
	// 	method.Call([]reflect.Value{reflect.ValueOf(18)})
	// } else {
	// 	t.Error("Fail to get Method")
	// }
	method := reflect.ValueOf(e).MethodByName("UpdateAge")
	method.Call([]reflect.Value{reflect.ValueOf(18)})
	t.Log("Age: ", e.Age)

}
