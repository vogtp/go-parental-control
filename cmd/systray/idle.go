package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/vogtp/go-win-session/pkg/cfg"
	"github.com/vogtp/go-win-session/pkg/idle"
	"github.com/vogtp/go-win-session/pkg/sessions"
	"github.com/vogtp/go-win-session/pkg/user"
)

func getIdle() {
	show := viper.GetBool(cfg.Show)
	repeat := viper.GetDuration(cfg.Repeat)
	if show {
		fmt.Printf("Loggin idle times very %v\n", repeat)
	}
	for {
		usr := user.Current()
		la := idle.Time()
		dur := idle.Duration()
		sessions.SendLastActivity(usr, la)
		l := fmt.Sprintf("%s\t%s\t%v\t%v\n",
			time.Now().Format(time.RFC3339),
			usr,
			dur,
			la.Format(time.RFC3339),
		)
		label.Text = fmt.Sprintf("%s -> %s", time.Now().Format("2006-01-02T15:04:05"), dur.Round(time.Second))
		label.Refresh()
		//	systray.SetTooltip(fmt.Sprintf("%s -> %s", time.Now().Format(time.RFC3339), la.Format(time.RFC3339)))
		if show {
			fmt.Println(l)
		}
		if repeat < time.Second {
			return
		}
		select {
		case <-time.After(repeat):
			continue
		case <-run:
			fmt.Println("Quit received")
			return
		}
	}
}
