package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextKey_String(t *testing.T) {
	key := ContextKey("test-key")
	assert.Equal(t, "test-key", key.String())
}
