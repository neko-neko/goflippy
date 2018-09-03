package collection_test

import (
	"testing"

	. "github.com/neko-neko/goflippy/pkg/collection"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func TestNewFeatureSetValues(t *testing.T) {
	f := NewFeature()

	assert.NotEqual(t, "", f.ID)
	assert.Equal(t, 0, len(f.Filters))
}

func TestFeatureSetProjectIdWhenInvalidIdThenReturnError(t *testing.T) {
	f := NewFeature()

	assert.Error(t, f.SetProjectID("Invalid-ID"))
}

func TestFeatureSetProjectIdWhenValidIdThenReturnNil(t *testing.T) {
	f := NewFeature()
	expected := bson.NewObjectId()

	assert.NoError(t, f.SetProjectID(expected.Hex()))
	assert.Equal(t, expected, f.ProjectID)
}

func TestSearchGroupNameCanBeFoundElem(t *testing.T) {
	cases := []struct {
		input    ToggleFilter
		arg      string
		expected int
	}{
		{
			input: ToggleFilter{
				Type: ToggleFilterTypeGroup,
				Groups: []ToggleFilterGroup{
					ToggleFilterGroup{
						Name: "Group-A",
					},
				},
			},
			arg:      "Group-A",
			expected: 0,
		},
		{
			input: ToggleFilter{
				Type: ToggleFilterTypeGroup,
				Groups: []ToggleFilterGroup{
					ToggleFilterGroup{
						Name: "Group-A",
					},
					ToggleFilterGroup{
						Name: "Group-B",
					},
					ToggleFilterGroup{
						Name: "Group-C",
					},
				},
			},
			arg:      "Group-B",
			expected: 1,
		},
	}

	for _, c := range cases {
		idx := c.input.SearchGroupName(c.arg)
		assert.True(t, idx < len(c.input.Groups))
		assert.Equal(t, c.expected, idx)
		assert.Equal(t, c.input.Groups[idx].Name, c.input.Groups[c.expected].Name)
	}
}

func TestSearchAttributeKeyCanBeFoundElem(t *testing.T) {
	cases := []struct {
		input    ToggleFilter
		arg      string
		expected int
	}{
		{
			input: ToggleFilter{
				Type: ToggleFilterTypeAttribute,
				Attributes: []ToggleFilterAttribute{
					ToggleFilterAttribute{
						Key:   "Key-A",
						Value: "Value-A",
					},
				},
			},
			arg:      "Key-A",
			expected: 0,
		},
		{
			input: ToggleFilter{
				Type: ToggleFilterTypeAttribute,
				Attributes: []ToggleFilterAttribute{
					ToggleFilterAttribute{
						Key:   "Key-A",
						Value: "Value-A",
					},
					ToggleFilterAttribute{
						Key:   "Key-B",
						Value: "Value-B",
					},
					ToggleFilterAttribute{
						Key:   "Key-C",
						Value: "Value-C",
					},
				},
			},
			arg:      "Key-B",
			expected: 1,
		},
	}

	for _, c := range cases {
		idx := c.input.SearchAttributeKey(c.arg)
		assert.True(t, idx < len(c.input.Attributes))
		assert.Equal(t, c.expected, idx)
		assert.Equal(t, c.input.Attributes[idx].Key, c.input.Attributes[c.expected].Key)
		assert.Equal(t, c.input.Attributes[idx].Value, c.input.Attributes[c.expected].Value)
	}
}

func TestSearchUUIDCanBeFoundElem(t *testing.T) {
	cases := []struct {
		input    ToggleFilter
		arg      string
		expected int
	}{
		{
			input: ToggleFilter{
				Type: ToggleFilterTypeUUID,
				UUIDs: []ToggleFilterUUID{
					ToggleFilterUUID{
						UUID: "UUID-A",
					},
				},
			},
			arg:      "UUID-A",
			expected: 0,
		},
		{
			input: ToggleFilter{
				Type: ToggleFilterTypeUUID,
				UUIDs: []ToggleFilterUUID{
					ToggleFilterUUID{
						UUID: "UUID-A",
					},
					ToggleFilterUUID{
						UUID: "UUID-B",
					},
					ToggleFilterUUID{
						UUID: "UUID-C",
					},
				},
			},
			arg:      "UUID-B",
			expected: 1,
		},
	}

	for _, c := range cases {
		idx := c.input.SearchUUID(c.arg)
		assert.True(t, idx < len(c.input.UUIDs))
		assert.Equal(t, c.expected, idx)
		assert.Equal(t, c.input.UUIDs[idx].UUID, c.input.UUIDs[c.expected].UUID)
	}
}

func TestSortCanBeSortElem(t *testing.T) {
	cases := []struct {
		input    ToggleFilter
		expected ToggleFilter
	}{
		{
			input: ToggleFilter{
				Type: ToggleFilterTypeGroup,
				Groups: []ToggleFilterGroup{
					ToggleFilterGroup{
						Name: "Group-A",
					},
					ToggleFilterGroup{
						Name: "Group-C",
					},
					ToggleFilterGroup{
						Name: "Group-B",
					},
				},
			},
			expected: ToggleFilter{
				Type: ToggleFilterTypeGroup,
				Groups: []ToggleFilterGroup{
					ToggleFilterGroup{
						Name: "Group-A",
					},
					ToggleFilterGroup{
						Name: "Group-B",
					},
					ToggleFilterGroup{
						Name: "Group-C",
					},
				},
			},
		},
		{
			input: ToggleFilter{
				Type: ToggleFilterTypeAttribute,
				Attributes: []ToggleFilterAttribute{
					ToggleFilterAttribute{
						Key:   "Key-C",
						Value: "Value-C",
					},
					ToggleFilterAttribute{
						Key:   "Key-B",
						Value: "Value-B",
					},
					ToggleFilterAttribute{
						Key:   "Key-A",
						Value: "Value-A",
					},
				},
			},
			expected: ToggleFilter{
				Type: ToggleFilterTypeGroup,
				Attributes: []ToggleFilterAttribute{
					ToggleFilterAttribute{
						Key:   "Key-A",
						Value: "Value-A",
					},
					ToggleFilterAttribute{
						Key:   "Key-B",
						Value: "Value-B",
					},
					ToggleFilterAttribute{
						Key:   "Key-C",
						Value: "Value-C",
					},
				},
			},
		},
		{
			input: ToggleFilter{
				Type: ToggleFilterTypeUUID,
				UUIDs: []ToggleFilterUUID{
					ToggleFilterUUID{
						UUID: "UUID-B",
					},
					ToggleFilterUUID{
						UUID: "UUID-A",
					},
					ToggleFilterUUID{
						UUID: "UUID-C",
					},
				},
			},
			expected: ToggleFilter{
				Type: ToggleFilterTypeUUID,
				UUIDs: []ToggleFilterUUID{
					ToggleFilterUUID{
						UUID: "UUID-A",
					},
					ToggleFilterUUID{
						UUID: "UUID-B",
					},
					ToggleFilterUUID{
						UUID: "UUID-C",
					},
				},
			},
		},
	}

	for _, c := range cases {
		actual := c.input.Sort()
		switch actual.Type {
		case ToggleFilterTypeGroup:
			for idx, g := range c.expected.Groups {
				assert.Equal(t, g.Name, actual.Groups[idx].Name)
			}
			break
		case ToggleFilterTypeAttribute:
			for idx, a := range c.expected.Attributes {
				assert.Equal(t, a.Key, actual.Attributes[idx].Key)
				assert.Equal(t, a.Value, actual.Attributes[idx].Value)
			}
			break
		case ToggleFilterTypeUUID:
			for idx, u := range c.expected.UUIDs {
				assert.Equal(t, u.UUID, actual.UUIDs[idx].UUID)
			}
			break
		}
	}
}
