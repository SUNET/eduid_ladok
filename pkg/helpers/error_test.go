package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *Error
		want string
	}{
		{
			name: "nil error",
			err:  nil,
			want: "",
		},
		{
			name: "title only",
			err:  &Error{Title: "some_error"},
			want: "Error: [some_error]",
		},
		{
			name: "title and details",
			err:  &Error{Title: "some_error", Details: "detail info"},
			want: "Error: [some_error] detail info",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.err.Error()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewError(t *testing.T) {
	err := NewError("test_error")
	assert.Equal(t, "test_error", err.Title)
	assert.Nil(t, err.Details)
}

func TestNewErrorDetails(t *testing.T) {
	err := NewErrorDetails("test_error", "some details")
	assert.Equal(t, "test_error", err.Title)
	assert.Equal(t, "some details", err.Details)
}

func TestNewErrorFromError(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		got := NewErrorFromError(nil)
		assert.Nil(t, got)
	})

	t.Run("already Error type", func(t *testing.T) {
		original := &Error{Title: "original", Details: "detail"}
		got := NewErrorFromError(original)
		assert.Equal(t, original, got)
	})

	t.Run("json UnmarshalTypeError", func(t *testing.T) {
		type testType struct {
			Age int `json:"age"`
		}
		err := json.Unmarshal([]byte(`{"age":"not_a_number"}`), &testType{})
		got := NewErrorFromError(err)
		assert.Equal(t, "json_type_error", got.Title)
		assert.NotNil(t, got.Details)
	})

	t.Run("json SyntaxError", func(t *testing.T) {
		err := json.Unmarshal([]byte(`{invalid`), &struct{}{})
		got := NewErrorFromError(err)
		assert.Equal(t, "json_syntax_error", got.Title)
		assert.NotNil(t, got.Details)
	})

	t.Run("validation errors", func(t *testing.T) {
		type testStruct struct {
			Name string `validate:"required"`
		}
		validate := validator.New()
		err := validate.Struct(testStruct{})
		got := NewErrorFromError(err)
		assert.Equal(t, "validation_error", got.Title)
		assert.NotNil(t, got.Details)
	})

	t.Run("generic error", func(t *testing.T) {
		err := errors.New("something went wrong")
		got := NewErrorFromError(err)
		assert.Equal(t, "internal_server_error", got.Title)
		assert.Equal(t, "something went wrong", got.Details)
	})
}

func TestFormatResponse(t *testing.T) {
	t.Run("nil response", func(t *testing.T) {
		got := FormatResponse(nil)
		assert.Equal(t, "(no response)", got)
	})

	t.Run("response with body", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader("hello")),
		}
		got := FormatResponse(resp)
		assert.Equal(t, "(status=200, body=hello)", got)
	})

	t.Run("response without body", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 404,
			Body:       io.NopCloser(strings.NewReader("")),
		}
		got := FormatResponse(resp)
		assert.Equal(t, "(status=404)", got)
	})
}

func TestProblem404(t *testing.T) {
	problem := Problem404()
	assert.NotNil(t, problem)
	assert.Equal(t, 404, problem.Status)
}
