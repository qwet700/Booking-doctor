package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type fysionTheme struct {
	fyne.Theme
}

func NewFysionTheme() fyne.Theme {
	return &fysionTheme{Theme: theme.DefaultTheme()}
}

func (t *fysionTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return t.Theme.Color(name, theme.VariantLight)
}

// custome text size
func (t *fysionTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 12
	}
	return t.Theme.Size(name)
}
