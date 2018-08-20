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

// FeatureEnabledByParams verify enabled feature
func (f *FeatureService) FeatureEnabled(key string, projectID string) (collection.Feature, bool, error) {
	feature, err := f.featureRepo.FindByKey(key, projectID)
	if err != nil {
		err = NewResourceNotFoundError(err.Error())
	}

	now := time.Now()
	feature.LastActivatedAt = now
	f.featureRepo.Update(&feature)

	return feature, false, nil
}

// FeatureEnabledByUser verify enabled feature from a user
func (f *FeatureService) FeatureEnabledByUser(key string, projectID string, uuid string) (collection.Feature, bool, error) {
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

// enabled verify enabled feature
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
func (f *FeatureService) groupToggleEnabled(strategy collection.ToggleStrategy, groups []collection.UserGroup) bool {
	for _, group := range groups {
		idx := strategy.SearchGroupName(group.Name)
		if idx < len(strategy.Groups) && strategy.Groups[idx].Name == group.Name {
			return true
		}
	}

	return false
}

// attributeToggleEnabled returns enabled flag when in the specified attributes
func (f *FeatureService) attributeToggleEnabled(strategy collection.ToggleStrategy, attributes []collection.UserAttribute) bool {
	for _, attribute := range attributes {
		idx := strategy.SearchAttributeKey(attribute.Key)
		if idx < len(strategy.Attributes) && strategy.Attributes[idx].Key == attribute.Key && strategy.Attributes[idx].Value == attribute.Value {
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
