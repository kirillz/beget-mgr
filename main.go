package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/kirillz/beget-mgr/v2/internal/assets"
	"github.com/kirillz/beget-mgr/v2/internal/ui"
)

const (
	version = "v0.0.1"
)

func main() {
	a := app.NewWithID("io.github.kirillz.beget-mgr")
	assets.SetIcon(a)
	w := a.NewWindow("Beget-mgr " + fmt.Sprint(version))
	w.SetContent(ui.Create(a, w))
	w.Resize(fyne.NewSize(500, 400))
	w.SetMaster()
	w.CenterOnScreen()
	w.ShowAndRun()
}
