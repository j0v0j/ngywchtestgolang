package fyne02

import (
	"fmt"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func Fyne02() {
	app := app.New()

	w := app.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("点击", func() {
			fmt.Println("点击了")
		}),
	))
	w.ShowAndRun()
}
