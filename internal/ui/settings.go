package ui

import (
	"errors"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	//appearance "fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/kirillz/beget-mgr/v2/internal/util"
)

type settings struct {
	downloadPathEntry *widget.Entry

	client      *Client
	preferences fyne.Preferences
	window      fyne.Window
	app         fyne.App
}

// NewClient returns a new client
func NewClient(app fyne.App) *Client {
	return &Client{app: app}
}

func newSettings(a fyne.App, w fyne.Window, c *Client) *settings {
	return &settings{app: a, window: w, client: c, preferences: a.Preferences()}
}

func (s *settings) onDownloadsPathSubmitted(path string) {
	path = filepath.Clean(path)
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		dialog.ShowInformation("Не существует", "Пожалуйста укажите существующий каталог.", s.window)
		return
	} else if err != nil {
		fyne.LogError("Ошибка проверки каталога", err)
		dialog.ShowError(err, s.window)
		return
	} else if !info.IsDir() {
		dialog.ShowInformation("Не является каталогом", "Please select a valid directory.", s.window)
		return
	}

	s.client.DownloadPath = path
	s.preferences.SetString("DownloadPath", s.client.DownloadPath)
	s.downloadPathEntry.SetText(s.client.DownloadPath)
}

func (s *settings) onDownloadsPathSelected() {
	folder := dialog.NewFolderOpen(func(folder fyne.ListableURI, err error) {
		if err != nil {
			fyne.LogError("Ошибка при выборе каталога", err)
			dialog.ShowError(err, s.window)
			return
		} else if folder == nil {
			return
		}

		s.client.DownloadPath = folder.Path()
		s.preferences.SetString("DownloadPath", s.client.DownloadPath)
		s.downloadPathEntry.SetText(s.client.DownloadPath)
	}, s.window)

	folder.Resize(util.WindowSizeToDialog(s.window.Canvas().Size()))
	folder.Show()
}

// getPreferences is used to set the preferences on startup without saving at the same time.
func (s *settings) getPreferences() {
	s.client.DownloadPath = s.preferences.StringWithFallback("DownloadPath", util.UserDownloadsFolder())
	s.downloadPathEntry.Text = s.client.DownloadPath

}

func (s *settings) buildUI() *container.Scroll {

	pathSelector := &widget.Button{Icon: theme.FolderOpenIcon(), Importance: widget.LowImportance, OnTapped: s.onDownloadsPathSelected}
	s.downloadPathEntry = &widget.Entry{Wrapping: fyne.TextWrapWord, OnSubmitted: s.onDownloadsPathSubmitted, ActionItem: pathSelector}

	s.getPreferences()

	//interfaceContainer := appearance.NewSettings().LoadAppearanceScreen(s.window)

	dataContainer := container.NewGridWithColumns(2,
		util.NewBoldLabel("Сохранить файл в:"), s.downloadPathEntry,
	)

	otherContainer := container.NewVBox(
		container.NewGridWithColumns(2,

			util.NewBoldLabel("Passphrase length"),
		),
		&widget.Accordion{Items: []*widget.AccordionItem{
			{Title: "Продвинутые настройки", Detail: container.NewGridWithColumns(2)},
		}},
	)

	return container.NewScroll(container.NewVBox(
		//&widget.Card{Title: "User Interface", Content: interfaceContainer},
		&widget.Card{Title: "Данные", Content: dataContainer},
		&widget.Card{Title: "Другие настройки", Content: otherContainer},
	))
}

func (s *settings) tabItem() *container.TabItem {
	return &container.TabItem{Text: "Настройки", Icon: theme.SettingsIcon(), Content: s.buildUI()}
}
