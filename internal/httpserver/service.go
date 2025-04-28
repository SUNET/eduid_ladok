package httpserver

import (
	"context"
	"eduid_ladok/internal/apiv1"
	"eduid_ladok/pkg/helpers"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/masv3971/goladok3/ladoktypes"
)

// Service is the service object for httpserver
type Service struct {
	config *model.Cfg
	logger *logger.Logger
	server *http.Server
	apiv1  Apiv1
	gin    *gin.Engine
}

// New creates a new httpserver service
func New(ctx context.Context, config *model.Cfg, api *apiv1.Client, logger *logger.Logger) (*Service, error) {
	s := &Service{
		config: config,
		logger: logger,
		apiv1:  api,
		server: &http.Server{Addr: config.APIServer.Host},
	}

	switch s.config.Production {
	case true:
		gin.SetMode(gin.ReleaseMode)
	case false:
		gin.SetMode(gin.DebugMode)
	}

	apiValidator := validator.New()
	binding.Validator = &defaultValidator{
		Validate: apiValidator,
	}

	s.gin = gin.New()
	s.server.Handler = s.gin
	s.server.ReadTimeout = time.Second * 5
	s.server.WriteTimeout = time.Second * 30
	s.server.IdleTimeout = time.Second * 90

	// Middlewares
	s.gin.Use(s.middlewareTraceID(ctx))
	s.gin.Use(s.middlewareDuration(ctx))
	s.gin.Use(s.middlewareLogger(ctx))
	s.gin.Use(s.middlewareCrash(ctx))
	s.gin.NoRoute(func(c *gin.Context) { c.JSON(http.StatusNotFound, helpers.Problem404()) })

	s.regEndpoint(ctx, "api/v1/:schoolName/ladokinfo", "POST", s.endpointLadokInfo)
	s.regEndpoint(ctx, "api/v1/schoolinfo", "GET", s.endpointSchoolInfo)

	s.regEndpoint(ctx, "/health", "GET", s.endpointStatus)
	s.regEndpoint(ctx, "/monitoring/cert/client", "GET", s.endpointMonitoringCertClient)

	// Run http server
	go func() {
		err := s.server.ListenAndServe()
		if err != nil {
			s.logger.New("http").Fatal("listen_error", "error", err)
		}
	}()

	s.logger.Info("started")

	return s, nil
}

func (s *Service) regEndpoint(ctx context.Context, path, method string, handler func(context.Context, *gin.Context) (interface{}, error)) {
	s.gin.Handle(method, path, func(c *gin.Context) {
		res, err := handler(ctx, c)

		var (
			status int = 200
		)

		if err != nil {
			switch err.(type) {
			case *ladoktypes.LadokError:
				status = 400
			case ladoktypes.PermissionErrors:
				status = 400
			case validator.ValidationErrors:
				status = 400
			default:
				status = 400
			}
		}

		renderContent(c, status, gin.H{"data": res, "error": helpers.NewErrorFromError(err)})
	})
}

func renderContent(c *gin.Context, code int, data interface{}) {
	switch c.NegotiateFormat(gin.MIMEJSON, "*/*") {
	case gin.MIMEJSON:
		c.JSON(code, data)
	case "*/*": // curl
		c.JSON(code, data)
	default:
		c.JSON(406, gin.H{"data": nil, "error": helpers.NewErrorDetails("not_acceptable", "Accept header is invalid. It should be \"application/json\".")})
	}
}

// Close closing httpserver
func (s *Service) Close(ctx context.Context) error {
	s.logger.Info("Quit")
	return nil
}
