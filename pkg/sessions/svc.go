package sessions

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/judwhite/go-svc"
	"github.com/spf13/viper"

	"github.com/vogtp/go-hcl"
	"github.com/vogtp/go-win-session/pkg/cfg"
	"github.com/vogtp/go-win-session/pkg/db"
)

type Service struct {
	mu         sync.Mutex
	db         db.Access
	quit       chan any
	session    *sessionData
	lastRun    time.Time
	lastUser   string
	jsonDBFile string
	logWriter  io.WriteCloser
}

func (s *Service) Run() error {
	if err := s.Init(nil); err != nil {
		return fmt.Errorf("cannot init sessions: %w", err)
	}
	s.collectSessions(context.Background())
	return nil
}

func (s *Service) Init(env svc.Environment) error {
	s.quit = make(chan any)
	s.session = &sessionData{}
	go func() {
		path := "C:/ProgramData/go-win-sessions/"
		_, err := os.Stat(path)
		if errors.Is(err, os.ErrNotExist) {
			os.Mkdir(path, os.ModeAppend)
		}
		s.jsonDBFile = path + "/sessions.json"
		s.session = &sessionData{
			LastActivity: make(map[string]*lastActData),
		}
		s.lastRun = time.Now()
		fi, err := os.Create(path + "/go-parental-controll.log")
		if err != nil {
			hcl.Errorf("Cannot open log: %v", err)
		} else {
			var w io.Writer
			if hcl.IsGoRun() {
				w = io.MultiWriter(fi, os.Stdout)
			} else {
				w = fi
			}
			hcl := hcl.New(hcl.WithWriter(w))
			if hcl.IsGoRun() {
				hcl.SetLevel(hclog.Debug)
			}
			s.logWriter = fi
			hcl.Info("test")
		}
		s.load()
		if viper.GetBool(cfg.Web) {
			hcl.Infof("Starting web server: %v", viper.GetBool(cfg.Web))
			go s.serveHTTP()
		}
	}()
	return nil
}

func (s *Service) Start() error {
	go func() {
		ctx := context.Background()
		s.collectSessions(ctx)
		repeat := viper.GetDuration(cfg.Repeat)
		for {
			select {
			case <-s.quit:
				return
			case <-time.After(repeat):
				s.collectSessions(ctx)
			}
		}
	}()
	return nil
}

func (s *Service) Stop() error {
	defer func() {
		if s.logWriter != nil {
			s.logWriter.Close()
		}
	}()
	close(s.quit)
	return nil
}
