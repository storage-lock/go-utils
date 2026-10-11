package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomID(t *testing.T) {
	id := RandomID("foo")
	assert.True(t, strings.HasPrefix(id, "foo"))
	t.Log(id)
}

func TestRandomID_NoPrefix(t *testing.T) {
	id := RandomID()
	assert.NotEmpty(t, id)
	assert.NotContains(t, id, "-")
	assert.Len(t, id, 32)
	other := RandomID()
	assert.NotEqual(t, id, other)
}

func TestRandomID_EmptyPrefix(t *testing.T) {
	id := RandomID("")
	assert.NotEmpty(t, id)
	assert.Len(t, id, 32)
}
