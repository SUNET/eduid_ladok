package goladok3

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SUNET/goladok3/ladoktypes"
)

// kataloginformationService handles kataloginformation
type kataloginformationService struct {
	client  *Client
	service string
}

func (s *kataloginformationService) acceptHeader() string {
	return ladokAcceptHeader[s.service][s.client.format]
}

// GetAnvandareAutentiserad gets kataloginformation/anvandare/autentiserad
func (s *kataloginformationService) GetAnvandareAutentiserad(ctx context.Context) (*ladoktypes.KataloginformationAnvandareAutentiserad, *http.Response, error) {
	url := fmt.Sprintf("%s/%s", s.service, "anvandare/autentiserad")
	reply := &ladoktypes.KataloginformationAnvandareAutentiserad{}
	resp, err := s.client.call(ctx, s.acceptHeader(), http.MethodGet, url, nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// GetBehorighetsprofilerReq configuration for GetBehorighetsprofil
type GetBehorighetsprofilerReq struct {
	UID string `validate:"required"`
}

// GetBehorighetsprofil return structure of rights for uid
func (s *kataloginformationService) GetBehorighetsprofil(ctx context.Context, req *GetBehorighetsprofilerReq) (*ladoktypes.KataloginformationBehorighetsprofil, *http.Response, error) {
	if err := Check(req); err != nil {
		return nil, nil, err
	}

	url := fmt.Sprintf("%s/%s/%s", s.service, "behorighetsprofil", req.UID)
	reply := &ladoktypes.KataloginformationBehorighetsprofil{}
	resp, err := s.client.call(ctx, s.acceptHeader(), http.MethodGet, url, nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// GetAnvandarbehorighetEgna return structure of ladok permission
func (s *kataloginformationService) GetAnvandarbehorighetEgna(ctx context.Context) (*ladoktypes.KataloginformationAnvandarbehorighetEgna, *http.Response, error) {
	url := fmt.Sprintf("%s/%s", s.service, "anvandarbehorighet/egna")
	reply := &ladoktypes.KataloginformationAnvandarbehorighetEgna{}
	resp, err := s.client.call(ctx, s.acceptHeader(), http.MethodGet, url, nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// GetAnvandarbehorighetEgna return structure of ladok permission
func (s *kataloginformationService) GetGrunddataLarosatesinformation(ctx context.Context) (*ladoktypes.KataloginformationGrunddataLarosatesinformation, *http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s", s.service, "grunddata", "larosatesinformation")
	reply := &ladoktypes.KataloginformationGrunddataLarosatesinformation{}
	resp, err := s.client.call(ctx, s.acceptHeader(), http.MethodGet, url, nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}
