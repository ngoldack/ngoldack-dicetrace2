package users

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog/log"
)

type UserRepository interface {
	RegisterUser(newUser *User) (user *User, err error)
	GetUser(username string) (*User, error)
	DeleteUser(username string) (err error)
}

type UserNeo4JRepository struct {
	Driver *neo4j.Driver
}

func (u UserNeo4JRepository) RegisterUser(newUser *User) (user *User, err error) {
	// Create Session
	session := (*u.Driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func() {
		err = session.Close()
	}()

	// Create a transaction
	tsx, err := session.BeginTransaction()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = tsx.Rollback()
	}()

	// run the cypher
	// language=cypher
	cypher := "CREATE (u:USER {username: $username, email: $email, name: $name}) RETURN u {.username, .email, .name}"
	props := map[string]interface{}{
		"username": newUser.Username,
		"email":    newUser.Email,
		"name":     newUser.Name,
	}
	response, err := tsx.Run(cypher, props)
	if err != nil {
		return nil, err
	}

	// Unwrap the record
	if response.Next() {
		user = &User{}
		record := response.Record()

		if value, ok := record.Get("username"); ok {
			user.Username = value.(string)
		}
		if value, ok := record.Get("email"); ok {
			user.Email = value.(string)
		}
		if value, ok := record.Get("name"); ok {
			user.Name = value.(string)
		}
	}

	// Commit the changes
	if user != nil {
		err = tsx.Commit()
		if err != nil {
			log.Err(err).Msg("commit failed")
			return nil, err
		}
	}

	return nil, err
}

func (u UserNeo4JRepository) GetUser(username string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserNeo4JRepository) DeleteUser(username string) (err error) {
	//TODO implement me
	panic("implement me")
}
