package systray

import (
	"context"
	_ "embed"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/vogtp/go-hcl"
	"github.com/vogtp/go-parental-control/pkg/cfg"
)

var CancelApp context.CancelFunc

func Run() {
	ctx, CancelApp := context.WithCancel(context.Background())
	defer CancelApp()
	hcl.Info("Run")
	pflag.Duration(cfg.Repeat, time.Minute, "Collect the stats every")
	pflag.Bool(cfg.Show, hcl.IsGoRun(), "Print output to stdout")
	viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()

	a := runSystray(ctx)
	a.Lifecycle().SetOnStopped(func() { checkClose(ctx, a) })

	go getIdle(ctx, a)
	hcl.Info("starting main loop")
	a.Run()

}
