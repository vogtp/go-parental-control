package sessions

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/vogtp/go-hcl"
)

const ActivtiyPage = "/activity"

type lastActData struct {
	User         string
	lastUpdate   time.Time
	LastActivity time.Time
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
	hcl.Infof("Got last activity: %s %s", la.User, la.LastActivity.Format(time.RFC3339))
}

func SendLastActivity(user string, last time.Time) {
	var body bytes.Buffer
	la := &lastActData{
		User:         user,
		LastActivity: last,
	}
	err := json.NewEncoder(&body).Encode(la)
	if err != nil {
		hcl.Errorf("cannot encode last activity: %v", err)
		return
	}
	r, err := http.Post("http://localhost:4711"+ActivtiyPage, "application/json", &body)
	if err != nil {
		hcl.Errorf("cannot send last activity: %v", err)
		return
	}
	if r.StatusCode != http.StatusOK {
		hcl.Errorf("cannot send last activity: %v", r.Status)
		return
	}
	if hcl.IsGoRun() {
		hcl.Infof("sent last activity: %v", r.Status)
	}
}
