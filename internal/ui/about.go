package ui

import (
	"net/url"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/kirillz/beget-mgr/v2/internal/util"
)

type about struct {
	icon        *clickableIcon
	nameLabel   *widget.Label
	helloLabel  *widget.Label
	spacerLabel *widget.Label
	hyperlink   *widget.Hyperlink

	app fyne.App
}

func newAbout(app fyne.App) *about {
	return &about{app: app}
}

func (a *about) buildUI() *fyne.Container {
	const (
		https   = "https"
		github  = "github.com"
		version = "v0.0.1"
	)

	repoURL := &url.URL{Scheme: https, Host: github, Path: "/kirillz/beget-mgr"}
	a.icon = newClickableIcon(a.app.Icon(), repoURL, a.app)
	dt := helloCounter()
	a.helloLabel = util.NewBoldLabel(dt)
	a.nameLabel = util.NewBoldLabel("Beget-mgr")
	a.spacerLabel = util.NewBoldLabel("-")

	releaseURL := &url.URL{
		Scheme: https, Host: github,
		Path: "/kirillz/beget-mgr/releases/tag/" + version,
	}
	a.hyperlink = &widget.Hyperlink{Text: version, URL: releaseURL, TextStyle: fyne.TextStyle{Bold: true}}

	spacer := &layout.Spacer{}
	return container.NewVBox(
		spacer,
		container.NewHBox(spacer, a.helloLabel, spacer),
		container.NewHBox(spacer, a.icon, spacer),

		container.NewHBox(
			spacer,
			a.nameLabel,
			a.spacerLabel,
			a.hyperlink,
			spacer,
		),
		spacer,
	)
}

func (a *about) tabItem() *container.TabItem {
	return &container.TabItem{Text: "О программе", Icon: theme.InfoIcon(), Content: a.buildUI()}
}

type clickableIcon struct {
	widget.BaseWidget
	app  fyne.App
	url  *url.URL
	icon *canvas.Image
}

func (c *clickableIcon) Tapped(_ *fyne.PointEvent) {
	err := c.app.OpenURL(c.url)
	if err != nil {
		fyne.LogError("Не могу открыть репозиторий: ", err)
	}
}

func helloCounter() string {
	t := time.Now()
	var res string
	switch {
	case t.Hour() < 12:
		res = "Доброе утро!"
	case t.Hour() < 17:
		res = "Добрый день."
	default:
		res = "Добрый вечер."
	}
	return res
}

func (c *clickableIcon) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func (c *clickableIcon) CreateRenderer() fyne.WidgetRenderer {
	c.ExtendBaseWidget(c)
	return widget.NewSimpleRenderer(c.icon)
}

func (c *clickableIcon) MinSize() fyne.Size {
	return fyne.Size{Width: 100, Height: 100}
}

func newClickableIcon(res fyne.Resource, url *url.URL, app fyne.App) *clickableIcon {
	return &clickableIcon{app: app, url: url, icon: canvas.NewImageFromResource(res)}
}
