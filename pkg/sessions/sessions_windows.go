package sessions

import (
	"fmt"

	"github.com/gentlemanautomaton/winsession"
	"github.com/gentlemanautomaton/winsession/lockstate"
	"github.com/vogtp/go-win-session/pkg/user"
)

func (s *Service) getSessions() ([]activeSession, error) {
	sessions, err := winsession.Local.Sessions(
		//winsession.Include(winsession.MatchID(0)),
		winsession.CollectSessionInfo)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve session list: %v\n", err)
	}
	sess := make([]activeSession, 0, len(sessions))
	for _, session := range sessions {
		if session.Info.LockState == lockstate.Unknown {
			continue
		}
		usr := session.Info.User()
		if len(usr) < 1 {
			continue
		}
		s := activeSession{}
		s.user = user.RemoveHostPart(usr)
		s.connectTime = session.Info.ConnectTime
		s.lastInput = session.Info.LastInputTime
		s.lockstate = LockState(session.Info.LockState)
		s.id = fmt.Sprintf("%v", session.ID)
		s.hostname = session.WindowStation
		sess = append(sess, s)
	}
	return sess, nil
}
