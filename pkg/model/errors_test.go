package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		msg  string
	}{
		{"ErrNotFound", ErrNotFound, "NOT_FOUND"},
		{"ErrPrivateKeyNotRSA", ErrPrivateKeyNotRSA, "ERR_PRIVATE_KEY_NOT_RSA"},
		{"ErrPrivateKeyEmpty", ErrPrivateKeyEmpty, "ERR_PRIVATE_KEY_EMPTY"},
		{"ErrCRTEmpty", ErrCRTEmpty, "ERR_CRT_EMPTY"},
		{"ErrCRTNotCertificate", ErrCRTNotCertificate, "ERR_CRT_NOT_CERTIFICATE"},
		{"ErrCertificateNotValid", ErrCertificateNotValid, "ERR_CERTIFICATE_NOT_VALID"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualError(t, tt.err, tt.msg)
			assert.True(t, errors.Is(tt.err, tt.err))
		})
	}
}
