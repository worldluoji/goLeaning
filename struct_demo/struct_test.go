package struct_demo

import "testing"

type User struct {
	Name    string
	Age     int
	address string
}

func TestCreateAndInitStructCase1(t *testing.T) {
	var user User = User{
		Name:    "luoji",
		Age:     29,
		address: "Chengdu",
	}
	t.Log(user.Name, user.Age, user.address)
}

func TestCreateAndInitStructCase2(t *testing.T) {
	var user *User = &User{
		Name:    "luoji",
		Age:     29,
		address: "Chengdu",
	}
	t.Log(user.Name, user.Age, user.address)
}

func TestCreateAndInitStructCase3(t *testing.T) {
	user := new(User)
	user.Name = "luoji"
	user.Age = 29
	user.address = "Chendu"
	t.Log(user.Name, user.Age, user.address)
}
