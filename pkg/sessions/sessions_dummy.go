//go:build !windows

package sessions

import (
	"math/rand"
	"time"
)

func (s *Service) getSessions() ([]activeSession, error) {
	username := "vogtp"
	id := "1000"
	if rand.Intn(10) > 5 {
		username = "rendep00"
		id = "1001"
	}
	return []activeSession{
		{
			user:        username,
			connectTime: time.Time{},
			lastInput:   time.Time{},
			lockstate:   Unlocked,
			id:          id,
			hostname:    "dev-box",
		},
	}, nil
}
