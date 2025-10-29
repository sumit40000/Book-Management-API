package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title string `json:"title"`
	Author string `json:"author"`
	Price float64 `json:"price"`
	Rating float64 `json:"rating"`
	Read bool `json:"read"`
}

