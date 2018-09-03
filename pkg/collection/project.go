package collection

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Project collection
type Project struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	APIKeys     []string      `json:"api_keys" bson:"api_keys"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}

// NewProject returns new project object
func NewProject() *Project {
	now := time.Now()
	return &Project{
		ID:        bson.NewObjectId(),
		APIKeys:   make([]string, 0),
		CreatedAt: now,
		UpdatedAt: now,
	}
}
