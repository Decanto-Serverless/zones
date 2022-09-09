package models

import (
	"zones/db"
	"zones/env"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Country struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Country string             `json:"country,omitempty"`
}

var CountryCollection *mongo.Collection = db.GetInstance().Client().Database(env.GetInstance().DB).Collection("g_countries")
