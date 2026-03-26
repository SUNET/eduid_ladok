package ladok

import (
	"context"
	"crypto/x509"
	"eduid_ladok/pkg/model"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.step.sm/crypto/x509util"
)

// importCertificate read crt, key and pem file and put it the certificate handler.
func (s *CertificateService) importCertificate(ctx context.Context) error {
	var (
		certPath = filepath.Join(s.Service.config.Ladok.Certificate.Folder, fmt.Sprintf("%s.%s", s.Service.schoolName, "crt"))
		keyPath  = filepath.Join(s.Service.config.Ladok.Certificate.Folder, fmt.Sprintf("%s.%s", s.Service.schoolName, "key"))
	)

	for _, file := range []string{certPath, keyPath} {
		if _, err := os.Stat(file); err != nil {
			return err
		}
	}

	certFile, err := os.ReadFile(certPath)
	if err != nil {
		return err
	}

	s.CertPEM = certFile

	pemBlock, _ := pem.Decode(certFile)

	s.Cert, err = x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return err
	}

	keyFile, err := os.ReadFile(keyPath)
	if err != nil {
		return err
	}
	s.PrivateKeyPEM = keyFile

	if s.isCertificateInvalid(ctx) {
		return model.ErrCertificateNotValid
	}

	return nil
}

var (
	// Cert90DaysWarning 90 days left until not_after
	Cert90DaysWarning = "CERT_90DAYS_WARNING"
	// CertExpired is expired
	CertExpired = "CERT_EXPIRED"
	// CertOK is ok
	CertOK = "CERT_OK"
)

// CheckValidTime keeps track of time left on certificate, return status
func (s *CertificateService) CheckValidTime(ctx context.Context) (string, time.Time) {
	days90 := time.Now().AddDate(0, 0, 90)
	if s.Cert.NotAfter.Before(days90) {
		return Cert90DaysWarning, s.Cert.NotAfter
	}
	if s.Cert.NotAfter.Before(time.Now()) {
		return CertExpired, s.Cert.NotAfter
	}

	return CertOK, s.Cert.NotAfter
}

func certDaysLeft(notAfter time.Time) int {
	return int(notAfter.Sub(time.Now()).Hours() / 24)
}

// NewSHA256Fingerprint return fingerprint from a *x509.Certificate certificate
func (s *CertificateService) NewSHA256Fingerprint(ctx context.Context) string {
	return x509util.Fingerprint(s.Cert)
}

func (s *CertificateService) isCertificateInvalid(ctx context.Context) bool {
	NAfterA := time.Now().After(s.Cert.NotAfter)
	NBeforeB := time.Now().Before(s.Cert.NotBefore)
	BAfterA := s.Cert.NotBefore.After(s.Cert.NotAfter)
	ABeforeB := s.Cert.NotAfter.Before(s.Cert.NotBefore)

	if NAfterA || NBeforeB || BAfterA || ABeforeB {
		return true
	}

	return false
}
