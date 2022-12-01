package db

import (
	"context"
	"fmt"

	"github.com/vogtp/go-parental-control/ent"
	_ "modernc.org/sqlite"
)

type Access struct {
	Client *ent.Client
}

func (a *Access) UserClient() (*UserClient, error) {
	err := a.createClient()
	if err != nil {
		return nil, fmt.Errorf("cannot create user client: %w", err)
	}
	return &UserClient{access: a}, nil
}

func (a *Access) Close() {
	a.Client.Close()
}

func (a *Access) createClient() error {
	if a.Client == nil {
		client, err := ent.Open("sqlite3", "file:sessions.sqlite?cache=shared&_fk=1")
		if err != nil {
			return fmt.Errorf("failed opening connection to sqlite: %w", err)
		}

		// Run the auto migration tool.
		if err := client.Schema.Create(context.Background()); err != nil {
			return fmt.Errorf("failed creating schema resources: %v", err)
		}
		a.Client = client
	}
	return nil
}
