package configuration

import (
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"gopkg.in/yaml.v3"
)

var mockConfig = []byte(`
---
eduid:
  worker:
    common:
      debug: yes
    ladok-x:
      api_server:
        host: :8080
    ladok:
      api_server:
        host: :8080
      production: false
      tracing:
        kind: jaeger
        endpoint: "http://localhost:14268/api/traces"
      schools:
        kf: 
          saml_name: student.konstfack.se
        lnu:
          saml_name: lnu.se 
      ladok:
        url: https://api.integrationstest.ladok.se
        certificate:
          folder: cert
        atom:
          periodicity: 60 
        permissions:
          90019: "rattighetsniva.las"
          51001: "rattighetsniva.las"
          61001: "rattighetsniva.las"
          11004: "rattighetsniva.las"
          860131: "rattighetsniva.las"
      eduid:
        iam:
          url: https://api.dev.eduid.se/scim/test 
      sunet:
        auth:
          url: https://auth-test.sunet.se 
      redis:
        db: 3
        host: localhost:6379
        sentinel_hosts:
        #  - localhost:1231
        #  - localhost:12313
        sentinel_service_name: redis-cluster
    x_service:
      api_server:
        host: 8080
`)

func testLog(t *testing.T) *logger.Logger {
	return &logger.Logger{
		Logger: *zaptest.NewLogger(t, zaptest.Level(zap.PanicLevel)),
	}
}

func TestParse(t *testing.T) {
	tempDir := t.TempDir()

	tts := []struct {
		name           string
		setEnvVariable bool
	}{
		{
			name:           "OK",
			setEnvVariable: true,
		},
	}

	for _, tt := range tts {
		path := fmt.Sprintf("%s/test.cfg", tempDir)
		if err := os.WriteFile(path, mockConfig, 0666); err != nil {
			assert.NoError(t, err)
		}
		if tt.setEnvVariable {
			os.Setenv("EDUID_CONFIG_YAML", path)
		}

		want := &model.Config{}
		err := yaml.Unmarshal(mockConfig, want)
		assert.NoError(t, err)

		t.Run(tt.name, func(t *testing.T) {
			cfg, err := Parse(testLog(t))
			assert.NoError(t, err)

			assert.Equal(t, &want.EduID.Worker.Ladok, cfg)

		})
	}

}

func TestParse_Errors(t *testing.T) {
	tts := []struct {
		name    string
		envVal  string
		setup   func(t *testing.T) string
		wantErr string
	}{
		{
			name:    "missing env variable",
			envVal:  "",
			setup:   func(t *testing.T) string { return "" },
			wantErr: "required",
		},
		{
			name:   "file does not exist",
			envVal: "set",
			setup: func(t *testing.T) string {
				return "/tmp/nonexistent_config_file_12345.yaml"
			},
			wantErr: "no such file",
		},
		{
			name:   "path is a directory",
			envVal: "set",
			setup: func(t *testing.T) string {
				return t.TempDir()
			},
			wantErr: "is a directory",
		},
		{
			name:   "invalid yaml",
			envVal: "set",
			setup: func(t *testing.T) string {
				path := fmt.Sprintf("%s/bad.cfg", t.TempDir())
				os.WriteFile(path, []byte("{{invalid yaml"), 0666)
				return path
			},
			wantErr: "",
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.setup(t)
			if tt.envVal == "" {
				os.Unsetenv("EDUID_CONFIG_YAML")
			} else {
				os.Setenv("EDUID_CONFIG_YAML", path)
			}

			_, err := Parse(testLog(t))
			assert.Error(t, err)
			if tt.wantErr != "" {
				assert.Contains(t, err.Error(), tt.wantErr)
			}
		})
	}
}
