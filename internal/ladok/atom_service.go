package ladok

import (
	"context"
	"time"

	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"

	"github.com/go-redis/redis/v8"
	"github.com/SUNET/goladok3"
)

// AtomService holds the service object
type AtomService struct {
	Service     *Service
	logger      *logger.Logger
	Channel     chan *model.LadokToAggregateMSG
	ladok       *goladok3.Client
	db          *redis.Client
	quitRunChan chan bool
}

// NewAtomService create a new instance of ladok rest
func NewAtomService(ctx context.Context, service *Service, channel chan *model.LadokToAggregateMSG, logger *logger.Logger) (*AtomService, error) {
	s := &AtomService{
		Channel:     channel,
		Service:     service,
		logger:      logger,
		quitRunChan: make(chan bool),
	}

	switch s.Service.config.Redis.Addr != "" {
	case true:
		s.db = redis.NewClient(&redis.Options{
			Addr: s.Service.config.Redis.Addr,
			DB:   s.Service.config.Redis.DB,
		})
	default:
		s.db = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    s.Service.config.Redis.SentinelServiceName,
			SentinelAddrs: s.Service.config.Redis.SentinelHosts,
			DB:            service.config.Redis.DB,
		})
	}

	// Non-intrusive check if redis is reachable, this will not stop the program even if non-contactable.
	if status := s.StatusRedis(ctx); !status.Healthy {
		s.logger.Warn("Cant connect to redis")
	}

	var err error
	s.ladok, err = goladok3.NewX509(goladok3.X509Config{
		URL:            s.Service.config.Ladok.URL,
		Certificate:    s.Service.Certificate.Cert,
		CertificatePEM: s.Service.Certificate.CertPEM,
		PrivateKeyPEM:  s.Service.Certificate.PrivateKeyPEM,
	})
	if err != nil {
		return nil, err
	}

	ticker := time.NewTicker(time.Duration(s.Service.config.Ladok.Atom.Periodicity) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.run(ctx)
			case <-s.quitRunChan:
				s.logger.Info("Consumer stopped")
				ticker.Stop()
				return
			}
		}
	}()

	s.logger.Info("Started")
	return s, nil
}

// StatusRedis return the status of redis
func (s *AtomService) StatusRedis(ctx context.Context) *model.Status {
	ping := s.db.Ping(ctx).String()
	status := &model.Status{
		Name:       "redis",
		SchoolName: s.Service.schoolName,
		Healthy:    false,
		Status:     model.StatusFail,
		Timestamp:  time.Now(),
	}

	switch ping {
	case "ping: PONG":
		status.Status = model.StatusOK
		status.Healthy = true

		return status
	default:
		return status
	}
}

// Close closes ladok atom service
func (s *AtomService) Close(ctx context.Context) error {
	s.quitRunChan <- true
	ctx.Done()
	s.logger.Info("Quit")
	return nil
}
