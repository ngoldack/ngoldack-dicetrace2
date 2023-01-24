package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	Id          primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt   primitive.Timestamp `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   primitive.Timestamp `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Name        string              `json:"name,omitempty" bson:"name,omitempty"`
	BggId       *int                `json:"bgg_id,omitempty" bson:"bgg_id,omitempty"`
	MinPlayers  *int                `json:"min_players,omitempty" bson:"min_players,omitempty"`
	MaxPlayers  *int                `json:"max_players,omitempty" bson:"max_players,omitempty"`
	AvgPlaytime *int                `json:"avg_playtime,omitempty" bson:"avg_playtime,omitempty"`
}

type Match struct {
	Id           primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt    primitive.Timestamp `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    primitive.Timestamp `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Game         []Game              `json:"game,omitempty" bson:"game,omitempty"`
	StartTime    primitive.Timestamp `json:"start_time,omitempty" bson:"start_time,omitempty"`
	EndTime      primitive.Timestamp `json:"end_time" bson:"end_time,omitempty"`
	Participants []User              `json:"participants" bson:"participants"`
	Location     string              `json:"location,omitempty" bson:"location,omitempty"`
}

type Event struct {
	Id           primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt    primitive.Timestamp `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    primitive.Timestamp `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	StartTime    primitive.Timestamp `json:"start_time,omitempty" bson:"start_time,omitempty"`
	EndTime      primitive.Timestamp `json:"end_time" bson:"end_time,omitempty"`
	Participants []User              `json:"participants" bson:"participants"`
	Matches      []Match             `json:"matches,omitempty" bson:"matches,omitempty"`
}

type Group struct {
	Id        primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt primitive.Timestamp `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.Timestamp `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Owner     *User               `json:"owner,omitempty" bson:"owner,omitempty"`
	Members   []User              `json:"members,omitempty" bson:"members,omitempty"`
	Events    []Event             `json:"events,omitempty" bson:"events,omitempty"`
}

type User struct {
	Id        primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt primitive.Timestamp `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.Timestamp `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Username  string              `json:"username,omitempty" bson:"username,omitempty"`
	Email     string              `json:"email,omitempty" bson:"email,omitempty"`
	Name      string              `json:"name,omitempty" bson:"name,omitempty"`
}
