package service

import (
	"fmt"
	"hash/adler32"
	"time"

	"github.com/neko-neko/goflippy/pkg/collection"
	"github.com/neko-neko/goflippy/pkg/repository"
)

// FeatureService is service of feature resource
type FeatureService struct {
	userRepo    repository.UserRepository
	featureRepo repository.FeatureRepository
}

// NewFeature returns a new FeatureService instance
func NewFeature(userRepo repository.UserRepository, featureRepo repository.FeatureRepository) *FeatureService {
	return &FeatureService{
		userRepo:    userRepo,
		featureRepo: featureRepo,
	}
}

// FetchFeatures returns a few FeatureService resources
func (f *FeatureService) FetchFeatures(projectID string) ([]collection.Feature, error) {
	features, err := f.featureRepo.Find(projectID)
	if err != nil {
		err = NewStoreSystemError(err.Error())
	}

	return features, err
}

// FetchFeature returns a few FeatureService resources
func (f *FeatureService) FetchFeature(key string, projectID string) (collection.Feature, error) {
	feature, err := f.featureRepo.FindByKey(key, projectID)
	if err != nil {
		err = NewResourceNotFoundError(err.Error())
	}

	return feature, err
}

// FeatureEnabled returns a feature and verify enabled for uuid of user
func (f *FeatureService) FeatureEnabled(uuid string, key string, projectID string) (collection.Feature, bool, error) {
	var feature collection.Feature

	now := time.Now()
	user, err := f.userRepo.FindByUUID(uuid, projectID)
	if err != nil {
		return feature, false, NewResourceNotFoundError(fmt.Sprintf("user not found %s", uuid))
	}
	user.LastActivatedAt = now
	f.userRepo.Update(&user)

	feature, err = f.featureRepo.FindByKey(key, projectID)
	if err != nil {
		return feature, false, NewResourceNotFoundError(fmt.Sprintf("feature not found %s", key))
	}
	feature.LastActivatedAt = now
	f.featureRepo.Update(&feature)

	return feature, f.enabled(feature, user), nil
}

// RegisterFeature register a feature
func (f *FeatureService) RegisterFeature(feature collection.Feature) error {
	// sort strategy attributes before insert
	for i := 0; i < len(feature.Strategies); i++ {
		feature.Strategies[i] = feature.Strategies[i].Sort()
	}

	return f.featureRepo.Add(&feature)
}

// UpdateFeature update a feature
func (f *FeatureService) UpdateFeature(feature collection.Feature) error {
	// sort strategy attributes before update
	for i := 0; i < len(feature.Strategies); i++ {
		feature.Strategies[i] = feature.Strategies[i].Sort()
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

// enabled verify feature?
func (f *FeatureService) enabled(feature collection.Feature, user collection.User) bool {
	for _, strategy := range feature.Strategies {
		switch strategy.Type {
		case collection.ToggleStrategyTypeSimple:
			if strategy.Enable {
				return true
			}
			break
		case collection.ToggleStrategyTypeGroup:
			if f.groupToggleEnabled(strategy, user.Groups) {
				return true
			}
			break
		case collection.ToggleStrategyTypeAttribute:
			if f.attributeToggleEnabled(strategy, user.Attributes) {
				return true
			}
			break
		case collection.ToggleStrategyTypeUUID:
			if f.uuidEnabled(strategy, user.UUID) {
				return true
			}
			break
		case collection.ToggleStrategyTypeReleaseDateTime:
			if f.releaseDateTimeEnabled(strategy.ReleaseDateTime.DateTime) {
				return true
			}
			break
		case collection.ToggleStrategyTypePercentage:
			if f.percentageEnabled(user.UUID, strategy.Percentage.Percent) {
				return true
			}
			break
		}
	}

	return false
}

// groupToggleEnabled returns enabled flag when in the specified groups
func (f *FeatureService) groupToggleEnabled(strategy collection.ToggleStrategy, userGroups []collection.UserGroup) bool {
	for _, userGroup := range userGroups {
		idx := strategy.SearchGroupName(userGroup.Name)
		if idx < len(strategy.Groups) && strategy.Groups[idx].Name == userGroup.Name {
			return true
		}
	}

	return false
}

// attributeToggleEnabled returns enabled flag when in the specified attributes
func (f *FeatureService) attributeToggleEnabled(strategy collection.ToggleStrategy, userAttributes []collection.UserAttribute) bool {
	for _, userAttribute := range userAttributes {
		idx := strategy.SearchAttributeKey(userAttribute.Key)
		if idx < len(strategy.Attributes) && strategy.Attributes[idx].Key == userAttribute.Key && strategy.Attributes[idx].Value == userAttribute.Value {
			return true
		}
	}

	return false
}

// uuidEnabled returns enabled flag when in the specified uuid
func (f *FeatureService) uuidEnabled(strategy collection.ToggleStrategy, UUID string) bool {
	idx := strategy.SearchUUID(UUID)
	if idx < len(strategy.UUIDs) && strategy.UUIDs[idx].UUID == UUID {
		return true
	}

	return false
}

// releaseDateTimeEnabled checks and returns releasable datetime now
func (f *FeatureService) releaseDateTimeEnabled(releaseDateTime time.Time) bool {
	now := time.Now()
	if now.Equal(releaseDateTime) || now.After(releaseDateTime) {
		return true
	}

	return false
}

// percentageEnabled checks and returns releasable user
func (f *FeatureService) percentageEnabled(uuid string, percentage uint32) bool {
	checksum := adler32.Checksum([]byte(uuid))

	return checksum%100000 < percentage*1000
}
