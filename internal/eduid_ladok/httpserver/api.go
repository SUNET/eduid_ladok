package httpserver

import "eduid_ladok/internal/eduid_ladok/publicapi"

// InternalAPI interface
type InternalAPI interface {
}

// PublicAPI interface
type PublicAPI interface {
	Public(indata *publicapi.RequestPublic) (*publicapi.ReplyPublic, error)
}