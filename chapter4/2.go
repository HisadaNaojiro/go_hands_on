package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hellow")
	w.SetContent(
		container.NewVBox(
			widget.NewLabel("Hello Fyne!"),
			widget.NewLabel("This is sample applications!"),
		),
	)

	w.ShowAndRun()
}
