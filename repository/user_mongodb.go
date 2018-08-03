package repository

import (
	"github.com/neko-neko/goflippy/collection"
	"github.com/neko-neko/goflippy/store"
	"gopkg.in/mgo.v2/bson"
)

// UserRepositoryMongoDB is project collection interface
type UserRepositoryMongoDB struct {
	db *store.DB
}

// NewUserRepositoryMongoDB returns new object
func NewUserRepositoryMongoDB(db *store.DB) *UserRepositoryMongoDB {
	return &UserRepositoryMongoDB{
		db: db,
	}
}

// Add a user document
func (u *UserRepositoryMongoDB) Add(user *collection.User) error {
	return u.db.MongoDB.C("users").Insert(user)
}

// Update a user document
func (u *UserRepositoryMongoDB) Update(user *collection.User) error {
	return u.db.MongoDB.C("users").UpdateId(user.ID, user)
}

// Delete a user document
func (u *UserRepositoryMongoDB) Delete(id string) error {
	return u.db.MongoDB.C("users").RemoveId(bson.ObjectIdHex(id))
}

// Find user documents
func (u *UserRepositoryMongoDB) Find(projectID string) ([]collection.User, error) {
	var users []collection.User

	err := u.db.MongoDB.C("users").Find(bson.M{"project_id": bson.ObjectIdHex(projectID)}).All(&users)
	return users, err
}

// FindByID returns a user document
func (u *UserRepositoryMongoDB) FindByID(id string, projectID string) (collection.User, error) {
	var user collection.User

	err := u.db.MongoDB.C("users").Find(bson.M{"_id": id, "project_id": bson.ObjectIdHex(projectID)}).One(&user)

	return user, err
}

// FindByUUID returns a user document
func (u *UserRepositoryMongoDB) FindByUUID(uuid string, projectID string) (collection.User, error) {
	var user collection.User

	err := u.db.MongoDB.C("users").Find(bson.M{"uuid": uuid, "project_id": bson.ObjectIdHex(projectID)}).One(&user)

	return user, err
}
