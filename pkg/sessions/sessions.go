package sessions

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/vogtp/go-hcl"
)

var (
	timeFormatString = "2006.01.02 15:04:05"

	dateFormatString = "2006.01.02"
)

type sessionData struct {
	current      string
	LastActivity map[string]*lastActData
}

type activeSession struct {
	user        string
	connectTime time.Time
	lastInput   time.Time
	lockstate   LockState
	id          string
	hostname    string
}

func (s *Service) collectSessions(ctx context.Context) {
	sessions, err := s.getSessions()
	if err != nil {
		fmt.Printf("Could not collect sessions: %v", err)
		return
	}
	userClient, err := s.db.UserClient()
	if err != nil {
		fmt.Printf("Cannot get client: %v", err)
		return
	}
	s.session.current = ""

	for _, session := range sessions {
		usr := session.user

		lastActivity, hasLastActivity := s.session.LastActivity[usr]
		if hasLastActivity {
			if lastActivity.LastActivity.Before(session.connectTime) && time.Since(lastActivity.lastUpdate) < 5*time.Minute {
				hasLastActivity = false
			} else {
				s.lastUser = usr
			}
		}

		delta := time.Since(s.lastRun)

		if hasLastActivity && lastActivity.LastActivity.Before(s.lastRun) && lastActivity.lastUpdate.After(s.lastRun) {
			hcl.Infof("last activity: %v -> delta", lastActivity.LastActivity)
			delta = 0
		}
		// else if !lastInput.IsZero() {
		// 	hcl.Infof("last input: %v", lastActivity)
		// 	delta = s.lastRun.Sub(lastInput)
		// }
		hcl.Infof("%s delta %v (last %v) has last activity %v", usr, delta, s.lastUser, hasLastActivity)
		s.session.current = fmt.Sprintf("%s%-10s %-9s %-8s %s\n", s.session.current, usr, session.lockstate, session.hostname, session.connectTime.Format(time.RFC3339))
		//s.OLDsession.current = fmt.Sprintf("%sSession %s: %s (%s) %s\n  connect: %v\n  input: %v\n", s.OLDsession.current, session.id, session.hostname, session.lockstate, usr, session.connectTime.Format(time.RFC3339), session.lastInput.Format(time.RFC3339))
		if session.lockstate == Unlocked && s.lastUser == usr && delta <= time.Since(s.lastRun) {

			if err := userClient.AddTime(ctx, usr, delta); err != nil {
				hcl.Warnf("Cannot save user data of %s: %v", usr, err)
			}
		}
		if session.lockstate == Unlocked {
			s.lastUser = usr
		}
	}
	s.lastRun = time.Now()
	s.printSessions(ctx, os.Stdout)

	s.save()

}

func (s *Service) save() {
	if err := s.wirteJSONFile(); err != nil {
		fmt.Printf("Cannot write session data to %v: %v\n", s.jsonDBFile, err)
	}
}

func (s *Service) wirteJSONFile() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, err := json.Marshal(s.session)
	if err != nil {
		return fmt.Errorf("cannot marshal json: %w", err)
	}
	err = ioutil.WriteFile(s.jsonDBFile, b, 0644)
	if err != nil {
		return fmt.Errorf("cannot write file %s: %w", s.jsonDBFile, err)
	}
	return nil
}

func (s *Service) load() {
	if err := s.readJSONFile(); err != nil {
		fmt.Printf("Cannot read session data: %v\n", err)
	}
}

func (s *Service) readJSONFile() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, err := os.Stat(s.jsonDBFile)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("not found %s: %w", s.jsonDBFile, err)
	}
	plan, _ := ioutil.ReadFile(s.jsonDBFile)
	err = json.Unmarshal(plan, s.session)
	if err != nil {
		fmt.Printf("JSON:\n%s\n", plan)
		return fmt.Errorf("error loading json from %v: %w", s.jsonDBFile, err)
	}
	return nil
}
