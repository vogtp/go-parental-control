package sessions

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/vogtp/go-hcl"
	"github.com/vogtp/go-parental-control/build"
	"github.com/vogtp/go-parental-control/pkg/cfg"
)

const (
	UpdatePage    = "/update"
	versionHeader = "version"
)

func (s *Service) handleUpdate(w http.ResponseWriter, r *http.Request) {
	build.ListFiles()
	file, err := build.GetFile("go-parental-control-systray")
	if err != nil {
		hcl.Errorf("Cannot get binary: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set(versionHeader, cfg.Version.String())
	w.Write(file)
}

func LoadSystrayBinary(ctx context.Context) ([]byte, string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:4711"+UpdatePage, nil)
	if err != nil {
		return nil, "", fmt.Errorf("cannot create post request: %w", err)
	}
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("cannot request new binary: %w", err)
	}

	bdy, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("Cannot read body: %w", err)
	}
	if r.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("server error: %s %v", r.Status, string(bdy))
	}

	return bdy, r.Header.Get(versionHeader), nil
}
