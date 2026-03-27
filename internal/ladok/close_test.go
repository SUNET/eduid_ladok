package ladok

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCertificateServiceClose(t *testing.T) {
	service, server, _, _ := mockService(t, 200, 0, 100, t.TempDir())
	defer server.Close()

	err := service.Certificate.Close(t.Context())
	assert.NoError(t, err)
}
