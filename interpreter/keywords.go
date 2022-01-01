package interpreter

import (
	"errors"
	"strings"
)

type Keyword string
type Args []string

const (
	EQUALS Keyword = "EQUALS"
	SET Keyword = "SET"
	ACTION Keyword = "ACTION"
	GOTO Keyword = "GOTO"
	NONE Keyword = "NONE"
	PRINT Keyword = "PRINT"
	SLEEP Keyword = "SLEEP"
)
var (
	KeywordList = map[string]Keyword {
		"EQUALS": EQUALS,
		"SET": SET,
		"ACTION": ACTION,
		"GOTO": GOTO,
		"NONE": NONE,
		"SLEEP": SLEEP,
		"PRINT": PRINT,
	}
)

func getKeywordAndArgsFromLine(rawLine string) (Keyword, Args, error) {
	splits := strings.Split(rawLine, "[")
	if len(splits) < 2 {
		return "", nil, errors.New("keyword not found")
	}

	keyword, err := getKeyword(splits[0])
	if err != nil {
		return "", nil, errors.New("not a valid keyword")
	}

	args, err := SplitBetween(rawLine, splits[0] + "[", "]")
	if err != nil {
		return "", nil, errors.New("could not parse arguments")
	}

	return keyword, strings.Split(args, ","), nil
}

func getKeyword(rawKey string) (Keyword, error) {
	keyword, ok := KeywordList[rawKey]
	if !ok {
		return "", errors.New("not a valid keyword")
	}

	return keyword, nil
}
