package collection_test

import (
	"testing"

	. "github.com/neko-neko/goflippy/pkg/collection"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func TestNewUserSetValues(t *testing.T) {
	u := NewUser()

	assert.NotEqual(t, "", u.ID)
	assert.Equal(t, 0, len(u.Attributes))
	assert.Equal(t, 0, len(u.Groups))
	assert.False(t, u.LastActivatedAt.IsZero())
	assert.False(t, u.CreatedAt.IsZero())
	assert.False(t, u.UpdatedAt.IsZero())
}

func TestUserSetProjectIdWhenInvalidIdThenReturnError(t *testing.T) {
	u := NewUser()

	assert.Error(t, u.SetProjectID("Invalid-ID"))
}

func TestUserSetProjectIdWhenValidIdThenReturnNil(t *testing.T) {
	u := NewUser()
	expected := bson.NewObjectId()

	assert.NoError(t, u.SetProjectID(expected.Hex()))
	assert.Equal(t, expected, u.ProjectID)
}

func TestAppendGroupCanBeSetValue(t *testing.T) {
	cases := [][]UserGroup{
		[]UserGroup{
			UserGroup{
				Name: "Group-A",
			},
		},
		[]UserGroup{
			UserGroup{
				Name: "Group-A",
			},
			UserGroup{
				Name: "Group-B",
			},
		},
	}

	for _, c := range cases {
		u := NewUser()
		for _, expected := range c {
			u.AppendGroup(expected.Name)
		}

		for idx, g := range u.Groups {
			assert.Equal(t, c[idx].Name, g.Name)
		}
	}
}

func TestAppendAttributeCanBeSetValue(t *testing.T) {
	cases := [][]UserAttribute{
		[]UserAttribute{
			UserAttribute{
				Key:   "Key-A",
				Value: "Value-A",
			},
		},
		[]UserAttribute{
			UserAttribute{
				Key:   "Key-A",
				Value: "Value-A",
			},
			UserAttribute{
				Key:   "Key-B",
				Value: "Value-B",
			},
		},
	}

	for _, c := range cases {
		u := NewUser()
		for _, expected := range c {
			u.AppendAttribute(expected.Key, expected.Value)
		}

		for idx, a := range u.Attributes {
			assert.Equal(t, c[idx].Key, a.Key)
			assert.Equal(t, c[idx].Value, a.Value)
		}
	}
}

func TestResetGroupsCanBeResetGroup(t *testing.T) {
	u := NewUser()

	u.AppendGroup("TEST")
	u.ResetGroups()
	assert.Equal(t, 0, len(u.Groups))
}

func TestResetAttributesCanBeResetAttribute(t *testing.T) {
	u := NewUser()

	u.AppendAttribute("Key", "Value")
	u.ResetAttributes()
	assert.Equal(t, 0, len(u.Attributes))
}
