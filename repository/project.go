package repository

import "github.com/neko-neko/goflippy/collection"

// ProjectRepository is data access interface
type ProjectRepository interface {
	// FindProjectIDByAPIKey returns ProjectID from APIKey
	FindProjectIDByAPIKey(key string) (string, error)

	// FindProjectByID returns Project from ID
	FindProjectByID(id string) (collection.Project, error)
}
