package systray

import (
	"context"
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

var (
	perventQuit bool
	IdleLoop    chan any
)

func getIdle(ctx context.Context, app fyne.App) {
	hcl.Info("Idle loop started")
	defer hcl.Info("Idle loop stopped")
	IdleLoop = make(chan any)
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
		act, err := sessions.SendLastActivity(ctx, usr, la)
		if err != nil {
			hcl.Warnf("Sending activity: %v", err)
		} else {
			if act.ActivityExceeded {
				msg := fmt.Sprintf("%s Du hast %v gespielt, hör mal auf", act.Username, act.Duration())
				app.SendNotification(fyne.NewNotification("Zeit aufzuhören", msg))
				perventQuit = true
			}
			if act.Version.IsNewer(cfg.Version) {
				hcl.Warnf("Peer version %s is newer than ours %s -> updating", act.Version, cfg.Version)
				update(ctx, app)
			}
		}
		select {
		case <-time.After(repeat):
			continue
		case <-ctx.Done():
			fmt.Println("Conext canceled")
			return
		case <-IdleLoop:
			fmt.Println("IdleLoop closed")
			return
		}
	}
}
