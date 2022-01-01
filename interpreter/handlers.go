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
		if val, ok := i.InterpreterTask.Storage[args[0]]; ok {
			fmt.Println(val.Get())
			return nil
		} else {
			return errors.New("no variable found: " + args[0])
		}
	}, nil
}

func (i *Interpreter) handleSet(args Args) (func() error, error) {

	if len(args) <= 1 {
		return nil, errors.New("invalid args")
	}

	return func() error {
		i.InterpreterTask.Storage[args[0]] = &InterpreterValue {
			Value: args[1],
		}
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

func (i *Interpreter) handleGoto(args Args) (func() error, error) {

	if len(args) <= 0 {
		return nil, errors.New("invalid args")
	}

	index, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.New("could not convert arg to number")
	}

	return func() error {
		if index >= len(i.InterpreterTask.Actions) {
			return errors.New("GOTO index is higher than action count")
		}
		i.InterpreterTask.LastWasGoto = true
		i.InterpreterTask.CurrentAction = index
		return nil
	}, nil
}

func (i *Interpreter) handleNone() (func() error, error) {
	return func() error {
		return nil
	}, nil
}