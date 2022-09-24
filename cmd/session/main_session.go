package main

import (
	"fmt"
	"time"

	"github.com/judwhite/go-svc"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/vogtp/go-hcl"
	"github.com/vogtp/go-win-session/pkg/cfg"
	"github.com/vogtp/go-win-session/pkg/sessions"
)

func main() {
	pflag.Bool(cfg.Web, true, "start a webserver")
	pflag.Bool(cfg.History, true, "")
	pflag.Duration(cfg.Repeat, 1*time.Minute, "")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	hcl.Infof("Web: %v repeat %v hist %v", viper.Get(cfg.Web), viper.Get(cfg.Repeat), viper.Get(cfg.History))

	if viper.GetDuration(cfg.Repeat) == 0*time.Second {
		s := sessions.Service{}
		if err := s.Run(); err != nil {
			fmt.Printf("Cannot run session: %v", err)
		}
		return
	}
	if err := svc.Run(&sessions.Service{}); err != nil {
		fmt.Printf("Cannot start service: %v", err)
	}
}
