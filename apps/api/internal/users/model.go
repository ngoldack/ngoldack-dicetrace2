package users

type User struct {
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
}
