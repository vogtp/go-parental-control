package systray

import (
	"os"

	"fyne.io/fyne/v2"
	"github.com/emersion/go-autostart"
)

func autostartMenu(menu *fyne.Menu) *fyne.MenuItem {
	app := &autostart.App{
		Name:        "go-win-idle",
		DisplayName: "idle checker",
		Exec:        []string{os.Args[0]},
	}
	m := fyne.NewMenuItem(autostartLabel(app), func() {})
	m.Checked = app.IsEnabled()

	m.Action = func() {
		//if app.IsEnabled() {
		//	app.Disable()
		//} else {
		app.Enable()
		//}
		m.Label = autostartLabel(app)
		m.Checked = app.IsEnabled()
		menu.Refresh()
	}
	return m
}

func autostartLabel(app *autostart.App) string {
	//	if app.IsEnabled() {
	//		return "deinstallieren"
	//	} else {
	return "installieren"
	//	}
}
