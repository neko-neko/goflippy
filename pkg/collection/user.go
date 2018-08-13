package collection

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User collection
type User struct {
	ID              bson.ObjectId   `json:"_id" bson:"_id,omitempty"`
	ProjectID       bson.ObjectId   `json:"project_id" bson:"project_id"`
	FirstName       string          `json:"first_name" bson:"first_name"`
	LastName        string          `json:"last_name" bson:"last_name"`
	UUID            string          `json:"uuid" bson:"uuid"`
	Email           string          `json:"email" bson:"email"`
	LastActivatedAt time.Time       `json:"last_activated_at" bson:"last_activated_at"`
	CreatedAt       time.Time       `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at" bson:"updated_at"`
	Groups          []UserGroup     `json:"groups" bson:"groups"`
	Attributes      []UserAttribute `json:"attributes" bson:"attributes"`
}

// UserGroup is nested struct
type UserGroup struct {
	Name string `json:"name" bson:"name"`
}

// UserAttribute is nested strcut
type UserAttribute struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}

// NewUser returns new user object
func NewUser() *User {
	now := time.Now()

	return &User{
		ID:              bson.NewObjectId(),
		LastActivatedAt: now,
		CreatedAt:       now,
		UpdatedAt:       now,
		Groups:          make([]UserGroup, 0),
		Attributes:      make([]UserAttribute, 0),
	}
}

// SetProjectID is set string projectID to bson.ObjectId projectID
func (u *User) SetProjectID(projectID string) {
	u.ProjectID = bson.ObjectIdHex(projectID)
}

// AppendGroup append group to user
func (u *User) AppendGroup(name string) {
	u.Groups = append(u.Groups, UserGroup{Name: name})
}

// AppendAttribute append attribute to user
func (u *User) AppendAttribute(key string, value string) {
	u.Attributes = append(u.Attributes, UserAttribute{Key: key, Value: value})
}

// ResetGroups reset user groups
func (u *User) ResetGroups() {
	u.Groups = make([]UserGroup, 0)
}

// ResetAttributes reset user attributes
func (u *User) ResetAttributes() {
	u.Attributes = make([]UserAttribute, 0)
}
