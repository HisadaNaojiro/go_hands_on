package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
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
			widget.NewButton("Click me!", func() {
				n, _ := strconv.Atoi(e.Text)
				l.SetText("Total: " + strconv.Itoa(total(n)))
			}),
		),
	)

	w.ShowAndRun()
}

func total(n int) int {
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	return t
}
