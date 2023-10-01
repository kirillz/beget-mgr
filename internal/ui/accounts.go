package ui

import "fyne.io/fyne/v2"

func newAccounts(a fyne.App, w fyne.Window) *settings {
	return &settings{app: a, window: w}
}
