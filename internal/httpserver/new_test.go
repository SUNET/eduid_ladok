package httpserver

import (
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestNewService(t *testing.T) {
	tts := []struct {
		name       string
		production bool
	}{
		{name: "development mode", production: false},
		{name: "production mode", production: true},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			log := &logger.Logger{Logger: *zaptest.NewLogger(t, zaptest.Level(zap.PanicLevel))}
			cfg := &model.Cfg{Production: tt.production}
			cfg.APIServer.Host = "127.0.0.1:0"

			s, err := New(t.Context(), cfg, nil, log)
			assert.NoError(t, err)
			assert.NotNil(t, s)

			err = s.Close(t.Context())
			assert.NoError(t, err)
		})
	}
}
