package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("production", func(t *testing.T) {
		log := New("test", true)
		assert.NotNil(t, log)
	})

	t.Run("development", func(t *testing.T) {
		log := New("test", false)
		assert.NotNil(t, log)
	})
}

func TestNewSimple(t *testing.T) {
	log := NewSimple("test")
	assert.NotNil(t, log)
}

func TestLogger_New(t *testing.T) {
	log := New("parent", false)
	child := log.New("child")
	assert.NotNil(t, child)
}

func TestLogger_Methods(t *testing.T) {
	log := New("test", false)

	// These should not panic
	log.Info("info message", "key", "value")
	log.Warn("warn message", "key", "value")
	log.Error("error message", "key", "value")
	log.Debug("debug message", "key", "value")
}
