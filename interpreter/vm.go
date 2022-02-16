package interpreter

import (
	"Ven/requests"
	"Ven/strats"
	"errors"
	"fmt"
	"strconv"
)

type InterpreterTask struct {
	Actions []func()error
	Task *strats.VenariTask
	CurrentAction int
	LastWasGoto bool
	Storage map[string]*InterpreterValue
	TagStorage map[string]int
}

func (it *InterpreterTask) Run() error {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Script fatally crashed")
		}
	}()

	it.Task = &strats.VenariTask {
		Client: requests.CreateNewClient(nil),
	}
	it.Task.Login("")

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

func (it *InterpreterTask) GetInterpreterValue(key string) (*InterpreterValue, error) {
	if val, ok := it.Storage[key]; !ok {
		return nil, errors.New("variable does not exist")
	} else {
		return val, nil
	}
}

func (it *InterpreterTask) GetValue(key string) (interface{}, error) {
	if val, ok := it.Storage[key]; !ok {
		return nil, errors.New("variable does not exist")
	} else {
		return val.Get(), nil
	}
}

func (it *InterpreterTask) SetValue(key string, value interface{}) {
	it.Storage[key] = &InterpreterValue {
		Value: value,
	}
}

func (it *InterpreterTask) GetTagValue(key string) (int, error) {
	if val, ok := it.TagStorage[key]; !ok {
		return 0, errors.New("variable does not exist")
	} else {
		return val, nil
	}
}

func (it *InterpreterTask) SetTagValue(key string, val int) {
	it.TagStorage[key] = val
}

type InterpreterValue struct {
	Value interface{}
}

func (c *InterpreterValue) Equals(cVal interface{}) bool {

	strX := ""
	strY := ""

	switch c.Value.(type) {
	case string:
		strX = c.Value.(string)
	case int64:
		strX = strconv.FormatInt(c.Value.(int64), 10)
	case int:
		strX = strconv.Itoa(c.Value.(int))
	}

	switch cVal.(type) {
	case string:
		strY = cVal.(string)
	case int64:
		strY = strconv.FormatInt(cVal.(int64), 10)
	case int:
		strY = strconv.Itoa(cVal.(int))
	}

	return strX == strY
}

func (c *InterpreterValue) LessThan(cVal interface{}) bool {

	var intX int64
	var intY int64

	switch c.Value.(type) {
	case string:
		i, err := strconv.ParseInt(c.Value.(string), 10, 64)
		if err != nil {
			intX = 0
		}

		intX = i
	case int64:
		intX = c.Value.(int64)
	case int:
		intX = int64(c.Value.(int))
	}

	switch cVal.(type) {
	case string:
		i, err := strconv.ParseInt(cVal.(string), 10, 64)
		if err != nil {
			intX = 0
		}

		intY = i
	case int64:
		intY = cVal.(int64)
	case int:
		intY = int64(cVal.(int))
	}

	return intX < intY
}

func (c *InterpreterValue) GreaterThan(cVal interface{}) bool {

	var intX int64
	var intY int64

	switch c.Value.(type) {
	case string:
		i, err := strconv.ParseInt(c.Value.(string), 10, 64)
		if err != nil {
			intX = 0
		}

		intX = i
	case int64:
		intX = c.Value.(int64)
	case int:
		intX = int64(c.Value.(int))
	}

	switch cVal.(type) {
	case string:
		i, err := strconv.ParseInt(cVal.(string), 10, 64)
		if err != nil {
			intX = 0
		}

		intY = i
	case int64:
		intY = cVal.(int64)
	case int:
		intY = int64(cVal.(int))
	}

	return intX > intY
}

func (c *InterpreterValue) LessThanOrEqual(cVal interface{}) bool {

	val := c.LessThan(cVal)
	if !val {
		return c.Equals(cVal)
	}
	return val
}

func (c *InterpreterValue) GreaterThanOrEqual(cVal interface{}) bool {

	val := c.GreaterThan(cVal)
	if !val {
		return c.Equals(cVal)
	}
	return val
}

func (c *InterpreterValue) Get() interface{} {
	return c.Value
}

func (c *InterpreterValue) Set(val interface{}) {
	c.Value = val
}

