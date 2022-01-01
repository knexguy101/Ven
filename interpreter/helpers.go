package interpreter

import (
	"errors"
	"strings"
)

func SplitBetween(str, sep1, sep2 string) (string, error) {
	splits := strings.Split(str, sep1)
	if len(splits) <= 1 {
		return "", errors.New("split length on sep1 equals 1")
	}

	splits = strings.Split(splits[1], sep2)
	if len(splits) <= 1 {
		return "", errors.New("split length on sep2 equals 1")
	}

	return splits[0], nil
}
