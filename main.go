package main

import (
	"Ven/interpreter"
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 2 {
		fmt.Println("Not enough args, try:\nven run\nven build")
		return
	}

	fileData, err := readLines(argsWithoutProg[1])
	if err != nil {
		fmt.Println("Could not find .ven file")
		return
	}

	i := interpreter.NewInterpreter(fileData)

	switch argsWithoutProg[0] {
	case "run":
		err := i.Parse()
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("Ven script built successfully")
		}

		err = i.InterpreterTask.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "build":
		err := i.Parse()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Ven script built successfully")
		}
	}
}
