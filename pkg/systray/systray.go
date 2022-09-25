package systray

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/vogtp/go-parental-control/pkg/cfg"
)

var (
	label      *widget.Label
	mainWindow fyne.Window
)

func runSystray(ctx context.Context) fyne.App {
	a := app.NewWithID("Check activity")

	if mainWindow == nil {
		mainWindow = a.NewWindow("Check activity")
		label = widget.NewLabel("Starting activity checker...")
		mainWindow.SetContent(container.NewVBox(
			widget.NewLabel(fmt.Sprintf("Version: %v", cfg.Version)),
			label,
			widget.NewButton("Quit", func() { checkClose(ctx, a) }),
		),
		)
	}
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("Check activity",
			fyne.NewMenuItem("Show", func() {
				mainWindow.Show()
			}))
		m.Items = append(m.Items, autostartMenu(m))
		desk.SetSystemTrayMenu(m)
	}
	return a
}
