package repository

import (
	"github.com/neko-neko/goflippy/pkg/collection"
	"github.com/neko-neko/goflippy/pkg/store"
	"gopkg.in/mgo.v2/bson"
)

// FeatureRepositoryMongoDB is project collection interface
type FeatureRepositoryMongoDB struct {
	db *store.DB
}

// NewFeatureRepositoryMongoDB returns new object
func NewFeatureRepositoryMongoDB(db *store.DB) *FeatureRepositoryMongoDB {
	return &FeatureRepositoryMongoDB{
		db: db,
	}
}

// Find returns features
func (f FeatureRepositoryMongoDB) Find(projectID string) ([]collection.Feature, error) {
	var features []collection.Feature

	err := f.db.MongoDB.C("features").Find(bson.M{"project_id": bson.ObjectIdHex(projectID)}).All(&features)
	return features, err
}

// FindByKey returns a feature document
func (f *FeatureRepositoryMongoDB) FindByKey(key string, projectID string) (collection.Feature, error) {
	var feature collection.Feature

	err := f.db.MongoDB.C("features").Find(bson.M{"key": key, "project_id": bson.ObjectIdHex(projectID)}).One(&feature)
	return feature, err
}

// Add a feature document
func (f *FeatureRepositoryMongoDB) Add(feature *collection.Feature) error {
	return f.db.MongoDB.C("features").Insert(feature)
}

// Update feature a document
func (f *FeatureRepositoryMongoDB) Update(feature *collection.Feature) error {
	return f.db.MongoDB.C("features").UpdateId(feature.ID, feature)
}

// Delete a feature document
func (f *FeatureRepositoryMongoDB) Delete(id string) error {
	return f.db.MongoDB.C("features").RemoveId(bson.ObjectIdHex(id))
}
