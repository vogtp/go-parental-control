package sessions

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/vogtp/go-hcl"
	"github.com/vogtp/go-parental-control/ent"
	"github.com/vogtp/go-parental-control/pkg/user"
)

const ActivtiyPage = "/activity"

type lastActData struct {
	User         string
	lastUpdate   time.Time
	LastActivity time.Time
}

type UserActivityReport struct {
	*ent.Activity
	ActivityExceeded bool
}

func (uar UserActivityReport) Duration() time.Duration {
	return time.Duration(uar.Activity.Duration).Truncate(time.Minute)
}

func checkActivityExcced(act *ent.Activity) bool {
	allowed := user.GetAllowedDuration(act.Username)
	return time.Duration(act.Duration) > allowed.Today()
}

func (s *Service) handleLastActivity(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	la := new(lastActData)
	err = json.Unmarshal(data, la)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	la.lastUpdate = time.Now()
	s.session.LastActivity[la.User] = la
	w.WriteHeader(http.StatusOK)
	uc, err := s.db.UserClient()
	if err != nil {
		hcl.Warnf("Cannot userclient: %v", err)
	}
	act, err := uc.GetCurrentUserActivity(r.Context(), la.User)
	if err != nil {
		hcl.Warnf("Cannot get activity: %v", err)
	}
	uar := &UserActivityReport{
		Activity:         act,
		ActivityExceeded: checkActivityExcced(act),
	}
	err = json.NewEncoder(w).Encode(uar)
	if err != nil {
		hcl.Warnf("Cannot encode activity: %v", err)
	}
	hcl.Infof("Got last activity: %s %s", la.User, la.LastActivity.Format(time.RFC3339))
}

func SendLastActivity(ctx context.Context, user string, last time.Time) (*UserActivityReport, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	var body bytes.Buffer
	la := &lastActData{
		User:         user,
		LastActivity: last,
	}
	err := json.NewEncoder(&body).Encode(la)
	if err != nil {
		return nil, fmt.Errorf("cannot encode last activity: %w", err)

	}

	// "application/json",
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:4711"+ActivtiyPage, &body)
	if err != nil {
		return nil, fmt.Errorf("cannot create post request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot send last activity: %w", err)
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot send last activity: %v", r.Status)
	}
	bdy, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot read body: %w", err)
	}
	if hcl.IsGoRun() {
		hcl.Infof("sent last activity: %v", r.Status)
	}
	act := &UserActivityReport{}
	err = json.Unmarshal(bdy, act)
	return act, err
}
