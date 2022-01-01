package main

import "VenariBot/interpreter"

func main() {

	script := []string {
		"SET[test,1]",
		"PRINT[test]",
		"SLEEP[1000]",
		"NONE[]",
		"SET[test2,22]",
		"PRINT[test2]",
		"GOTO[0]",
	}
	i := interpreter.NewInterpreter(script)
	err := i.Parse()
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
