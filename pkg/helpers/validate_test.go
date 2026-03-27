package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	type testStruct struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}

	tts := []struct {
		name    string
		have    testStruct
		wantErr bool
	}{
		{
			name:    "valid struct",
			have:    testStruct{Name: "test", Email: "test@example.com"},
			wantErr: false,
		},
		{
			name:    "missing all fields",
			have:    testStruct{},
			wantErr: true,
		},
		{
			name:    "invalid email",
			have:    testStruct{Name: "test", Email: "not-an-email"},
			wantErr: true,
		},
		{
			name:    "missing name",
			have:    testStruct{Email: "test@example.com"},
			wantErr: true,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			err := Check(tt.have, nil)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
