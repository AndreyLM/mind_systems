package db

import (
	"context"

	"github.com/andreylm/mind_systems/pkg/ent"
	// sqllite dialect
	_ "github.com/mattn/go-sqlite3"
)

// CreateConnection - creates connection
func CreateConnection() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, err
	}
	// run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		client.Close()
		return nil, err
	}

	return client, nil
}
