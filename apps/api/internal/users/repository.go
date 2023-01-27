package users

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog/log"
)

type UserRepository interface {
	RegisterUser(newUser *User) (user *User, err error)
	GetUser(uuid uuid.UUID) (*User, error)
	DeleteUser(uuid uuid.UUID) (err error)
	FindUserWithUsername(username string) ([]User, error)
}

type UserNeo4JRepository struct {
	Driver *neo4j.Driver
}

var validate = validator.New()

func unwrapRecordToUser(record *neo4j.Record) (user *User, err error) {
	user = &User{}
	if value, ok := record.Get("uuid"); ok {
		user.UUID, err = uuid.Parse(value.(string))
		if err != nil {
			return nil, err
		}
	}
	if value, ok := record.Get("username"); ok {
		user.Username = value.(string)
	}
	if value, ok := record.Get("email"); ok {
		user.Email = value.(string)
	}
	if value, ok := record.Get("name"); ok {
		user.Name = value.(string)
	}

	err = validate.Struct(user)
	if err != nil {
		return nil, err
	}

	return
}

func (u UserNeo4JRepository) FindUserWithUsername(username string) (users []User, err error) {
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
		err = tsx.Close()
	}()

	// run the cypher
	// language=cypher
	cypher := "MATCH (u:USER) WHERE toLower(u.username) CONTAINS toLower($username) RETURN u.uuid AS uuid, u.username AS username, u.email AS email, u.name AS name"
	props := map[string]interface{}{
		"username": username,
	}
	log.Debug().Str("system", "user/repository").Str("cypher", cypher).Interface("props", props).Msg("executing cypher")
	response, err := tsx.Run(cypher, props)
	if err != nil {
		return nil, err
	}

	users = make([]User, 0)
	for response.Next() {
		log.Debug().Interface("record", response.Record()).Msg("record")
	}

	return nil, nil
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
		err = tsx.Close()
	}()

	// run the cypher
	// language=cypher
	cypher := "CREATE (u:USER {uuid: apoc.create.uuid(), username: $username, email: $email, name: $name}) RETURN u.uuid AS uuid, u.username AS username, u.email AS email, u.name AS name"
	props := map[string]interface{}{
		"username": newUser.Username,
		"email":    newUser.Email,
		"name":     newUser.Name,
	}
	log.Debug().Str("system", "user/repository").Str("cypher", cypher).Interface("props", props).Msg("executing cypher")
	response, err := tsx.Run(cypher, props)
	if err != nil {
		return nil, err
	}

	// Unwrap the record
	if response.Next() {
		user, err = unwrapRecordToUser(response.Record())
		if err != nil {
			return nil, err
		}
	}

	// Commit the changes
	if user != nil {
		err = tsx.Commit()
		if err != nil {
			log.Err(err).Msg("database commit failed")
			return nil, err
		}
	}

	return user, err
}

func (u UserNeo4JRepository) GetUser(uuid uuid.UUID) (user *User, err error) {
	// Create Session
	session := (*u.Driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer func() {
		err = session.Close()
	}()

	// Create a transaction
	tsx, err := session.BeginTransaction()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = tsx.Close()
	}()

	// language=cypher
	cypher := "MATCH (u:USER {uuid: $uuid}) RETURN u.uuid AS uuid, u.username AS username, u.email AS email, u.name AS name"
	props := map[string]interface{}{"uuid": uuid.String()}
	log.Debug().Str("system", "user/repository").Str("cypher", cypher).Interface("props", props).Msg("executing cypher")
	response, err := tsx.Run(cypher, props)
	if err != nil {
		return nil, err
	}

	// Unwrap the record
	if response.Next() {
		user, err = unwrapRecordToUser(response.Record())
		if err != nil {
			return nil, err
		}
	}

	// Commit the changes
	if user != nil {
		err = tsx.Commit()
		if err != nil {
			log.Err(err).Msg("database commit failed")
			return nil, err
		}
	}

	return
}

func (u UserNeo4JRepository) DeleteUser(uuid uuid.UUID) (err error) {
	// Create Session
	session := (*u.Driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func() {
		err = session.Close()
	}()

	// Create a transaction
	tsx, err := session.BeginTransaction()
	if err != nil {
		return err
	}
	defer func() {
		err = tsx.Close()
	}()

	// language=cypher
	cypher := "MATCH (u:USER {uuid: $uuid}) DELETE u"
	props := map[string]interface{}{"uuid": uuid.String()}
	log.Debug().Str("system", "user/repository").Str("cypher", cypher).Interface("props", props).Msg("executing cypher")
	_, err = tsx.Run(cypher, props)
	if err != nil {
		return err
	}

	// commit the changes
	err = tsx.Commit()
	if err != nil {
		log.Err(err).Msg("database commit failed")
		return err
	}

	return
}
