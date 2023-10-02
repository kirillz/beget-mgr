// Package ui handles all logic related to the user interface.
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// Клиентское окно
type Client struct {
	app fyne.App

	// Notification holds the settings value for if we have notifications enabled or not.
	Notifications bool

	// OverwriteExisting holds the settings value for if we should overwrite already existing files.
	OverwriteExisting bool

	// DownloadPath holds the download path used for saving received files.
	DownloadPath string
}

// Create will set up and create the ui components.
func Create(app fyne.App, window fyne.Window) *container.AppTabs {
	client := NewClient(app)

	return &container.AppTabs{Items: []*container.TabItem{
		newAccounts(app, window, client).tabItem(),
		newSettings(app, window, client).tabItem(),
		newAbout(app).tabItem(),
	}}
}
