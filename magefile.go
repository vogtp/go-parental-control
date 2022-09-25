//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// var Default = All

const (
	baseName = "go-parental-control"
)

func All() error {
	if err := Build("systray"); err != nil {
		return err
	}
	if err := Build("service"); err != nil {
		return err
	}
	return nil
}

func binName(comp string) string {
	return fmt.Sprintf("%s-%s", baseName, comp)
}

func Build(comp string) error {
	name := binName(comp)
	src := fmt.Sprintf("./cmd/%s", comp)
	if runtime.GOOS == "windows" {
		name += ".exe"
	}

	gocmd := mg.GoCmd()

	root := "build/"

	if err := os.Mkdir(root, 0700); err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create %q: %v", root, err)
	}
	path := filepath.Join(root, name)
	fmt.Printf("Build %s\n", path)
	return sh.RunV(gocmd, "build", "-o", path, "-ldflags="+flags(), src)
}

func flags() string {
	timestamp := time.Now().Format(time.RFC3339)
	return fmt.Sprintf(`-X "github.com/vogtp/go-parental-control/pkg/cfg.BuildInfo=%s" `, timestamp)
}
