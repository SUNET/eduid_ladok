package httpserver

import (
	"context"
	"eduid_ladok/internal/apiv1"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

type mockApiv1 struct {
	schoolInfo  *apiv1.ReplySchoolInfo
	status      *model.Status
	certClients *model.MonitoringCertClients
	ladokInfo   *apiv1.ReplyLadokInfo
	err         error
}

func (m *mockApiv1) LadokInfo(_ context.Context, _ *apiv1.RequestLadokInfo) (*apiv1.ReplyLadokInfo, error) {
	return m.ladokInfo, m.err
}

func (m *mockApiv1) SchoolInfo(_ context.Context, _ *apiv1.RequestSchoolInfo) (*apiv1.ReplySchoolInfo, error) {
	return m.schoolInfo, m.err
}

func (m *mockApiv1) Status(_ context.Context) (*model.Status, error) {
	return m.status, m.err
}

func (m *mockApiv1) MonitoringCertClient(_ context.Context) (*model.MonitoringCertClients, error) {
	return m.certClients, m.err
}

func testLogger(t *testing.T) *logger.Logger {
	return &logger.Logger{Logger: *zaptest.NewLogger(t, zaptest.Level(zap.PanicLevel))}
}

func setupTestService(t *testing.T, api Apiv1) *Service {
	gin.SetMode(gin.TestMode)
	log := testLogger(t)
	s := &Service{
		config: &model.Cfg{},
		logger: log,
		apiv1:  api,
		gin:    gin.New(),
	}

	ctx := t.Context()
	s.gin.Use(s.middlewareTraceID(ctx))
	s.gin.Use(s.middlewareDuration(ctx))
	s.gin.Use(s.middlewareCrash(ctx))

	s.regEndpoint(ctx, "api/v1/:schoolName/ladokinfo", "POST", s.endpointLadokInfo)
	s.regEndpoint(ctx, "api/v1/schoolinfo", "GET", s.endpointSchoolInfo)
	s.regEndpoint(ctx, "/health", "GET", s.endpointStatus)
	s.regEndpoint(ctx, "/monitoring/cert/client", "GET", s.endpointMonitoringCertClient)

	return s
}

func TestEndpointSchoolInfo(t *testing.T) {
	tts := []struct {
		name       string
		mock       *mockApiv1
		wantStatus int
		wantSchool string
	}{
		{
			name: "returns school info",
			mock: &mockApiv1{
				schoolInfo: &apiv1.ReplySchoolInfo{
					Schools: map[string]model.SchoolInfo{
						"testSchool": {LongNameSv: "Testskolan", LongNameEn: "Test School"},
					},
				},
			},
			wantStatus: 200,
			wantSchool: "testSchool",
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			s := setupTestService(t, tt.mock)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/schoolinfo", nil)
			req.Header.Set("Accept", "application/json")
			s.gin.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantSchool)
		})
	}
}

func TestEndpointStatus(t *testing.T) {
	tts := []struct {
		name       string
		mock       *mockApiv1
		wantStatus int
		wantHealth bool
	}{
		{
			name: "healthy status",
			mock: &mockApiv1{
				status: &model.Status{Healthy: true, Status: model.StatusOK, Timestamp: time.Now()},
			},
			wantStatus: 200,
			wantHealth: true,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			s := setupTestService(t, tt.mock)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			req.Header.Set("Accept", "application/json")
			s.gin.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestEndpointMonitoringCertClient(t *testing.T) {
	clients := model.MonitoringCertClients{
		"school1": {Valid: true, Fingerprint: "abc123"},
	}

	tts := []struct {
		name       string
		mock       *mockApiv1
		wantStatus int
	}{
		{
			name:       "returns cert clients",
			mock:       &mockApiv1{certClients: &clients},
			wantStatus: 200,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			s := setupTestService(t, tt.mock)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/monitoring/cert/client", nil)
			req.Header.Set("Accept", "application/json")
			s.gin.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), "abc123")
		})
	}
}

func TestEndpointLadokInfo(t *testing.T) {
	tts := []struct {
		name       string
		mock       *mockApiv1
		body       string
		wantStatus int
	}{
		{
			name: "successful request",
			mock: &mockApiv1{
				ladokInfo: &apiv1.ReplyLadokInfo{
					ESI:             "urn:schac:personalUniqueCode:int:esi:ladok.se:externtstudentuid-abc",
					LadokExterntUID: "abc",
					IsStudent:       false,
				},
			},
			body:       `{"data":{"nin":"199001011234"}}`,
			wantStatus: 200,
		},
		{
			name:       "error from api",
			mock:       &mockApiv1{err: assert.AnError},
			body:       `{"data":{"nin":"199001011234"}}`,
			wantStatus: 400,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			s := setupTestService(t, tt.mock)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/testSchool/ladokinfo", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "application/json")
			s.gin.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestRenderContent(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tts := []struct {
		name       string
		accept     string
		wantStatus int
	}{
		{
			name:       "json response",
			accept:     "application/json",
			wantStatus: 200,
		},
		{
			name:       "wildcard accept",
			accept:     "*/*",
			wantStatus: 200,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
			c.Request.Header.Set("Accept", tt.accept)

			renderContent(c, tt.wantStatus, gin.H{"data": "test"})

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestMiddlewareTraceID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	log := testLogger(t)
	s := &Service{logger: log}
	engine := gin.New()
	engine.Use(s.middlewareTraceID(t.Context()))
	engine.GET("/test", func(c *gin.Context) {
		id := c.GetString("sunet-request-id")
		c.JSON(200, gin.H{"id": id})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, w.Header().Get("sunet-request-id"))

	var body map[string]string
	json.Unmarshal(w.Body.Bytes(), &body)
	assert.NotEmpty(t, body["id"])
}

func TestMiddlewareDuration(t *testing.T) {
	gin.SetMode(gin.TestMode)
	log := testLogger(t)
	s := &Service{logger: log}
	engine := gin.New()
	engine.Use(s.middlewareDuration(t.Context()))
	engine.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestMiddlewareCrash(t *testing.T) {
	gin.SetMode(gin.TestMode)
	log := testLogger(t)
	s := &Service{logger: log}
	engine := gin.New()
	engine.Use(s.middlewareCrash(t.Context()))
	engine.GET("/panic", func(c *gin.Context) {
		panic("test panic")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Contains(t, w.Body.String(), "internal_server_error")
}

func TestMiddlewareLogger(t *testing.T) {
	gin.SetMode(gin.TestMode)
	log := testLogger(t)
	s := &Service{logger: log}
	engine := gin.New()
	engine.Use(s.middlewareLogger(t.Context()))
	engine.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestClose(t *testing.T) {
	log := testLogger(t)
	s := &Service{logger: log}
	err := s.Close(t.Context())
	assert.NoError(t, err)
}
