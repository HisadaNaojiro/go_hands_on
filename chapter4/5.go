package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hellow")
	l := widget.NewLabel("Hellow fyne!")
	e := widget.NewEntry()
	e.SetText("0")
	w.SetContent(
		container.NewVBox(
			l, e,
			widget.NewButton("Click me!", nil),
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.ShowAndRun()
}
