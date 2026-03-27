package ladok

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSchoolID(t *testing.T) {
	tts := []struct {
		name       string
		statusCode int
		wantErr    bool
		wantID     int
	}{
		{
			name:       "successful school ID retrieval",
			statusCode: 200,
			wantErr:    false,
			wantID:     96,
		},
		{
			name:       "server error",
			statusCode: 500,
			wantErr:    true,
			wantID:     0,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			service, server, _, _ := mockService(t, tt.statusCode, 0, 100, t.TempDir())
			defer server.Close()

			err := service.getSchoolID(t.Context())
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantID, service.SchoolID)
			}
		})
	}
}
