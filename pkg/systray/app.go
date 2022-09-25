package systray

import (
	_ "embed"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/vogtp/go-hcl"
	"github.com/vogtp/go-parental-control/pkg/cfg"
)

func Run() {
	hcl.Info("Run")
	pflag.Duration(cfg.Repeat, time.Minute, "Collect the stats every")
	pflag.Bool(cfg.Show, hcl.IsGoRun(), "Print output to stdout")
	viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()

	a := runSystray()

	go getIdle(a)
	hcl.Info("starting main loop")
	a.Run()

	close(stopIdle)
}
