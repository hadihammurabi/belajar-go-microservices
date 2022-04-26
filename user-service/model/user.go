package model

type User struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Name     string `bson:"name,omitempty" json:"name,omitempty"`
	Email    string `bson:"email,omitempty" json:"email,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
}
