package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GUI() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {}),
	)

	left := widget.NewLabel("left")
	right := widget.NewLabel("right")

	content := widget.NewLabel("content")
	content.Alignment = fyne.TextAlignCenter

	return container.NewBorder(toolbar, nil, left, right, content)
}
