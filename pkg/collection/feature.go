package collection

import (
	"fmt"
	"sort"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Toggle Filter definitions
const (
	ToggleFilterTypeGroup           = "group"
	ToggleFilterTypeAttribute       = "attribute"
	ToggleFilterTypeUUID            = "uuid"
	ToggleFilterTypeReleaseDateTime = "release_date_time"
	ToggleFilterTypePercentage      = "percentage"
)

// Feature is nested structure
type Feature struct {
	ID              bson.ObjectId  `json:"_id" bson:"_id"`
	ProjectID       bson.ObjectId  `json:"project_id" bson:"project_id"`
	Key             string         `json:"key" bson:"key"`
	Name            string         `json:"name" bson:"name"`
	Enabled         bool           `json:"enabled" bson:"enabled"`
	LastActivatedAt time.Time      `json:"last_activated_at" bson:"last_activated_at"`
	Filters         []ToggleFilter `json:"filters" bson:"filters"`
}

// ToggleFilterGroup is feature publish target group
type ToggleFilterGroup struct {
	Name string `json:"name" bson:"name"`
}

// ToggleFilterAttribute is feature publish target attribute
type ToggleFilterAttribute struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}

// ToggleFilterUUID is feature publish specification UUID
type ToggleFilterUUID struct {
	UUID string `json:"uuid" bson:"uuid"`
}

// ToggleFilterReleaseDateTime is feature publish release datetime
type ToggleFilterReleaseDateTime struct {
	DateTime time.Time `json:"date_time" bson:"date_time,omitempty"`
}

// ToggleFilterPercentage is feature publish percentage
type ToggleFilterPercentage struct {
	Percent uint32 `json:"percent" bson:"percent,omitempty"`
}

// ToggleFilter is filter of feature toggle
type ToggleFilter struct {
	Type            string                      `json:"type" bson:"type"`
	Groups          []ToggleFilterGroup         `json:"groups" bson:"groups,omitempty"`
	Attributes      []ToggleFilterAttribute     `json:"attributes" bson:"attributes,omitempty"`
	UUIDs           []ToggleFilterUUID          `json:"uuids" bson:"uuids,omitempty"`
	ReleaseDateTime ToggleFilterReleaseDateTime `json:"release_date_time" bson:"release_date_time,omitempty"`
	Percentage      ToggleFilterPercentage      `json:"percentage" bson:"percentage,omitempty"`
}

// NewFeature returns feature object
func NewFeature() *Feature {
	return &Feature{
		ID:      bson.NewObjectId(),
		Filters: make([]ToggleFilter, 0),
	}
}

// SetProjectID is set string projectID to bson.ObjectId projectID
func (f *Feature) SetProjectID(projectID string) error {
	if !bson.IsObjectIdHex(projectID) {
		return fmt.Errorf("%s is not valid format", projectID)
	}

	f.ProjectID = bson.ObjectIdHex(projectID)
	return nil
}

// SearchGroupName returns index of x into ToggleStrategyGroup
// SearchGroupName using binary search algorithm, so must sort t.Groups
func (t ToggleFilter) SearchGroupName(x string) int {
	return sort.Search(len(t.Groups), func(i int) bool { return t.Groups[i].Name >= x })
}

// SearchAttributeKey returns index of x into ToggleStrategyAttribute
// SearchAttributeKey using binary search algorithm, so must sort t.Attributes
func (t ToggleFilter) SearchAttributeKey(x string) int {
	return sort.Search(len(t.Attributes), func(i int) bool { return t.Attributes[i].Key >= x })
}

// SearchUUID returns index of x into ToggleStrategyUUID
// SearchUUID using binary search algorithm, so must sort t.UUIDs
func (t ToggleFilter) SearchUUID(x string) int {
	return sort.Search(len(t.UUIDs), func(i int) bool { return t.UUIDs[i].UUID >= x })
}

// Sort elements
func (t ToggleFilter) Sort() ToggleFilter {
	switch t.Type {
	case ToggleFilterTypeGroup:
		sort.Slice(t.Groups, func(i, j int) bool {
			return t.Groups[i].Name < t.Groups[j].Name
		})
		break
	case ToggleFilterTypeAttribute:
		sort.Slice(t.Attributes, func(i, j int) bool {
			return t.Attributes[i].Key < t.Attributes[j].Key
		})
		break
	case ToggleFilterTypeUUID:
		sort.Slice(t.UUIDs, func(i, j int) bool {
			return t.UUIDs[i].UUID < t.UUIDs[j].UUID
		})
		break
	}

	return t
}
