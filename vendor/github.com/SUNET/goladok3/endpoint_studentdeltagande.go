package goladok3

import (
	"context"
	"net/http"

	"github.com/SUNET/goladok3/ladoktypes"
)

// studentdeltagandeService xx
type studentdeltagandeService struct {
	client  *Client
	service string
}

func (s *studentdeltagandeService) acceptHeader() string {
	return ladokAcceptHeader[s.service][s.client.format]
}

// GetTillfallesdeltagandePagaendeStudentReq request
type GetTillfallesdeltagandePagaendeStudentReq struct {
	StudentUID string `validate:"required"`
}

// TODO(masv): fix!
func (s *studentdeltagandeService) GetTillfallesdeltagandePagaendeStudent(ctx context.Context, req GetAktivPaLarosateReq) (*ladoktypes.TillfallesdeltagandePagaendeStudent, *http.Response, error) {
	return nil, nil, nil
}
