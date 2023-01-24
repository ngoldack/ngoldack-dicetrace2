package database

import (
	"context"
	"fmt"
	"github.com/ngoldack/dicetrace/internal/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DB struct {
	client   *mongo.Client
	database *mongo.Database
}

// Force interface implementation
var _ app.Controllable = &DB{}

func NewDBClient(uri string, database string) (*DB, error) {
	db := &DB{}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("error while creating mongo client: %w", err)
	}

	db.client = client
	db.database = client.Database(database)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) Start(ctx context.Context) error {
	err := db.client.Connect(ctx)
	if err != nil {
		return fmt.Errorf("error while connecting to database: %w", err)
	}
	return nil
}

func (db *DB) Stop(ctx context.Context) error {
	err := db.client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("error while closing datase connection: %w", err)
	}
	return nil
}

func (db *DB) GetCollection(name string) *mongo.Collection {
	return db.database.Collection(name)
}
