package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func RunUI() {
	myApp := app.New()
	myApp.Settings().SetTheme(newFysionTheme())
	defaultWindow := myApp.NewWindow("Welcome to app")
	defaultWindow.SetMaster()
	defaultWindow.Show()

	userWindow := myApp.NewWindow("Welcome to app")
	userWindow.SetContent(GUI())
	userWindow.Resize(fyne.NewSize(300, 400))
	userWindow.Show()

	myApp.Run()
}
