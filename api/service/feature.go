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
func (f *FeatureService) FetchFeature(key string, projectID string) (collection.Feature, error) {
	feature, err := f.featureRepo.FindByKey(key, projectID)
	if err != nil {
		err = NewResourceNotFoundError(err.Error())
	}

	now := time.Now()
	feature.LastActivatedAt = now
	f.featureRepo.Update(&feature)

	return feature, nil
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

	return feature, feature.Enabled && f.enabled(feature, user), nil
}

// enabled verify enabled feature
func (f *FeatureService) enabled(feature collection.Feature, user collection.User) bool {
	if len(feature.Filters) < 1 {
		return true
	}

	for _, filter := range feature.Filters {
		switch filter.Type {
		case collection.ToggleFilterTypeGroup:
			if f.groupToggleEnabled(filter, user.Groups) {
				return true
			}
			break
		case collection.ToggleFilterTypeAttribute:
			if f.attributeToggleEnabled(filter, user.Attributes) {
				return true
			}
			break
		case collection.ToggleFilterTypeUUID:
			if f.uuidEnabled(filter, user.UUID) {
				return true
			}
			break
		case collection.ToggleFilterTypeReleaseDateTime:
			if f.releaseDateTimeEnabled(filter.ReleaseDateTime.DateTime) {
				return true
			}
			break
		case collection.ToggleFilterTypePercentage:
			if f.percentageEnabled(user.UUID, filter.Percentage.Percent) {
				return true
			}
			break
		}
	}

	return false
}

// groupToggleEnabled returns enabled flag when in the specified groups
func (f *FeatureService) groupToggleEnabled(filter collection.ToggleFilter, groups []collection.UserGroup) bool {
	for _, group := range groups {
		idx := filter.SearchGroupName(group.Name)
		if idx < len(filter.Groups) && filter.Groups[idx].Name == group.Name {
			return true
		}
	}

	return false
}

// attributeToggleEnabled returns enabled flag when in the specified attributes
func (f *FeatureService) attributeToggleEnabled(filter collection.ToggleFilter, attributes []collection.UserAttribute) bool {
	for _, attribute := range attributes {
		idx := filter.SearchAttributeKey(attribute.Key)
		if idx < len(filter.Attributes) && filter.Attributes[idx].Key == attribute.Key && filter.Attributes[idx].Value == attribute.Value {
			return true
		}
	}

	return false
}

// uuidEnabled returns enabled flag when in the specified uuid
func (f *FeatureService) uuidEnabled(filter collection.ToggleFilter, UUID string) bool {
	idx := filter.SearchUUID(UUID)
	if idx < len(filter.UUIDs) && filter.UUIDs[idx].UUID == UUID {
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
