package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"id,omitempty"`
	MovieName string             `json:"movie_name,omitempty"`
	Watched   bool               `json:"watched,omitempty"`
	UniqueId  string             `json:"unique_id,omitempty"`
}
