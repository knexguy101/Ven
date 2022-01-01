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
			continue
		}

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
		}
		if err != nil {
			return err
		}

		i.InterpreterTask.Actions = append(i.InterpreterTask.Actions, action)
	}
	return nil
}