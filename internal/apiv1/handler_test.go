package apiv1

import (
	"eduid_ladok/internal/ladok"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SUNET/goladok3"
	"github.com/SUNET/goladok3/ladokmocks"
	"github.com/SUNET/goladok3/ladoktypes"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestSchoolInfo(t *testing.T) {
	tts := []struct {
		name       string
		schools    map[string]struct{ SwamidName string `yaml:"swamid_name" validate:"required"` }
		schoolInfo map[string]model.SchoolInfo
		wantKeys   []string
		dontWant   []string
	}{
		{
			name: "school with info returned",
			schools: map[string]struct{ SwamidName string `yaml:"swamid_name" validate:"required"` }{
				"testSchool":  {SwamidName: "test"},
				"otherSchool": {SwamidName: "other"},
			},
			schoolInfo: map[string]model.SchoolInfo{
				"testSchool": {LongNameSv: "Testskolan", LongNameEn: "Test School"},
			},
			wantKeys: []string{"testSchool"},
			dontWant: []string{"otherSchool"},
		},
		{
			name:    "no schools",
			schools: nil,
			wantKeys: nil,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &model.Cfg{}
			cfg.Schools = tt.schools
			cfg.SchoolInformation = tt.schoolInfo

			c := &Client{config: cfg}

			reply, err := c.SchoolInfo(t.Context(), &RequestSchoolInfo{})
			assert.NoError(t, err)
			assert.NotNil(t, reply)
			for _, key := range tt.wantKeys {
				assert.Contains(t, reply.Schools, key)
			}
			for _, key := range tt.dontWant {
				assert.NotContains(t, reply.Schools, key)
			}
		})
	}
}

func TestLadokInfo_NoMatchingInstance(t *testing.T) {
	tts := []struct {
		name       string
		schoolName string
		wantErr    string
	}{
		{
			name:       "nonexistent school",
			schoolName: "nonexistent",
			wantErr:    "can't find any matching ladok instance",
		},
		{
			name:       "empty school name",
			schoolName: "",
			wantErr:    "can't find any matching ladok instance",
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{}
			_, err := c.LadokInfo(t.Context(), &RequestLadokInfo{SchoolName: tt.schoolName})
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestMonitoringCertClient(t *testing.T) {
	tts := []struct {
		name           string
		ladokInstances map[string]*ladok.Service
		wantLen        int
	}{
		{
			name:           "empty instances",
			ladokInstances: nil,
			wantLen:        0,
		},
		{
			name: "with certificate status",
			ladokInstances: map[string]*ladok.Service{
				"school1": {
					Certificate: &ladok.CertificateService{
						ClientCertificateStatus: &model.MonitoringCertClient{
							Valid:       true,
							Fingerprint: "abc",
						},
					},
				},
			},
			wantLen: 1,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{ladokInstances: tt.ladokInstances}
			reply, err := c.MonitoringCertClient(t.Context())
			assert.NoError(t, err)
			assert.NotNil(t, reply)
			assert.Len(t, *reply, tt.wantLen)
		})
	}
}

func TestStatus(t *testing.T) {
	tts := []struct {
		name           string
		ladokInstances map[string]*ladok.Service
		wantHealthy    bool
	}{
		{
			name:           "empty instances - healthy",
			ladokInstances: nil,
			wantHealthy:    true,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{ladokInstances: tt.ladokInstances}
			reply, err := c.Status(t.Context())
			assert.NoError(t, err)
			assert.NotNil(t, reply)
			assert.Equal(t, tt.wantHealthy, reply.Healthy)
		})
	}
}

func TestLadokInfo_Success(t *testing.T) {
	// Set up mock Ladok HTTP server for studentinformation endpoint
	mux := http.NewServeMux()
	for _, student := range ladokmocks.Students {
		payload := ladokmocks.StudentJSON(student)
		mux.HandleFunc("/studentinformation/student/personnummer/"+student.Personnummer,
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", goladok3.ContentTypeKataloginformationJSON)
				w.WriteHeader(200)
				w.Write(payload)
			},
		)
	}
	server := httptest.NewServer(mux)
	defer server.Close()

	// Create goladok3 client with mock certificate
	certPEM, cert, keyPEM, _ := ladokmocks.MockCertificateAndKey(t, ladoktypes.EnvIntTestAPI, 0, 100)
	goladokClient, err := goladok3.NewX509(goladok3.X509Config{
		URL:            server.URL,
		Certificate:    cert,
		CertificatePEM: certPEM,
		PrivateKeyPEM:  keyPEM,
	})
	assert.NoError(t, err)

	svc := &ladok.Service{
		Rest: &ladok.RestService{
			Ladok: goladokClient,
		},
	}

	c := &Client{
		ladokInstances: map[string]*ladok.Service{
			"testSchool": svc,
		},
	}

	tts := []struct {
		name       string
		schoolName string
		nin        string
		wantErr    bool
	}{
		{
			name:       "valid student lookup",
			schoolName: "testSchool",
			nin:        ladokmocks.Students[0].Personnummer,
			wantErr:    false,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			reply, err := c.LadokInfo(t.Context(), &RequestLadokInfo{
				SchoolName: tt.schoolName,
				Data:       model.UserData{NIN: tt.nin},
			})
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, reply)
				assert.NotEmpty(t, reply.ESI)
				assert.NotEmpty(t, reply.LadokExterntUID)
			}
		})
	}
}

func TestNew(t *testing.T) {
	log := &logger.Logger{Logger: *zaptest.NewLogger(t, zaptest.Level(zap.PanicLevel))}
	cfg := &model.Cfg{}

	client, err := New(t.Context(), cfg, nil, log)
	assert.NoError(t, err)
	assert.NotNil(t, client)
}
