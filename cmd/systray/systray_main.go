package main

import (
	"fmt"

	"github.com/vogtp/go-parental-control/pkg/cfg"
	"github.com/vogtp/go-parental-control/pkg/systray"
)

func main() {
	fmt.Printf("Satrting systray %v\n", cfg.Version)

	systray.Run()
}
