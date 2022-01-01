package interpreter

import (
	"fmt"
	"reflect"
)

type InterpreterTask struct {
	Actions []func()error
	CurrentAction int
	LastWasGoto bool
	Storage map[string]*InterpreterValue
}

func (it *InterpreterTask) Run() error {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Script fatally crashed")
		}
	}()

	it.CurrentAction = 0

	for it.CurrentAction < len(it.Actions) {

		funcA := it.Actions[it.CurrentAction]
		err := funcA()
		if err != nil {
			return err
		}

		if !it.LastWasGoto {
			it.CurrentAction++
		} else {
			it.LastWasGoto = false
		}
	}

	return nil
}

type InterpreterValue struct {
	Value interface{}
}

func (c *InterpreterValue) Compare(cVal interface{}) bool {
	return reflect.DeepEqual(c.Value, cVal)
}

func (c *InterpreterValue) Get() interface{} {
	return c.Value
}

func (c *InterpreterValue) Set(val interface{}) {
	c.Value = val
}

