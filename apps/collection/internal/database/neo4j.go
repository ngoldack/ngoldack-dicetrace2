package database

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4JClient struct {
	Driver *neo4j.Driver
}

func NewNeo4JClient(uri, username, password string) (*Neo4JClient, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))

	if err != nil {
		return nil, err
	}

	return &Neo4JClient{
		Driver: &driver,
	}, nil
}

func (n *Neo4JClient) Start(ctx context.Context) error {
	return nil
}

func (n *Neo4JClient) Stop(_ context.Context) error {
	err := (*n.Driver).Close()
	if err != nil {
		return err
	}
	return nil
}
