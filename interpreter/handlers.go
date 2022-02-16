package interpreter

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func (i *Interpreter) handlePrint(args Args) (func() error, error) {

	if len(args) <= 0 {
		return nil, errors.New("invalid args")
	}

	return func() error {
		val, err := i.InterpreterTask.GetValue(args[0])
		if err != nil {
			fmt.Println(args[0])
		} else {
			fmt.Println(val)
		}
		return nil
	}, nil
}

func (i *Interpreter) handleSet(args Args) (func() error, error) {

	if len(args) <= 1 {
		return nil, errors.New("invalid args")
	}

	return func() error {
		i.InterpreterTask.SetValue(args[0], args[1])
		return nil
	}, nil
}

func (i *Interpreter) handleSleep(args Args) (func() error, error) {

	if len(args) <= 0 {
		return nil, errors.New("invalid args")
	}

	sleepTime, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.New("could not convert arg to number")
	}

	return func() error {
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		return nil
	}, nil
}

func (i *Interpreter) handleAdd(args Args) (func() error, error) {

	if len(args) <= 1 {
		return nil, errors.New("invalid args")
	}

	addAmount, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.New("could not convert arg to number")
	}

	return func() error {

		current, err := i.InterpreterTask.GetValue(args[0])
		if err != nil {
			return err
		}

		switch current.(type) {
		case int:
			i.InterpreterTask.SetValue(args[0], current.(int) + addAmount)
		case int64:
			i.InterpreterTask.SetValue(args[0], current.(int64) + int64(addAmount))
		case string:
			currentAmount, err := strconv.Atoi(current.(string))
			if err != nil {
				return errors.New("could not convert arg to number")
			}

			i.InterpreterTask.SetValue(args[0], strconv.Itoa(currentAmount + addAmount))
		}

		return nil
	}, nil
}

func (i *Interpreter) handleGoto(args Args) (func() error, error) {

	if len(args) <= 0 {
		return nil, errors.New("invalid args")
	}

	return func() error {

		var index int

		tagIndex, err := i.InterpreterTask.GetTagValue(args[0])
		if err != nil {
			index, err = strconv.Atoi(args[0])
			if err != nil {
				return errors.New("could not convert arg to number: " + args[0])
			}
		} else {
			index = tagIndex + 1
		}
		index = index - 1

		if index >= len(i.InterpreterTask.Actions) || index < 0 {
			return errors.New("GOTO index is higher than action count or lower than 0")
		}
		i.InterpreterTask.LastWasGoto = true
		i.InterpreterTask.CurrentAction = index
		return nil
	}, nil
}


func (i *Interpreter) handleTime(args Args) (func() error, error) {

	if len(args) <= 0 {
		return nil, errors.New("invalid args")
	}

	return func() error {
		i.InterpreterTask.SetValue(args[0], time.Now().Unix())
		return nil
	}, nil
}

func (i *Interpreter) handleAction(args Args) (func() error, error) {

	if len(args) <= 1 {
		return nil, errors.New("invalid args")
	}

	action, ok := ActionList[args[0]]
	if !ok {
		return nil, errors.New("invalid action")
	}

	return func() error {

		err := action(i.InterpreterTask, i.InterpreterTask.Task, args[1:])
		if err != nil {
			fmt.Printf("action error: %v\n", err)
			//TODO: handle error case
		}

		return nil
	}, nil
}

func (i *Interpreter) handleEnd() (func() error, error) {

	return func() error {

		i.InterpreterTask.CurrentAction = len(i.InterpreterTask.Actions)
		i.InterpreterTask.LastWasGoto = true

		return nil
	}, nil
}

func (i *Interpreter) handleNone() (func() error, error) {
	return func() error {
		return nil
	}, nil
}

func (i *Interpreter) getComparitiveAction(keyword Keyword, args Args) (func() error, error) {
	var ifcaseaction func() error
	var err error
	switch keyword {
	case PRINT:
		ifcaseaction, err = i.handlePrint(args)
		break
	case SET:
		ifcaseaction, err = i.handleSet(args)
		break
	case SLEEP:
		ifcaseaction, err = i.handleSleep(args)
		break
	case GOTO:
		ifcaseaction, err = i.handleGoto(args)
		break
	case NONE:
		ifcaseaction, err = i.handleNone()
		break
	case TIME:
		ifcaseaction, err = i.handleTime(args)
		break
	case END:
		ifcaseaction, err = i.handleEnd()
		break
	case ACTION:
		ifcaseaction, err = i.handleAction(args)
		break
	case ADD:
		ifcaseaction, err = i.handleAdd(args)
		break
	}
	return ifcaseaction, err
}

func (i *Interpreter) handleEmpty(args Args) (func() error, error) {

	if len(args) <= 2 {
		return nil, errors.New("invalid args")
	}

	keyword, ifArgs, err := getKeywordAndArgsFromLine(args[1])
	if err != nil {
		return nil, fmt.Errorf("error parsing keywords from EQUALS: %s %v", args[2], err)
	}
	ifHandleErr := keyword != ACTION
	ifCase, err := i.getComparitiveAction(keyword, ifArgs)
	if err != nil {
		return nil, fmt.Errorf("error parsing ifCase from EQUALS: %s %v", args[2], err)
	}

	keyword, elseArgs, err := getKeywordAndArgsFromLine(args[2])
	if err != nil {
		return nil, fmt.Errorf("error parsing keywords from EQUALS: %s %v", args[2], err)
	}
	elseHandleErr := keyword != ACTION
	elseCase, err := i.getComparitiveAction(keyword, elseArgs)
	if err != nil {
		return nil, fmt.Errorf("error parsing elseCase from EQUALS: %s %v", args[2], err)
	}

	return func() error {

		xVal, err := i.InterpreterTask.GetInterpreterValue(args[0])
		if err != nil {
			return err
		}

		if xVal.Equals("") {
			err = ifCase()
			if err != nil && ifHandleErr {
				return err
			}
		} else {
			err = elseCase()
			if err != nil && elseHandleErr {
				return err
			}
		}

		return nil
	}, nil
}

func (i *Interpreter) handleEquals(args Args) (func() error, error) {

	if len(args) <= 3 {
		return nil, errors.New("invalid args")
	}

	keyword, ifArgs, err := getKeywordAndArgsFromLine(args[2])
	if err != nil {
		return nil, fmt.Errorf("error parsing keywords from EQUALS: %s %v", args[2], err)
	}
	ifHandleErr := keyword != ACTION
	ifCase, err := i.getComparitiveAction(keyword, ifArgs)
	if err != nil {
		return nil, fmt.Errorf("error parsing ifCase from EQUALS: %s %v", args[2], err)
	}

	keyword, elseArgs, err := getKeywordAndArgsFromLine(args[3])
	if err != nil {
		return nil, fmt.Errorf("error parsing keywords from EQUALS: %s %v", args[2], err)
	}
	elseHandleErr := keyword != ACTION
	elseCase, err := i.getComparitiveAction(keyword, elseArgs)
	if err != nil {
		return nil, fmt.Errorf("error parsing elseCase from EQUALS: %s %v", args[2], err)
	}

	return func() error {

		xVal, err := i.InterpreterTask.GetInterpreterValue(args[0])
		if err != nil {
			return err
		}

		yVal, err := i.InterpreterTask.GetInterpreterValue(args[1])
		if err != nil {
			return err
		}

		if xVal.Equals(yVal.Get()) {
			err = ifCase()
			if err != nil && ifHandleErr {
				return err
			}
		} else {
			err = elseCase()
			if err != nil && elseHandleErr {
				return err
			}
		}

		return nil
	}, nil
}

func (i *Interpreter) handleLessThan(args Args) (func() error, error) {

	if len(args) <= 3 {
		return nil, errors.New("invalid args")
	}

	keyword, ifArgs, err := getKeywordAndArgsFromLine(args[2])
	if err != nil {
		return nil, fmt.Errorf("error parsing keywords from EQUALS: %s %v", args[2], err)
	}
	ifHandleErr := keyword != ACTION
	ifCase, err := i.getComparitiveAction(keyword, ifArgs)
	if err != nil {
		return nil, fmt.Errorf("error parsing ifCase from EQUALS: %s %v", args[2], err)
	}

	keyword, elseArgs, err := getKeywordAndArgsFromLine(args[3])
	if err != nil {
		return nil, fmt.Errorf("error parsing keywords from EQUALS: %s %v", args[2], err)
	}
	elseHandleErr := keyword != ACTION
	elseCase, err := i.getComparitiveAction(keyword, elseArgs)
	if err != nil {
		return nil, fmt.Errorf("error parsing elseCase from EQUALS: %s %v", args[2], err)
	}

	return func() error {

		xVal, err := i.InterpreterTask.GetInterpreterValue(args[0])
		if err != nil {
			return err
		}

		yVal, err := i.InterpreterTask.GetInterpreterValue(args[1])
		if err != nil {
			return err
		}

		if xVal.LessThan(yVal.Get()) {
			err = ifCase()
			if err != nil && ifHandleErr {
				return err
			}
		} else {
			err = elseCase()
			if err != nil && elseHandleErr {
				return err
			}
		}

		return nil
	}, nil
}

func (i *Interpreter) handleGreaterThan(args Args) (func() error, error) {

	if len(args) <= 3 {
		return nil, errors.New("invalid args")
	}

	keyword, ifArgs, err := getKeywordAndArgsFromLine(args[2])
	if err != nil {
		return nil, fmt.Errorf("error parsing keywords from EQUALS: %s %v", args[2], err)
	}
	ifHandleErr := keyword != ACTION
	ifCase, err := i.getComparitiveAction(keyword, ifArgs)
	if err != nil {
		return nil, fmt.Errorf("error parsing ifCase from EQUALS: %s %v", args[2], err)
	}

	keyword, elseArgs, err := getKeywordAndArgsFromLine(args[3])
	if err != nil {
		return nil, fmt.Errorf("error parsing keywords from EQUALS: %s %v", args[2], err)
	}
	elseHandleErr := keyword != ACTION
	elseCase, err := i.getComparitiveAction(keyword, elseArgs)
	if err != nil {
		return nil, fmt.Errorf("error parsing elseCase from EQUALS: %s %v", args[2], err)
	}

	return func() error {

		xVal, err := i.InterpreterTask.GetInterpreterValue(args[0])
		if err != nil {
			return err
		}

		yVal, err := i.InterpreterTask.GetInterpreterValue(args[1])
		if err != nil {
			return err
		}

		if xVal.GreaterThan(yVal.Get()) {
			err = ifCase()
			if err != nil && ifHandleErr {
				return err
			}
		} else {
			err = elseCase()
			if err != nil && elseHandleErr {
				return err
			}
		}

		return nil
	}, nil
}

func (i *Interpreter) handleLessThanOrEquals(args Args) (func() error, error) {

	if len(args) <= 3 {
		return nil, errors.New("invalid args")
	}

	keyword, ifArgs, err := getKeywordAndArgsFromLine(args[2])
	if err != nil {
		return nil, fmt.Errorf("error parsing keywords from EQUALS: %s %v", args[2], err)
	}
	ifHandleErr := keyword != ACTION
	ifCase, err := i.getComparitiveAction(keyword, ifArgs)
	if err != nil {
		return nil, fmt.Errorf("error parsing ifCase from EQUALS: %s %v", args[2], err)
	}

	keyword, elseArgs, err := getKeywordAndArgsFromLine(args[3])
	if err != nil {
		return nil, fmt.Errorf("error parsing keywords from EQUALS: %s %v", args[2], err)
	}
	elseHandleErr := keyword != ACTION
	elseCase, err := i.getComparitiveAction(keyword, elseArgs)
	if err != nil {
		return nil, fmt.Errorf("error parsing elseCase from EQUALS: %s %v", args[2], err)
	}

	return func() error {

		xVal, err := i.InterpreterTask.GetInterpreterValue(args[0])
		if err != nil {
			return err
		}

		yVal, err := i.InterpreterTask.GetInterpreterValue(args[1])
		if err != nil {
			return err
		}

		if xVal.LessThanOrEqual(yVal.Get()) {
			err = ifCase()
			if err != nil && ifHandleErr {
				return err
			}
		} else {
			err = elseCase()
			if err != nil && elseHandleErr {
				return err
			}
		}

		return nil
	}, nil
}

func (i *Interpreter) handleGreaterThanOrEquals(args Args) (func() error, error) {

	if len(args) <= 3 {
		return nil, errors.New("invalid args")
	}

	keyword, ifArgs, err := getKeywordAndArgsFromLine(args[2])
	if err != nil {
		return nil, fmt.Errorf("error parsing keywords from EQUALS: %s %v", args[2], err)
	}
	ifHandleErr := keyword != ACTION
	ifCase, err := i.getComparitiveAction(keyword, ifArgs)
	if err != nil {
		return nil, fmt.Errorf("error parsing ifCase from EQUALS: %s %v", args[2], err)
	}

	keyword, elseArgs, err := getKeywordAndArgsFromLine(args[3])
	if err != nil {
		return nil, fmt.Errorf("error parsing keywords from EQUALS: %s %v", args[2], err)
	}
	elseHandleErr := keyword != ACTION
	elseCase, err := i.getComparitiveAction(keyword, elseArgs)
	if err != nil {
		return nil, fmt.Errorf("error parsing elseCase from EQUALS: %s %v", args[2], err)
	}

	return func() error {

		xVal, err := i.InterpreterTask.GetInterpreterValue(args[0])
		if err != nil {
			return err
		}

		yVal, err := i.InterpreterTask.GetInterpreterValue(args[1])
		if err != nil {
			return err
		}

		if xVal.GreaterThanOrEqual(yVal.Get()) {
			err = ifCase()
			if err != nil && ifHandleErr {
				return err
			}
		} else {
			err = elseCase()
			if err != nil && elseHandleErr {
				return err
			}
		}

		return nil
	}, nil
}