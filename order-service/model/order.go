package model

type Order struct {
	ID        string `bson:"_id,omitempty" json:"id"`
	UserID    string `bson:"user_id,omitempty" json:"user_id,omitempty"`
	ProductID string `bson:"product_id,omitempty" json:"product_id,omitempty"`
	Quantity  uint   `bson:"quantity,omitempty" json:"quantity,omitempty"`
}
