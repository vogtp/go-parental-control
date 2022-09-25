package systray

import (
	"os"

	"fyne.io/fyne/v2"
	"github.com/emersion/go-autostart"
	"github.com/vogtp/go-hcl"
)

func autostartMenu(menu *fyne.Menu) *fyne.MenuItem {
	auto := &autostart.App{
		Name:        "go-win-idle",
		DisplayName: "idle checker",
		Exec:        []string{os.Args[0]},
	}
	m := fyne.NewMenuItem(autostartLabel(auto), func() {})
	m.Checked = auto.IsEnabled()

	m.Action = func() {
		if auto.IsEnabled() {
			//if err := auto.Disable(); err != nil {
			//	hcl.Warnf("Cannot disable binary: %v", err)
			//}
		} else {
			if err := auto.Enable(); err != nil {
				hcl.Warnf("Cannot enable binary: %v", err)
			}
		}
		m.Label = autostartLabel(auto)
		m.Checked = auto.IsEnabled()
		menu.Refresh()
	}
	return m
}

func autostartLabel(app *autostart.App) string {
	if app.IsEnabled() {
		//return "deinstallieren"
		return "installiert"
	}
	return "installieren"

}
