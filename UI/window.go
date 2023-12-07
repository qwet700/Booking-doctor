package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func RunUI() {
	myApp := app.New()
	myApp.Settings().SetTheme(newFysionTheme())
	myWindow := myApp.NewWindow("Test user")

	myWindow.SetContent(GUI())
	myWindow.Resize(fyne.NewSize(300, 400))
	myWindow.ShowAndRun()
}
