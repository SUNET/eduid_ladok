package aggregate

import (
	"context"
	"eduid_ladok/internal/ladok"
	"eduid_ladok/pkg/logger"
	"sync"
)

// Config holds the configuration for aggregate
type Config struct {
}

// Service holds the service object for aggregate
type Service struct {
	config   Config
	logger   *logger.Logger
	wg       *sync.WaitGroup
	ladok    *ladok.Service
	feedName string
}

// New creates a new instance of aggregate
func New(ctx context.Context, config Config, wg *sync.WaitGroup, feedName string, ladok *ladok.Service, logger *logger.Logger) (*Service, error) {
	s := &Service{
		logger:   logger,
		config:   config,
		ladok:    ladok,
		wg:       wg,
		feedName: feedName,
	}

	s.wg.Add(1)
	go s.run(ctx)

	return s, nil
}

// Close closes aggregate service
func (s *Service) Close(ctx context.Context) error {
	defer s.wg.Done()

	s.logger.Warn("Quit")
	return nil
}