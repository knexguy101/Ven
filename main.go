package main

import (
	"VenariBot/interpreter"
	"bufio"
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

	script, err := readLines("./script.ven")
	i := interpreter.NewInterpreter(script)
	err = i.Parse()
	if err != nil {
		panic(err)
	}

	err = i.InterpreterTask.Run()
	if err != nil {
		panic(err)
	}

	/*myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")
	myWindow.Resize(fyne.Size{
		Width: 650,
		Height: 600,
	})

	events := container.NewVBox()
	farmers := container.NewVBox()
	accounts := container.NewVBox()
	scripts := container.NewVBox()
	settings := container.NewVBox()

	tabs := container.NewAppTabs(
		container.NewTabItem("Farmers", farmers),
		container.NewTabItem("Accounts", accounts),
		container.NewTabItem("Scripts", scripts),
		container.NewTabItem("Events", events),
		container.NewTabItem("Settings", settings),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()*/
}
