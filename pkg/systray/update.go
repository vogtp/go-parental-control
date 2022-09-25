package systray

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"github.com/emersion/go-autostart"
	"github.com/vogtp/go-hcl"
	"github.com/vogtp/go-parental-control/pkg/sessions"
)

func update(ctx context.Context, app fyne.App) {
	binary, newVersion, err := sessions.LoadSystrayBinary(ctx)
	if err != nil {
		hcl.Errorf("updating binary failed: %v", err)
		return
	}
	hcl.Infof("Got new binary size: %d", len(binary))
	currentBin := os.Args[0]
	newBin := newFileName(currentBin, newVersion)

	hcl.Infof("Current binary: %s", currentBin)
	hcl.Infof("New binary: %s", newBin)
	err = ioutil.WriteFile(newBin, binary, 0777)
	if err != nil {
		hcl.Warnf("Cannot write new binary %s: %v", newBin, err)
		return
	}
	auto := &autostart.App{
		Name:        "go-win-idle",
		DisplayName: "idle checker",
		Exec:        []string{currentBin},
	}
	if err := auto.Disable(); err != nil {
		hcl.Warnf("Cannot disable old binary: %v", err)
	}
	auto.Exec = []string{newBin}

	if err := auto.Enable(); err != nil {
		hcl.Warnf("Cannot enable new binary: %v", err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		hcl.Errorf("cannot get cwd: %v", err)
	}
	cmd := exec.Command(cwd + "/" + newBin)
	cmd.Dir = cwd
	err = cmd.Start()
	if err != nil {
		hcl.Errorf("Cannot start new binary: %v", err)
	}
	cmd.Process.Release()
	os.Exit(0)
}

const exe = ".exe"

func newFileName(name string, version string) string {
	ext := ""
	if strings.HasSuffix(name, exe) {
		ext = exe
		name = name[:len(name)-len(exe)]
	}
	return fmt.Sprintf("%s-%s%s",
		"go-parental-control-systray",
		version,
		ext,
	)
}
