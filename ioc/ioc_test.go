package ioc

import (
	"errors"
	"testing"
)

/*
** 控制反转，不是由控制逻辑 Undo  来依赖业务逻辑 IntSet，而是由业务逻辑 IntSet 依赖 Undo 。
** 这里依赖的是其实是一个协议，这个协议是一个没有参数的函数数组。可以看到，这样一来，我们 Undo 的代码就可以复用了。
** 它的主要思想是把控制逻辑与业务逻辑分开，不要在业务逻辑里写控制逻辑，因为这样会让控制逻辑依赖于业务逻辑，
** 而是反过来，让业务逻辑依赖控制逻辑。
 */
type Undo []func()

type Inset struct {
	data map[int]bool
	undo Undo
}

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

// Undo last step
func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("find no function to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil
	}
	return nil
}

func (set *Inset) NewInSet() Inset {
	return Inset{data: make(map[int]bool)}
}

func (set *Inset) Undo() error {
	return set.undo.Undo()
}

func (set *Inset) Contains(x int) bool {
	return set.data[x]
}

func (set *Inset) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() {
			set.Delete(x)
		})
	} else {
		set.undo.Add(nil)
	}
}

func (set *Inset) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() {
			set.Add(x)
		})
	} else {
		set.undo.Add(nil)
	}
}

func TestIoc(t *testing.T) {
	var set Inset
	set = set.NewInSet()
	set.Add(1)
	set.Add(2)
	set.undo.Undo()
	t.Log(set.Contains(1), set.Contains(2))
}
