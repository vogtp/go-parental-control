package db

import (
	"context"
	"fmt"
	"time"

	"github.com/vogtp/go-hcl"
	"github.com/vogtp/go-parental-control/ent"
	"github.com/vogtp/go-parental-control/ent/activity"
	"github.com/vogtp/go-parental-control/ent/day"
)

type UserClient struct {
	access *Access
}

func (uc *UserClient) AddTime(ctx context.Context, usr string, delta time.Duration) error {
	day, err := uc.GetDay(ctx)
	if err != nil {
		return fmt.Errorf("cannot get day: %w", err)
	}
	now := time.Now()
	act, err := uc.GetActivity(ctx, day, usr)
	if err != nil {
		return fmt.Errorf("cannot get user activit for %v.%v.%v: %w", now.Year(), now.Month(), now.Day(), err)
	}
	_, err = act.Update().SetActivity(act.Activity + int64(delta)).Save(ctx)
	return err
}

func (uc *UserClient) GetDay(ctx context.Context) (*ent.Day, error) {
	y, mt, d := time.Now().Date()
	m := int(mt)
	day, err := uc.access.Client.Day.Query().Where(
		day.Year(y),
		day.Month(m),
		day.Day(d),
	).First(ctx)
	if err == nil {
		return day, nil
	}
	hcl.Warnf("Day not found, will create it: %v", err)
	return uc.access.Client.Day.Create().SetYear(y).SetMonth(m).SetDay(d).Save(ctx)
}

func (uc *UserClient) GetActivity(ctx context.Context, day *ent.Day, usr string) (*ent.Activity, error) {
	act, err := day.QueryActivity().Where(activity.UsernameEqualFold(usr)).First(ctx)
	if err == nil {
		return act, nil
	}
	act, err = uc.access.Client.Activity.Create().SetUsername(usr).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("Cannot create activit for user %s: %w", usr, err)
	}
	_, err = day.Update().AddActivity(act).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("Cannot add activit to day %s: %w", usr, err)
	}
	return act, nil
}

func (uc *UserClient) GetCurrentUserActivity(ctx context.Context, usr string) (*ent.Activity, error) {
	day, err := uc.GetDay(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot get day for user %s activity: %v", usr, err)
	}
	act, err := day.QueryActivity().Where(activity.UsernameEqualFold(usr)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot current activty of user %s: %v", usr, err)
	}
	return act, nil
}
