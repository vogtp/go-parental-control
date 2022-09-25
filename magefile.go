//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// var Default = All

const (
	baseName = "go-parental-control"
	buildDir = "./build/"
	runDir   = "./run/"
)

func Install() error {
	fmt.Println("Are you Admin?")
	if err := BuildAll(); err != nil {
		return err
	}
	return InstallWinService()
}

func BuildAll() error {
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

	root := buildDir

	if err := os.Mkdir(root, 0700); err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create %q: %v", root, err)
	}
	path := filepath.Join(root, name)
	fmt.Printf("Build %s \n", path)
	ldflags := "-ldflags="
	if comp == "systray" {
		ldflags += " -H=windowsgui "
	}
	return sh.RunV(gocmd, "build", "-buildvcs=false", "-o", path, ldflags+flags(), src)
}

func flags() string {
	timestamp := time.Now().Format(time.RFC3339)
	return fmt.Sprintf(`-X="github.com/vogtp/go-parental-control/pkg/cfg.BuildInfo=%s"`, timestamp)
}

func InstallWinService() error {
	name := "go-parental-control-service"
	action := "stop"
	if err := WinService("query", name); err != nil {
		fmt.Printf("Service %s not installed, will install", name)
		action = "install"
	}
	if err := WinService(action, name); err != nil {
		fmt.Printf("Cannot %s service %s: %v", action, name, err)
	}
	if err := sh.Copy(
		fmt.Sprintf("%s/%s.exe", buildDir, name),
		fmt.Sprintf("%s/%s.exe", runDir, name),
	); err != nil {
		return err
	}
	fmt.Println("Files copied starting service")
	return WinService("start", name)
}

func WinService(action, name string) error {
	action = strings.ToLower(action)
	switch action {
	case "start", "stop":
		fmt.Printf("%s service %s", action, name)
		return sh.Run("sc.exe", action, name)
	case "show", "query":
		return sh.RunV("sc.exe", "query", name)
	case "install":
		fmt.Printf("%s service %s", action, name)
		if err := sh.Run("sc.exe", "stop", name); err != nil {
			fmt.Printf("cannot create service %s: %v\n", name, err)
		}
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("cannot get cwd: %v", err)
		}
		binpath := fmt.Sprintf(`"%s/%s/%s.exe"`, cwd, runDir, name)
		displayname := fmt.Sprintf(`%s`, name)
		return sh.RunV("sc.exe", "create", name, "binpath=", binpath, "start=", "auto", "DisplayName=", displayname)
	case "delete":
		if err := sh.Run("net", "stop", name); err != nil {
			return fmt.Errorf("cannot stop service %s: %v", name, err)
		}
		return sh.RunV("sc.exe", "DELETE", name)

	}
	return fmt.Errorf("Action not found: %s", action)
}
