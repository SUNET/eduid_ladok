package ladok

import (
	"bytes"
	"context"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"

	"github.com/SUNET/goladok3"
	"github.com/SUNET/goladok3/ladokmocks"
	"github.com/SUNET/goladok3/ladoktypes"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func mockGenericEndpointServer(t *testing.T, mux *http.ServeMux, contentType, method, url, param string, payload []byte, statusCode int) {
	mux.HandleFunc(fmt.Sprintf("%s/%s", url, param),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentType)
			w.WriteHeader(statusCode)
			testMethod(t, r, method)
			testURL(t, r, fmt.Sprintf("%s/%s", url, param))
			w.Write(payload)
		},
	)
}

func mockLadokHTTPServer(t *testing.T, statusCode int) *httptest.Server {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)

	endpoints := []struct {
		url                  string
		param                string
		paramType            string
		numericFeedEndpoint  bool
		personnummerEndpoint bool
		statusCode           int
		contentType          string
		method               string
		serverReturnPayload  []byte
	}{
		{
			url:                 "/uppfoljning/feed/recent",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeAtomXML,
			method:              "GET",
			serverReturnPayload: ladokmocks.FeedXML(100),
		},
		{
			url:                 "/handelser/feed/recent",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeAtomXML,
			method:              "GET",
			serverReturnPayload: ladokmocks.FeedXML(100),
		},
		{
			url:                 "/uppfoljning/feed",
			paramType:           "historicalAtom",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeAtomXML,
			method:              "GET",
			serverReturnPayload: ladokmocks.XMLFeedRecent,
		},
		{
			url:                 "/handelser/feed",
			paramType:           "historicalAtom",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeAtomXML,
			method:              "GET",
			serverReturnPayload: ladokmocks.XMLFeedRecent,
		},
		{
			url:                 "/uppfoljning/feed/first",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeAtomXML,
			method:              "GET",
			serverReturnPayload: ladokmocks.FeedXML(1),
		},
		{
			url:                 "/handelser/feed/first",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeAtomXML,
			method:              "GET",
			serverReturnPayload: ladokmocks.FeedXML(1),
		},
		{
			url:                 "/kataloginformation/anvandare/autentiserad",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeKataloginformationJSON,
			method:              "GET",
			serverReturnPayload: ladokmocks.JSONKataloginformationAutentiserad,
		},
		{
			url:                 "/kataloginformation/anvandarbehorighet/egna",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeKataloginformationJSON,
			method:              "GET",
			serverReturnPayload: ladokmocks.JSONKataloginformationEgna,
		},
		{
			url:                 "/kataloginformation/behorighetsprofil",
			param:               ladokmocks.BehorighetsprofilUID,
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeKataloginformationJSON,
			method:              "GET",
			serverReturnPayload: ladokmocks.JSONKataloginformationBehorighetsprofil,
		},
		{
			url:                 "/studentinformation/student",
			param:               ladokmocks.Students[0].StudentUID,
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeKataloginformationJSON,
			method:              "GET",
			serverReturnPayload: ladokmocks.JSONStudentinformationStudent,
		},
		{
			url:                 "/studentinformation/student/personnummer",
			paramType:           "personnummer",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeKataloginformationJSON,
			method:              "GET",
			serverReturnPayload: ladokmocks.JSONStudentinformationStudent,
		},
	}

	for _, endpoint := range endpoints {
		switch endpoint.paramType {
		case "personnummer":
			for _, student := range ladokmocks.Students {
				mockGenericEndpointServer(t, mux, endpoint.contentType, endpoint.method, endpoint.url, student.Personnummer, ladokmocks.StudentJSON(student), endpoint.statusCode)
			}
		case "historicalAtom":
			for id := 1; id <= 100; id++ {
				mockGenericEndpointServer(t, mux, endpoint.contentType, endpoint.method, endpoint.url, strconv.Itoa(id), ladokmocks.FeedXML(id), endpoint.statusCode)
			}
		default:
			mockGenericEndpointServer(t, mux, endpoint.contentType, endpoint.method, endpoint.url, endpoint.param, endpoint.serverReturnPayload, endpoint.statusCode)
		}
	}

	// Register endpoint without trailing slash for grunddata/larosatesinformation
	mux.HandleFunc("/kataloginformation/grunddata/larosatesinformation", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", goladok3.ContentTypeKataloginformationJSON)
		w.WriteHeader(statusCode)
		w.Write(ladokmocks.JSONKataloginformationGrunddataLarosateinformation)
	})

	return server
}

func TestMockEndpoints(t *testing.T) {
	service, server, _, _ := mockService(t, 200, 0, 100, t.TempDir())
	defer server.Close()

	tts := []struct {
		name      string
		want      any
		fn        any
		attribute any
	}{
		{
			name:      "/handelser/feed/recent",
			want:      ladokmocks.MockSuperFeed(100),
			attribute: 100,
			fn:        service.Atom.ladok.Feed.Recent,
		},
		{
			name:      "/handelser/feed/50",
			want:      ladokmocks.MockSuperFeed(50),
			attribute: 50,
			fn:        service.Atom.ladok.Feed.Historical,
		},
		{
			name:      "/handelser/feed/first",
			want:      ladokmocks.MockSuperFeed(1),
			attribute: 1,
			fn:        service.Atom.ladok.Feed.First,
		},
		{
			name: "/kataloginformation/anvandare/autentiserad",
			want: ladokmocks.MockKataloginformationAutentiserad(),
			fn:   service.Rest.Ladok.Kataloginformation.GetAnvandareAutentiserad,
		},
		{
			name: "/kataloginformation/anvandare/egna",
			want: ladokmocks.MockKataloginformationEgna(),
			fn:   service.Rest.Ladok.Kataloginformation.GetAnvandarbehorighetEgna,
		},
		{
			name:      "/kataloginformation/behorighetsprofil",
			want:      ladokmocks.MockKataloginformationBehorighetsprofil(),
			fn:        service.Rest.Ladok.Kataloginformation.GetBehorighetsprofil,
			attribute: ladokmocks.BehorighetsprofilUID,
		},
		{
			name:      "/studentinformation/student UID",
			want:      ladokmocks.MockStudentinformationStudent(),
			fn:        service.Rest.Ladok.Studentinformation.GetStudent,
			attribute: ladokmocks.Students[0].StudentUID,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.fn.(type) {
			case func(context.Context) (*ladoktypes.SuperFeed, *http.Response, error):
				f := tt.fn.(func(context.Context) (*ladoktypes.SuperFeed, *http.Response, error))
				reply, _, err := f(t.Context())
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				assert.Equal(t, tt.want, reply)
			case func(context.Context, int) (*ladoktypes.SuperFeed, *http.Response, error):
				f := tt.fn.(func(context.Context, int) (*ladoktypes.SuperFeed, *http.Response, error))
				reply, _, err := f(t.Context(), tt.attribute.(int))
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				assert.Equal(t, tt.want, reply)
			case func(context.Context, *goladok3.HistoricalReq) (*ladoktypes.SuperFeed, *http.Response, error):
				f := tt.fn.(func(context.Context, *goladok3.HistoricalReq) (*ladoktypes.SuperFeed, *http.Response, error))
				reply, _, err := f(t.Context(), &goladok3.HistoricalReq{ID: tt.attribute.(int)})
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				assert.Equal(t, tt.want, reply)
			case func(context.Context) (*ladoktypes.KataloginformationAnvandareAutentiserad, *http.Response, error):
				f := tt.fn.(func(ctx context.Context) (*ladoktypes.KataloginformationAnvandareAutentiserad, *http.Response, error))
				reply, _, err := f(t.Context())
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				assert.Equal(t, tt.want, reply)
			case func(context.Context) (*ladoktypes.KataloginformationAnvandarbehorighetEgna, *http.Response, error):
				f := tt.fn.(func(context.Context) (*ladoktypes.KataloginformationAnvandarbehorighetEgna, *http.Response, error))
				reply, _, err := f(t.Context())
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				assert.Equal(t, tt.want, reply)
			case func(context.Context, *goladok3.GetBehorighetsprofilerReq) (*ladoktypes.KataloginformationBehorighetsprofil, *http.Response, error):
				f := tt.fn.(func(context.Context, *goladok3.GetBehorighetsprofilerReq) (*ladoktypes.KataloginformationBehorighetsprofil, *http.Response, error))
				reply, _, err := f(t.Context(), &goladok3.GetBehorighetsprofilerReq{UID: tt.attribute.(string)})
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				assert.Equal(t, tt.want, reply)
			case func(context.Context, *goladok3.GetStudentReq) (*ladoktypes.Student, *http.Response, error):
				f := tt.fn.(func(context.Context, *goladok3.GetStudentReq) (*ladoktypes.Student, *http.Response, error))
				reply, _, err := f(t.Context(), &goladok3.GetStudentReq{UID: tt.attribute.(string)})
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				assert.Equal(t, tt.want, reply)
			default:
				t.Errorf("ERROR: can't find a matching reflecting type: %T", tt.fn)
				t.FailNow()
			}
		})
	}
}

func mockService(t *testing.T, statusCode, notBefore, notAfter int, tempDir string) (*Service, *httptest.Server, redismock.ClientMock, error) {
	server := mockLadokHTTPServer(t, statusCode)

	ladokToAggregateChan := make(chan *model.LadokToAggregateMSG, 200)

	cfg := &model.Cfg{}
	cfg.Ladok.Certificate.Folder = tempDir
	cfg.Ladok.URL = server.URL
	cfg.Ladok.Atom.Periodicity = 60

	mockCertificate(t, notBefore, notAfter, tempDir)
	testLog := logger.Logger{
		Logger: *zaptest.NewLogger(t, zaptest.Level(zap.PanicLevel)),
	}

	service, err := New(t.Context(), cfg, &sync.WaitGroup{}, "testSchoolName", ladokToAggregateChan, testLog.New("test"))
	if err != nil {
		return nil, nil, nil, err
	}

	r, redisMock := redismock.NewClientMock()
	service.Atom.db = r

	return service, server, redisMock, nil
}

func testMethod(t *testing.T, r *http.Request, want string) {
	got := r.Method
	assert.Equal(t, want, got)
}

func testURL(t *testing.T, r *http.Request, want string) {
	got := r.RequestURI
	assert.Equal(t, want, got)
}

func testBody(t *testing.T, r *http.Request, want string) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	assert.NoError(t, err)

	got := buffer.String()
	require.JSONEq(t, want, got)
}
