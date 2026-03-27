package ladok

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusLadok(t *testing.T) {
	tts := []struct {
		name        string
		statusCode  int
		wantHealthy bool
	}{
		{
			name:        "healthy ladok",
			statusCode:  200,
			wantHealthy: true,
		},
		{
			name:        "unhealthy ladok",
			statusCode:  500,
			wantHealthy: false,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			service, server, _, _ := mockService(t, tt.statusCode, 0, 100, t.TempDir())
			defer server.Close()

			status := service.Rest.StatusLadok(t.Context())
			assert.Equal(t, tt.wantHealthy, status.Healthy)
			assert.Equal(t, "Ladok rest", status.Name)
		})
	}
}

func TestRestServiceClose(t *testing.T) {
	service, server, _, _ := mockService(t, 200, 0, 100, t.TempDir())
	defer server.Close()

	err := service.Rest.Close(t.Context())
	assert.NoError(t, err)
}
