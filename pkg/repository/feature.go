package repository

import "github.com/neko-neko/goflippy/pkg/collection"

// FeatureRepository is data access interface
type FeatureRepository interface {
	// Find returns features
	Find(projectID string) ([]collection.Feature, error)

	// FindByKey returns a feature
	FindByKey(key string, projectID string) (collection.Feature, error)

	// Add a feature
	Add(user *collection.Feature) error

	// Update a feature
	Update(feature *collection.Feature) error

	// Delete a feature
	Delete(id string) error
}
