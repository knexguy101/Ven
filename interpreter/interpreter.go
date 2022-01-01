package interpreter

import (
	"fmt"
)

type Interpreter struct {
	ScriptLines []string
	InterpreterTask *InterpreterTask
}

func NewInterpreter(lines []string) *Interpreter {
	return &Interpreter {
		ScriptLines: lines,
		InterpreterTask: &InterpreterTask{
			CurrentAction: 0,
			LastWasGoto: false,
			Storage: make(map[string]*InterpreterValue),
		},
	}
}

func (i *Interpreter) Parse() error {
	for _, line := range i.ScriptLines {
		if line == "" {
			action, err := i.handleNone()
			if err != nil {
				return err
			}
			i.InterpreterTask.Actions = append(i.InterpreterTask.Actions, action)
			continue
		}

		fmt.Println(line)
		keyword, args, err := getKeywordAndArgsFromLine(line)
		if err != nil {
			return fmt.Errorf("error parsing line: %s %v", line, err)
		}

		var action func() error

		switch keyword {
		case PRINT:
			action, err = i.handlePrint(args)
			break
		case SET:
			action, err = i.handleSet(args)
			break
		case SLEEP:
			action, err = i.handleSleep(args)
			break
		case GOTO:
			action, err = i.handleGoto(args)
			break
		case NONE:
			action, err = i.handleNone()
			break
		case TIME:
			action, err = i.handleTime(args)
			break
		case ACTION:
			action, err = i.handleAction(args)
			break
		case END:
			action, err = i.handleEnd()
			break
		case EQUALS:
			action, err = i.handleEquals(args)
			break
		case ISEMPTY:
			action, err = i.handleEmpty(args)
			break
		case GREATER:
			action, err = i.handleGreaterThan(args)
			break
		case LESSER:
			action, err = i.handleLessThan(args)
			break
		case GREATEREQUALS:
			action, err = i.handleGreaterThanOrEquals(args)
			break
		case LESSEREQUALS:
			action, err = i.handleLessThanOrEquals(args)
			break
		case ADD:
			action, err = i.handleAdd(args)
			break
		}
		if err != nil {
			return err
		}

		i.InterpreterTask.Actions = append(i.InterpreterTask.Actions, action)
	}
	return nil
}