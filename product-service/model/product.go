package model

type Product struct {
	ID    string `bson:"_id,omitempty" json:"id"`
	Name  string `bson:"name,omitempty" json:"name,omitempty"`
	Price uint   `bson:"price,omitempty" json:"price,omitempty"`
}
