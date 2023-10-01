// Package ui handles all logic related to the user interface.
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// Create will set up and create the ui components.
func Create(app fyne.App, window fyne.Window) *container.AppTabs {

	return &container.AppTabs{Items: []*container.TabItem{
		newSettings(app, window).tabItem(),
		newAccounts(app, window).tabItem(),
	}}
}
