package collection

import (
	"sort"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Toggle Strategy definitions
const (
	ToggleStrategyTypeSimple          = "simple"
	ToggleStrategyTypeGroup           = "group"
	ToggleStrategyTypeAttribute       = "attribute"
	ToggleStrategyTypeUUID            = "uuid"
	ToggleStrategyTypeReleaseDateTime = "release_date_time"
	ToggleStrategyTypePercentage      = "percentage"
)

// Feature is nested structure
type Feature struct {
	ID              bson.ObjectId    `json:"_id" bson:"_id"`
	ProjectID       bson.ObjectId    `json:"project_id" bson:"project_id"`
	Key             string           `json:"key" bson:"key"`
	Name            string           `json:"name" bson:"name"`
	LastActivatedAt time.Time        `json:"last_activated_at" bson:"last_activated_at"`
	Strategies      []ToggleStrategy `json:"strategies" bson:"strategies"`
}

// ToggleStrategy is toggle strategy
type ToggleStrategy struct {
	Type            string                        `json:"type" bson:"type"`
	Enable          bool                          `json:"enable" bson:"enable,omitempty"`
	Groups          []ToggleStrategyGroup         `json:"groups" bson:"groups,omitempty"`
	Attributes      []ToggleStrategyAttribute     `json:"attributes" bson:"attributes,omitempty"`
	UUIDs           []ToggleStrategyUUID          `json:"uuids" bson:"uuids,omitempty"`
	ReleaseDateTime ToggleStrategyReleaseDateTime `json:"release_date_time" bson:"release_date_time,omitempty"`
	Percentage      ToggleStrategyPercentage      `json:"percentage" bson:"percentage,omitempty"`
}

// SearchGroupName returns index of x into ToggleStrategyGroup
// SearchGroupName using binary search algorithm, so must sort t.Groups
func (t ToggleStrategy) SearchGroupName(x string) int {
	return sort.Search(len(t.Groups), func(i int) bool { return t.Groups[i].Name >= x })
}

// SearchAttributeKey returns index of x into ToggleStrategyAttribute
// SearchAttributeKey using binary search algorithm, so must sort t.Attributes
func (t ToggleStrategy) SearchAttributeKey(x string) int {
	return sort.Search(len(t.Attributes), func(i int) bool { return t.Attributes[i].Key >= x })
}

// SearchUUID returns index of x into ToggleStrategyUUID
// SearchUUID using binary search algorithm, so must sort t.UUIDs
func (t ToggleStrategy) SearchUUID(x string) int {
	return sort.Search(len(t.UUIDs), func(i int) bool { return t.UUIDs[i].UUID >= x })
}

// Sort elements
func (t ToggleStrategy) Sort() ToggleStrategy {
	switch t.Type {
	case ToggleStrategyTypeGroup:
		sort.Slice(&t.Groups, func(i, j int) bool {
			return t.Groups[i].Name > t.Groups[j].Name
		})
		break
	case ToggleStrategyTypeAttribute:
		sort.Slice(&t.Attributes, func(i, j int) bool {
			return t.Attributes[i].Key > t.Attributes[j].Key
		})
		break
	case ToggleStrategyTypeUUID:
		sort.Slice(&t.UUIDs, func(i, j int) bool {
			return t.UUIDs[i].UUID > t.UUIDs[j].UUID
		})
		break
	}

	return t
}

// ToggleStrategyGroup is feature publish target group
type ToggleStrategyGroup struct {
	Name string `json:"name" bson:"name"`
}

// ToggleStrategyAttribute is feature publish target attribute
type ToggleStrategyAttribute struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}

// ToggleStrategyUUID is feature publish specification UUID
type ToggleStrategyUUID struct {
	UUID string `json:"uuid" bson:"uuid"`
}

// ToggleStrategyReleaseDateTime is feature publish release datetime
type ToggleStrategyReleaseDateTime struct {
	DateTime time.Time `json:"date_time" bson:"date_time,omitempty"`
}

// ToggleStrategyPercentage is feature publish percentage
type ToggleStrategyPercentage struct {
	Percent uint32 `json:"percent" bson:"percent,omitempty"`
}
