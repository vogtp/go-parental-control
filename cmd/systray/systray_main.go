package main

import (
	_ "embed"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/vogtp/go-hcl"
	"github.com/vogtp/go-win-session/pkg/cfg"
)

var (
	label *widget.Label
)

func main() {
	pflag.Duration(cfg.Repeat, time.Minute, "Collect the stats every")
	pflag.Bool(cfg.Show, hcl.IsGoRun(), "Print output to stdout")
	viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()
	a := app.New()
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
	go getIdle()
	a.Run()
	run <- false
}

var run = make(chan any)
