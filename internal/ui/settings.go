package ui

import (
	"fyne.io/fyne/v2"
	appearance "fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type settings struct {
	componentSlider     *widget.Slider
	componentLabel      *widget.Label
	verifyRadio         *widget.RadioGroup
	appID               *widget.Entry
	rendezvousURL       *widget.Entry
	transitRelayAddress *widget.Entry

	window fyne.Window
	app    fyne.App
}

func newSettings(a fyne.App, w fyne.Window) *settings {
	return &settings{app: a, window: w}
}

func (s *settings) onComponentsChange(value float64) {

	s.componentLabel.SetText(string('0' + byte(value)))
}

func (s *settings) buildUI() *container.Scroll {
	onOffOptions := []string{"On", "Off"}

	s.componentSlider, s.componentLabel = &widget.Slider{Min: 2.0, Max: 6.0, Step: 1, OnChanged: s.onComponentsChange}, &widget.Label{}

	s.verifyRadio = &widget.RadioGroup{Options: onOffOptions, Horizontal: true, Required: true}

	interfaceContainer := appearance.NewSettings().LoadAppearanceScreen(s.window)

	wormholeContainer := container.NewVBox(
		container.NewGridWithColumns(2,
			newBoldLabel("Verify before accepting"), s.verifyRadio,
			newBoldLabel("Passphrase length"),
			container.NewBorder(nil, nil, nil, s.componentLabel, s.componentSlider),
		),
		&widget.Accordion{Items: []*widget.AccordionItem{
			{Title: "Advanced", Detail: container.NewGridWithColumns(2,
				newBoldLabel("AppID"), s.appID,
				newBoldLabel("Rendezvous URL"), s.rendezvousURL,
				newBoldLabel("Transit Relay Address"), s.transitRelayAddress,
			)},
		}},
	)

	return container.NewScroll(container.NewVBox(
		&widget.Card{Title: "User Interface", Content: interfaceContainer},
		&widget.Card{Title: "Wormhole Options", Content: wormholeContainer},
	))
}

func (s *settings) tabItem() *container.TabItem {
	return &container.TabItem{Text: "Settings", Icon: theme.SettingsIcon(), Content: s.buildUI()}
}

func newBoldLabel(text string) *widget.Label {
	return &widget.Label{Text: text, TextStyle: fyne.TextStyle{Bold: true}}
}
