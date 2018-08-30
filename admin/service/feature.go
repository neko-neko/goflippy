package service

import (
	"fmt"

	"github.com/neko-neko/goflippy/pkg/collection"
	"github.com/neko-neko/goflippy/pkg/repository"
)

type FeatureService struct {
	featureRepo repository.FeatureRepository
	userRepo    repository.UserRepository
}

func NewFeatureService(featureRepo repository.FeatureRepository, userRepo repository.UserRepository) *FeatureService {
	return &FeatureService{
		featureRepo: featureRepo,
		userRepo:    userRepo,
	}
}

func (f *FeatureService) FetchFeatures(projectID string) ([]collection.Feature, error) {
	features, err := f.featureRepo.Find(projectID)
	if err != nil {
		err = NewStoreSystemError(err.Error())
	}

	return features, err
}

func (f *FeatureService) FetchFeature(key string, projectID string) (collection.Feature, error) {
	feature, err := f.featureRepo.FindByKey(key, projectID)
	if err != nil {
		err = NewStoreSystemError(err.Error())
	}

	return feature, err
}

// RegisterFeature register a feature
func (f *FeatureService) RegisterFeature(feature collection.Feature) error {
	// sort filters attributes before insert
	for i := 0; i < len(feature.Filters); i++ {
		feature.Filters[i] = feature.Filters[i].Sort()
	}

	return f.featureRepo.Add(&feature)
}

// UpdateFeature update a feature
func (f *FeatureService) UpdateFeature(feature collection.Feature) error {
	// sort Filters attributes before update
	for i := 0; i < len(feature.Filters); i++ {
		feature.Filters[i] = feature.Filters[i].Sort()
	}

	return f.featureRepo.Update(&feature)
}

// DeleteFeature delete a feature
func (f *FeatureService) DeleteFeature(key string, projectID string) error {
	feature, err := f.featureRepo.FindByKey(key, projectID)
	if err != nil {
		return NewResourceNotFoundError(fmt.Sprintf("feature does not exists %s", feature.Key))
	}

	if err := f.featureRepo.Delete(feature.ID.Hex()); err != nil {
		return NewStoreSystemError(err.Error())
	}

	return nil
}
