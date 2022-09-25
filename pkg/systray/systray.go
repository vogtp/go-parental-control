package systray

import (
	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var (
	label *widget.Label
)

func runSystray() fyne.App {
	a := app.NewWithID("Check activity")

	w := a.NewWindow("Check activity")
	label = widget.NewLabel("...")
	w.SetContent(label)

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("Check activity",
			fyne.NewMenuItem("Show", func() {
				w.Show()

			}))
		m.Items = append(m.Items, autostartMenu(m))
		desk.SetSystemTrayMenu(m)
	}
	return a
}
