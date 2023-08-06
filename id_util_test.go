package utils

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRandomID(t *testing.T) {
	id := RandomID("foo")
	assert.True(t, strings.HasPrefix(id, "foo"))
	t.Log(id)
}
