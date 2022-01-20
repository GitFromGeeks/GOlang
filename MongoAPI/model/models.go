package model

// import (
// 	// "go.mongodb.org/mongo-driver/bson/primitive"
// )

type Netflix struct {
	ID      primitive.objectID `json:"_id,omitempty" bson:"_id,omitempty"`
	movie   string             `json:"movie,omitempty"`
	watched bool               `json:"watched,omitempty"`
}
