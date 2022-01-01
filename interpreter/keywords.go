package interpreter

import (
	"errors"
	"strings"
)

type Keyword string
type Args []string

const (
	EQUALS Keyword = "EQUALS"
	ISEMPTY Keyword = "ISEMPTY"
	GREATER Keyword = "GREATER"
	LESSER Keyword = "LESSER"
	GREATEREQUALS Keyword = "GREATEREQUALS"
	LESSEREQUALS Keyword = "LESSEREQUALS"
	SET Keyword = "SET"
	ACTION Keyword = "ACTION"
	GOTO Keyword = "GOTO"
	NONE Keyword = "NONE"
	PRINT Keyword = "PRINT"
	SLEEP Keyword = "SLEEP"
	TIME Keyword = "TIME"
	END Keyword = "END"
	ADD Keyword = "ADD"
)
var (
	KeywordList = map[string]Keyword {
		"EQUALS": EQUALS,
		"ISEMPTY": ISEMPTY,
		"GREATER": GREATER,
		"LESSER": LESSER,
		"LESSEREQUALS": LESSEREQUALS,
		"GREATEREQUALS": GREATEREQUALS,
		"END": END,
		"SET": SET,
		"ACTION": ACTION,
		"GOTO": GOTO,
		"NONE": NONE,
		"SLEEP": SLEEP,
		"PRINT": PRINT,
		"TIME": TIME,
		"ADD": ADD,
	}
)

func getKeywordAndArgsFromLine(rawLine string) (Keyword, Args, error) {

	splitBegin := "["
	splitEnd := "]"
	splitItem := ","
	var (
		keyword Keyword
		err error
	)

	splits := strings.Split(rawLine, "[")
	if len(splits) < 2 {

		splits = strings.Split(rawLine, "(")
		if len(splits) < 2 {
			return "", nil, errors.New("keyword not found")
		}

		keyword, err = getKeyword(splits[0])
		if err != nil {
			return "", nil, errors.New("invalid keyword")
		}

		splitBegin = "("
		splitEnd = ")"
		splitItem = "."
	} else {
		keyword, err = getKeyword(splits[0])
		if err != nil {
			splits = strings.Split(rawLine, "(")
			if len(splits) < 2 {
				return "", nil, errors.New("keyword not found")
			}

			keyword, err = getKeyword(splits[0])
			if err != nil {
				return "", nil, errors.New("invalid keyword")
			}

			splitBegin = "("
			splitEnd = ")"
			splitItem = "."
		}
	}

	args, err := SplitBetween(rawLine, splits[0] + splitBegin, splitEnd)
	if err != nil {
		return "", nil, errors.New("could not parse arguments")
	}

	return keyword, strings.Split(args, splitItem), nil
}

func getKeyword(rawKey string) (Keyword, error) {
	keyword, ok := KeywordList[rawKey]
	if !ok {
		return "", errors.New("not a valid keyword")
	}

	return keyword, nil
}
