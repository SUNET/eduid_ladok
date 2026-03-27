package httpserver

import (
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestDefaultValidator_ValidateStruct(t *testing.T) {
	v := &defaultValidator{Validate: validator.New()}

	type testStruct struct {
		Name string `validate:"required"`
	}

	t.Run("valid struct", func(t *testing.T) {
		err := v.ValidateStruct(testStruct{Name: "test"})
		assert.NoError(t, err)
	})

	t.Run("invalid struct", func(t *testing.T) {
		err := v.ValidateStruct(testStruct{})
		assert.Error(t, err)
	})

	t.Run("pointer to valid struct", func(t *testing.T) {
		err := v.ValidateStruct(&testStruct{Name: "test"})
		assert.NoError(t, err)
	})

	t.Run("non-struct type", func(t *testing.T) {
		err := v.ValidateStruct("not a struct")
		assert.NoError(t, err)
	})
}

func TestDefaultValidator_Engine(t *testing.T) {
	validate := validator.New()
	v := &defaultValidator{Validate: validate}
	assert.Equal(t, validate, v.Engine())
}

func TestKindOfData(t *testing.T) {
	tests := []struct {
		name string
		data any
		want reflect.Kind
	}{
		{"struct", struct{}{}, reflect.Struct},
		{"pointer to struct", &struct{}{}, reflect.Struct},
		{"string", "hello", reflect.String},
		{"int", 42, reflect.Int},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kindOfData(tt.data)
			assert.Equal(t, tt.want, got)
		})
	}
}
