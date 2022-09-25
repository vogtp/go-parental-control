package user

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/vogtp/go-hcl"
	"gopkg.in/yaml.v3"
)

const (
	defauleAllowed = "default"
)

var (
	allowed AllowedDuration
)

func GetAllowedDuration(user string) AllowedWeekdays {
	if allowed == nil {
		allowed = make(AllowedDuration)
		if err := allowed.load(); err != nil {
			hcl.Warnf("Cannot load allowed times: %v", err)
		}
	}
	if a, ok := allowed[user]; ok {
		return a
	}
	defer func() {
		if err := allowed.save(); err != nil {
			hcl.Errorf("Cannot save allowed times: %v", err)
		}

	}()
	d, ok := allowed[defauleAllowed]
	if !ok {
		d = defaultWeekdays()
		allowed[defauleAllowed] = d
	}
	allowed[user] = d
	return d
}

type AllowedDuration map[string]AllowedWeekdays

type AllowedWeekdays map[time.Weekday]time.Duration

type dayAllowed struct {
	Weekday string
	Allowed time.Duration
}

func (aw AllowedWeekdays) Today() time.Duration {
	return aw[time.Now().Weekday()]
}

func (ad AllowedDuration) save() error {
	var data bytes.Buffer
	for u, days := range ad {
		fmt.Fprintf(&data, "%s:\n", u)
		for i := 0; i < 7; i++ {
			w := time.Weekday(i)
			fmt.Fprintf(&data, "    %s: %v\n", days[w].String(), w)
		}
	}
	return ioutil.WriteFile("allowed.yml", data.Bytes(), 0600)
}

var daysOfWeek = map[string]time.Weekday{}

func init() {
	for d := time.Sunday; d <= time.Saturday; d++ {
		daysOfWeek[d.String()] = d
	}
}

func (ad AllowedDuration) load() error {
	data, err := ioutil.ReadFile("allowed.yml")
	if err != nil {
		return err
	}
	tmp := make(map[string]map[string]time.Duration)
	err = yaml.Unmarshal(data, tmp)
	if err != nil {
		return err
	}
	for u, days := range tmp {
		wd := make(AllowedWeekdays)
		for n, t := range days {
			d, ok := daysOfWeek[n]
			if !ok {
				return fmt.Errorf("Weekday %q not found", n)
			}
			wd[d] = t
		}
		ad[u] = wd
	}
	return nil
}

func defaultWeekdays() AllowedWeekdays {
	wd := make(AllowedWeekdays, 7)
	createDefaultDay(wd, time.Monday, time.Hour+30*time.Minute)
	createDefaultDay(wd, time.Tuesday, time.Hour+30*time.Minute)
	createDefaultDay(wd, time.Wednesday, time.Hour+30*time.Minute)
	createDefaultDay(wd, time.Thursday, time.Hour+30*time.Minute)
	createDefaultDay(wd, time.Friday, time.Hour+30*time.Minute)
	createDefaultDay(wd, time.Saturday, time.Hour+30*time.Minute)
	createDefaultDay(wd, time.Sunday, time.Hour+30*time.Minute)
	return wd
}

func createDefaultDay(days AllowedWeekdays, day time.Weekday, allowed time.Duration) {
	days[day] = allowed
}
