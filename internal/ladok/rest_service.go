package ladok

import (
	"context"
	"fmt"
	"time"

	"eduid_ladok/pkg/helpers"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"

	"github.com/SUNET/goladok3"
)

// RestService holds the rest service
type RestService struct {
	Service *Service
	logger  *logger.Logger
	Ladok   *goladok3.Client
}

// NewRestService creates a new instance of rest
func NewRestService(ctx context.Context, service *Service, logger *logger.Logger) (*RestService, error) {
	s := &RestService{
		logger:  logger,
		Service: service,
	}

	var err error
	s.Ladok, err = goladok3.NewX509(goladok3.X509Config{
		URL:            s.Service.config.Ladok.URL,
		ProxyURL:       s.Service.config.HTTPProxy,
		Certificate:    s.Service.Certificate.Cert,
		CertificatePEM: s.Service.Certificate.CertPEM,
		PrivateKeyPEM:  s.Service.Certificate.PrivateKeyPEM,
	})
	if err != nil {
		return nil, err
	}

	s.logger.Info("Started")
	return s, nil
}

// StatusLadok return the status of ladok
func (s *RestService) StatusLadok(ctx context.Context) *model.Status {
	status := &model.Status{
		Name:       "Ladok rest",
		SchoolName: s.Service.schoolName,
		Healthy:    false,
		Status:     model.StatusFail,
		Timestamp:  time.Now(),
	}

	data, resp, err := s.Ladok.Kataloginformation.GetGrunddataLarosatesinformation(ctx)
	if err != nil {
		status.Message = fmt.Sprintf("%s %s", err.Error(), helpers.FormatResponse(resp))
		return status
	}

	if data == nil {
		status.Message = "Empty return, no data"
		return status
	}

	status.Healthy = true
	status.Status = model.StatusOK

	return status
}

// Close closes service ladok rest
func (s *RestService) Close(ctx context.Context) error {
	s.logger.Info("Quit")
	ctx.Done()
	return nil
}
