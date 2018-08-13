package repository

import "github.com/neko-neko/goflippy/pkg/collection"

// UserRepository is data access interface
type UserRepository interface {
	// Add a user document
	Add(user *collection.User) error

	// Update a user document
	Update(user *collection.User) error

	// Delete a user document
	Delete(id string) error

	// Find return user documents
	Find(projectID string) ([]collection.User, error)

	// FindByID returns a user document
	FindByID(id string, projectID string) (collection.User, error)

	// FindByUUID returns a user document
	FindByUUID(uuid string, projectID string) (collection.User, error)
}
