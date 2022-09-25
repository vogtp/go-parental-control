package systray

import (
	"context"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/vogtp/go-hcl"
)

var (
	closeWindow fyne.Window
	code        binding.String
)

func checkClose(ctx context.Context, a fyne.App) {
	if perventQuit {
		hcl.Warn("Preventing quit")
		a.SendNotification(fyne.NewNotification("Kann nicht aufhören", "Game Zeit überschritten"))
	}
	if closeWindow == nil {
		closeWindow = a.NewWindow("Quit")
		code = binding.NewString()
		closeWindow.SetContent(container.NewVBox(
			widget.NewLabel("Enter Code"),
			widget.NewEntryWithData(code),
		))
		closeWindow.SetOnClosed(func() {

		})
	}
	closeWindow.Show()
	for range time.NewTicker(time.Second).C {
		if s, err := code.Get(); err == nil && s == "4711" {
			hcl.Warn("CLOSING")
			//_, c := context.WithCancel(ctx)
			close(IdleLoop)
			CancelApp()
			return
		}
	}

	<-ctx.Done()
}
