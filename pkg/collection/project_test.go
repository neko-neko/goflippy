package collection_test

import (
	"testing"

	. "github.com/neko-neko/goflippy/pkg/collection"
	"github.com/stretchr/testify/assert"
)

func TestNewProjectSetValues(t *testing.T) {
	p := NewProject()

	assert.NotEqual(t, "", p.ID)
	assert.Equal(t, 0, len(p.APIKeys))
	assert.False(t, p.CreatedAt.IsZero())
	assert.False(t, p.UpdatedAt.IsZero())
}
