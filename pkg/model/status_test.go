package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestManyStatus_Check_AllHealthy(t *testing.T) {
	statuses := ManyStatus{
		{Name: "redis", Healthy: true, Status: StatusOK, Timestamp: time.Now()},
		{Name: "ladok", Healthy: true, Status: StatusOK, Timestamp: time.Now()},
	}

	result := statuses.Check()
	assert.True(t, result.Healthy)
	assert.Equal(t, StatusOK, result.Status)
}

func TestManyStatus_Check_OneUnhealthy(t *testing.T) {
	unhealthy := &Status{Name: "redis", Healthy: false, Status: StatusFail, Message: "connection refused"}
	statuses := ManyStatus{
		unhealthy,
		{Name: "ladok", Healthy: true, Status: StatusOK},
	}

	result := statuses.Check()
	assert.False(t, result.Healthy)
	assert.Equal(t, "redis", result.Name)
}

func TestManyStatus_Check_Empty(t *testing.T) {
	statuses := ManyStatus{}

	result := statuses.Check()
	assert.True(t, result.Healthy)
	assert.Equal(t, StatusOK, result.Status)
}

func TestMonitoringCertClient(t *testing.T) {
	clients := MonitoringCertClients{
		"school1": {
			Valid:       true,
			Fingerprint: "abc123",
			NotAfter:    time.Now().AddDate(0, 0, 100),
			DaysLeft:    100,
			LastChecked: time.Now(),
		},
	}

	assert.Len(t, clients, 1)
	assert.True(t, clients["school1"].Valid)
	assert.Equal(t, "abc123", clients["school1"].Fingerprint)
}
