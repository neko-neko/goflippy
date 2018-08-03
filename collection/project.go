package collection

import (
	"gopkg.in/mgo.v2/bson"
)

// Project collection
type Project struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	APIKeys     []APIKey      `json:"api_keys" bson:"api_keys"`
	Features    []Feature     `json:"features" bson:"features"`
}

// APIKey is nested structure
type APIKey struct {
	Key string `json:"key" bson:"key"`
}
