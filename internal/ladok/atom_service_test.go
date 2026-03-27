package ladok

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusRedis(t *testing.T) {
	tts := []struct {
		name        string
		mockPong    bool
		wantHealthy bool
		wantStatus  string
	}{
		{
			name:        "redis not reachable",
			mockPong:    false,
			wantHealthy: false,
			wantStatus:  "STATUS_FAIL_eduid_ladok_",
		},
		{
			name:        "redis healthy (mock pong)",
			mockPong:    true,
			wantHealthy: true,
			wantStatus:  "STATUS_OK_eduid_ladok_",
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			service, server, redisMock, _ := mockService(t, 200, 0, 100, t.TempDir())
			defer server.Close()

			if tt.mockPong {
				redisMock.ExpectPing().SetVal("PONG")
			}

			status := service.Atom.StatusRedis(t.Context())
			assert.Equal(t, "redis", status.Name)
			assert.Equal(t, tt.wantHealthy, status.Healthy)
			assert.Equal(t, tt.wantStatus, status.Status)
		})
	}
}
