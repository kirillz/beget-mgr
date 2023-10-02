package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/kirillz/beget-mgr/v2/internal/util"
)

type accounts struct {
	addAccount *widget.Button
	delAccount *widget.Button
	addForm    *widget.Form
	table      *widget.Table

	client *Client
	window fyne.Window
	canvas fyne.Canvas
	app    fyne.App
}

func newAccounts(a fyne.App, w fyne.Window, c *Client) *accounts {
	return &accounts{app: a, window: w, client: c, canvas: w.Canvas()}
}

func (s *accounts) onAddAccount() {
	s.addForm.Resize(util.WindowSizeToDialog(s.canvas.Size()))
	s.addForm.Show()
	fmt.Println("account added")
}
func (s *accounts) onDelAccount() {
	fmt.Println("account delete clicked")
}

// Строим UI
func (s *accounts) buildUI() *fyne.Container {
	s.addAccount = &widget.Button{Text: "Добавить новый Аккаунт", Icon: theme.AccountIcon(), OnTapped: s.onAddAccount}
	s.delAccount = &widget.Button{Text: "Удалить Аккаунт", Icon: theme.AccountIcon(), OnTapped: s.onDelAccount}
	data := map[string][]string{
		"Account1": {"active", "username", "password"},
		"Account2": {"inactive", "username1", "password1"},
		"Account3": {"active", "username2", "password2"},
		"Account4": {"active", "username3", "password3"},
	}
	var kx string //temp var
	var names []string

	for k := range data {
		names = append(names, k)
		kx = k
	}

	s.table = widget.NewTable(
		func() (int, int) {
			return len(names), len(data[kx]) + 1
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Default")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			if tci.Col == 0 {
				co.(*widget.Label).SetText(names[tci.Row])
			} else {
				co.(*widget.Label).SetText(fmt.Sprint(data[names[tci.Row]][tci.Col-1]))
			}
		},
	)

	// Form

	// s.fileChoice = &widget.Button{Text: "Файл", Icon: theme.FileIcon(), OnTapped: s.onFileSend}
	// s.directoryChoice = &widget.Button{Text: "Каталог", Icon: theme.FolderOpenIcon(), OnTapped: s.onDirSend}

	// choiceContent := container.NewGridWithColumns(1, s.fileChoice, s.directoryChoice, s.textChoice)
	//s.contentPicker = dialog.NewCustom("Выберите тип контента", "Отменить", choiceContent, s.window)

	//s.contentToSend = &widget.Button{Text: "Добавить контент для отправки", Icon: theme.ContentAddIcon(), OnTapped: s.contentPicker.Show}

	box := container.NewVBox(s.addAccount, s.delAccount, &widget.Separator{}, s.table)
	return container.NewBorder(box, nil, nil, nil, s.table)
}

func (s *accounts) tabItem() *container.TabItem {
	return &container.TabItem{Text: "Аккаунты", Icon: theme.MailSendIcon(), Content: s.buildUI()}
}
