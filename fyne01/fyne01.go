package fyne01

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func Fyne01() {
	myApp := app.New()

	myWin := myApp.NewWindow("Hello")
	myWin.SetContent(widget.NewLabel("Hello Fyne!"))
	myWin.Resize(fyne.NewSize(200, 200))
	myWin.ShowAndRun()
}
