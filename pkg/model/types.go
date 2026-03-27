package model

import (
	"time"

	"github.com/SUNET/goladok3/ladoktypes"
)

// LadokToAggregateMSG is the message on the channel, ladok-atom --> aggregate
type LadokToAggregateMSG struct {
	Event *ladoktypes.SuperEvent
	TS    time.Time
}

// AggregateToEduIDMSG is message on the channel from Aggregate to EduID
type AggregateToEduIDMSG struct {
}

// UserData consists of identifiers for a user
type UserData struct {
	NIN string `json:"nin" validate:"required"`
}
