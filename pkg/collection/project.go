package collection

import (
	"gopkg.in/mgo.v2/bson"
)

// Project collection
type Project struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	APIKeys     []string      `json:"api_keys" bson:"api_keys"`
}
