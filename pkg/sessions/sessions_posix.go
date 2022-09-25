//go:build !windows

package sessions

import (
	"os"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

func (s *Service) getSessions() ([]activeSession, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}
	sessions := make(map[string]*activeSession, 0)
	for _, p := range procs {
		if b, err := p.Background(); b || err != nil {
			continue
		}
		if u, err := p.Username(); err == nil {
			create := time.Time{}
			if t, err := p.CreateTime(); err != nil {
				create = time.UnixMilli(t)
			}
			s, ok := sessions[u]
			if !ok {
				s = &activeSession{
					user:        u,
					connectTime: create,
					lastInput:   time.Time{},
					lockstate:   Unlocked,
					id:          "",
					hostname:    hostname,
				}
				sessions[u] = s
			}
			if create.After(s.connectTime) {
				s.connectTime = create
			}
		}
	}
	sess := make([]activeSession, 0)
	for _, s := range sessions {
		sess = append(sess, *s)
	}
	return sess, nil
}
