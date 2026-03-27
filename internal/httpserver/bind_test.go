package httpserver

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"eduid_ladok/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func bindTestService(t *testing.T) *Service {
	gin.SetMode(gin.TestMode)
	log := &logger.Logger{Logger: *zaptest.NewLogger(t, zaptest.Level(zap.PanicLevel))}
	return &Service{logger: log}
}

func TestBindRequest(t *testing.T) {
	s := bindTestService(t)

	type testRequest struct {
		SchoolName string `uri:"schoolName" validate:"required"`
		Name       string `json:"name"`
		Page       string `form:"page"`
	}

	tts := []struct {
		name        string
		contentType string
		body        string
		uri         string
		wantErr     bool
	}{
		{
			name:        "json body with uri param",
			contentType: "application/json",
			body:        `{"name":"test"}`,
			uri:         "/test/school1",
			wantErr:     false,
		},
		{
			name:        "query params only",
			contentType: "",
			body:        "",
			uri:         "/test/school1?page=5",
			wantErr:     false,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			var bindErr error
			router.POST("/test/:schoolName", func(c *gin.Context) {
				req := &testRequest{}
				bindErr = s.bindRequest(c, req)
				c.JSON(200, gin.H{"ok": true})
			})

			w := httptest.NewRecorder()
			httpReq := httptest.NewRequest(http.MethodPost, tt.uri, strings.NewReader(tt.body))
			if tt.contentType != "" {
				httpReq.Header.Set("Content-Type", tt.contentType)
			}
			router.ServeHTTP(w, httpReq)

			if tt.wantErr {
				assert.Error(t, bindErr)
			} else {
				assert.NoError(t, bindErr)
			}
		})
	}
}

func TestBindRequestQuery_MapStringString(t *testing.T) {
	s := bindTestService(t)

	type req struct {
		Filters map[string]string `form:"filters"`
	}

	tts := []struct {
		name  string
		query string
		want  map[string]string
	}{
		{
			name:  "with values",
			query: "?filters[name]=test&filters[age]=25",
			want:  map[string]string{"name": "test", "age": "25"},
		},
		{
			name:  "empty",
			query: "",
			want:  nil,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			var result req
			router.GET("/q", func(c *gin.Context) {
				result = req{}
				s.bindRequestQuery(c, &result)
				c.Status(200)
			})

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/q"+tt.query, nil)
			router.ServeHTTP(w, r)

			if tt.want == nil {
				assert.Empty(t, result.Filters)
			} else {
				assert.Equal(t, tt.want, result.Filters)
			}
		})
	}
}

func TestBindRequestQuery_PtrMapStringString(t *testing.T) {
	s := bindTestService(t)

	type req struct {
		Filters *map[string]string `form:"filters"`
	}

	tts := []struct {
		name  string
		query string
		want  *map[string]string
	}{
		{
			name:  "with values",
			query: "?filters[name]=test",
			want:  &map[string]string{"name": "test"},
		},
		{
			name:  "empty",
			query: "",
			want:  nil,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			var result req
			router.GET("/q", func(c *gin.Context) {
				result = req{}
				s.bindRequestQuery(c, &result)
				c.Status(200)
			})

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/q"+tt.query, nil)
			router.ServeHTTP(w, r)

			if tt.want == nil {
				assert.Nil(t, result.Filters)
			} else {
				assert.Equal(t, tt.want, result.Filters)
			}
		})
	}
}

func TestBindRequestQuery_MapStringSliceString(t *testing.T) {
	s := bindTestService(t)

	type req struct {
		Tags map[string][]string `form:"tags"`
	}

	tts := []struct {
		name  string
		query string
		want  map[string][]string
	}{
		{
			name:  "with values",
			query: "?tags[color]=red&tags[color]=blue",
			want:  map[string][]string{"color": {"red", "blue"}},
		},
		{
			name:  "empty",
			query: "",
			want:  nil,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			var result req
			router.GET("/q", func(c *gin.Context) {
				result = req{}
				s.bindRequestQuery(c, &result)
				c.Status(200)
			})

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/q"+tt.query, nil)
			router.ServeHTTP(w, r)

			if tt.want == nil {
				assert.Empty(t, result.Tags)
			} else {
				assert.Equal(t, tt.want, result.Tags)
			}
		})
	}
}

func TestBindRequestQuery_PtrMapStringSliceString(t *testing.T) {
	s := bindTestService(t)

	type req struct {
		Tags *map[string][]string `form:"tags"`
	}

	tts := []struct {
		name  string
		query string
		want  *map[string][]string
	}{
		{
			name:  "with values",
			query: "?tags[color]=red&tags[color]=blue",
			want:  &map[string][]string{"color": {"red", "blue"}},
		},
		{
			name:  "empty",
			query: "",
			want:  nil,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			var result req
			router.GET("/q", func(c *gin.Context) {
				result = req{}
				s.bindRequestQuery(c, &result)
				c.Status(200)
			})

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/q"+tt.query, nil)
			router.ServeHTTP(w, r)

			if tt.want == nil {
				assert.Nil(t, result.Tags)
			} else {
				assert.Equal(t, tt.want, result.Tags)
			}
		})
	}
}
