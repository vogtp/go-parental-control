package sessions

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"github.com/vogtp/go-hcl"
	"github.com/vogtp/go-parental-control/ent"
	"github.com/vogtp/go-parental-control/ent/day"
	"github.com/vogtp/go-parental-control/pkg/cfg"
)

func (s *Service) serveHTTP() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		s.printSessions(r.Context(), w)
		fmt.Fprintf(w, "\n\nLast Update: %s\n", s.lastRun.Format(cfg.TimeFormat))
		fmt.Fprintf(w, "Version: %s\n", cfg.Version)
	})
	http.HandleFunc(ActivtiyPage, s.handleLastActivity)
	if err := http.ListenAndServe(":4711", nil); err != nil {
		fmt.Printf("Cannot start http: %v", err)
	}
}

func (s *Service) printSessions(ctx context.Context, w io.Writer) {
	fmt.Fprint(w, "Current:\n\n")
	fmt.Fprint(w, s.session.current)
	if !viper.GetBool(cfg.History) {
		return
	}
	fmt.Fprint(w, "\n\nLast Activity:\n")
	for u, t := range s.session.LastActivity {
		fmt.Fprintf(w, "%s\t%10s\n",
			u,
			time.Now().Sub(t.LastActivity).Truncate(time.Second),
		)
	}
	fmt.Fprint(w, "\n\nHistory:\n\n")
	days, err := s.db.Client.Day.Query().Order(
		ent.Desc(day.FieldYear, day.FieldMonth, day.FieldDay),
	).All(ctx)
	if err != nil {
		fmt.Fprintf(w, "DB Access: %v", err)
		hcl.Warnf("Cannot access DB: %v", err)
		return
	}
	for _, day := range days {
		fmt.Fprintf(w, "==== %v.%v.%v ====\n", day.Day, day.Month, day.Year)
		acts, err := day.QueryActivity().All(ctx)
		if err != nil {
			fmt.Fprintf(w, "DB Access: %v", err)
			hcl.Warnf("Cannot access DB: %v", err)
			continue
		}
		for _, act := range acts {
			d := time.Duration(act.Activity).Truncate(time.Second)
			fmt.Fprintf(w, "  %-20s %v\n", act.Username, d)
		}
	}

}
