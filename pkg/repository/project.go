package repository

import "github.com/neko-neko/goflippy/pkg/collection"

// ProjectRepository is data access interface
type ProjectRepository interface {
	// FindAll returns all projects
	FindAll() ([]collection.Project, error)

	// FindProjectIDByAPIKey returns ProjectID from APIKey
	FindProjectIDByAPIKey(key string) (string, error)

	// FindProjectByID returns Project from ID
	FindProjectByID(id string) (collection.Project, error)

	// FindByID returns Project from ID
	FindByID(id string) (collection.Project, error)

	// Add a project
	Add(project *collection.Project) error

	// Update a project
	Update(project *collection.Project) error

	// Delete a project
	Delete(id string) error
}
