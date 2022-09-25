package systray

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"github.com/spf13/viper"
	"github.com/vogtp/go-hcl"
	"github.com/vogtp/go-parental-control/pkg/cfg"
	"github.com/vogtp/go-parental-control/pkg/idle"
	"github.com/vogtp/go-parental-control/pkg/sessions"
	"github.com/vogtp/go-parental-control/pkg/user"
)

var stopIdle = make(chan any)

func getIdle(app fyne.App) {
	show := viper.GetBool(cfg.Show)
	repeat := viper.GetDuration(cfg.Repeat)
	if show {
		fmt.Printf("Loggin idle times very %v\n", repeat)
	}
	for {
		usr := user.Current()
		la := idle.Time()
		dur := idle.Duration()
		label.Text = fmt.Sprintf("%s -> %s", time.Now().Format(cfg.TimeFormat), dur.Round(time.Second))
		label.Refresh()
		if show {
			fmt.Printf("%s\t%s\t%v\t%v\n", time.Now().Format(cfg.TimeFormat), usr, dur, la.Format(cfg.TimeFormat))
		}
		if repeat < time.Second {
			return
		}
		act, err := sessions.SendLastActivity(usr, la)
		if err != nil {
			hcl.Warnf("Sending activity: %v", err)
		} else {
			d := time.Duration(act.Activity)
			if d > time.Second {
				msg := fmt.Sprintf("%s Du hast %v gespielt, hör mal auf", act.Username, d.Truncate(time.Minute))
				app.SendNotification(fyne.NewNotification("Zeit aufzuhören", msg))
			}
		}
		select {
		case <-time.After(repeat):
			continue
		case <-stopIdle:
			fmt.Println("Quit received")
			return
		}
	}
}