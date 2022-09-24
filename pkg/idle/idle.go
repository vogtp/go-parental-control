//go:build !windows

package idle

import (
	"time"

	"github.com/vogtp/go-hcl"
)

func Duration() time.Duration {
	hcl.Warn("Idle not implemented")
	return 0
}

func Time() time.Time {
	d := Duration()
	return time.Now().Add(-d)
}
