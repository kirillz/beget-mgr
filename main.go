package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/kirillz/beget-mgr/v3/internal/ui"
)

func main() {
	a := app.New()
	w := a.NewWindow("Beget-mgr v0.0.1")
	w.Resize(fyne.NewSize(700, 400))
	w.SetMaster()

	w.SetContent(ui.Create(a, w))
	w.ShowAndRun()
}
