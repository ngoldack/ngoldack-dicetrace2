package database

import (
	"context"
	"fmt"
	"github.com/ngoldack/dicetrace/internal/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
}

// Force interface implementation
var _ app.Controllable = &Database{}

func NewDatabase(uri string) (*Database, error) {
	db := &Database{}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("error while creating mongo client: %w", err)
	}

	db.client = client
	return db, nil
}

func (db *Database) Start(ctx context.Context) error {
	err := db.client.Connect(ctx)
	if err != nil {
		return fmt.Errorf("error while connecting to database: %w", err)
	}
	return nil
}

func (db *Database) Stop(ctx context.Context) error {
	err := db.client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("error while closing datase connection: %w", err)
	}
	return nil
}
