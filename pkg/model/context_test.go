package model

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCopyTraceID(t *testing.T) {
	ginContext := &gin.Context{
		Keys: map[any]any{
			"sunet-request-id": "test-uuid",
		},
	}

	ctx := CopyTraceID(t.Context(), ginContext)
	assert.Equal(t, "test-uuid", ctx.Value(ContextKey("sunet-request-id")))
}
