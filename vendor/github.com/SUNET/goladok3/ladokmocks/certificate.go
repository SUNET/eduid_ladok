package ladokmocks

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// MockCertificatePassword mock password for certificate
var MockCertificatePassword = "testPassword"

func mockCACertificateAndKey(t *testing.T) (*x509.Certificate, *rsa.PrivateKey, []*x509.Certificate) {
	ca := &x509.Certificate{
		SignatureAlgorithm: x509.SHA256WithRSA,
		PublicKeyAlgorithm: x509.RSA,
		Version:            3,
		SerialNumber:       big.NewInt(2300),
		Issuer: pkix.Name{
			Organization:       []string{"Ladok3", "LED", "ca01", "ROOT", "CA"},
			OrganizationalUnit: []string{"pki-tomcat"},
			CommonName:         "CA Signing Certificate",
		},
		Subject: pkix.Name{
			Country:            []string{"SE"},
			Organization:       []string{"Ladok"},
			OrganizationalUnit: []string{"LED"},
			Locality:           []string{"Stockholm"},
			SerialNumber:       "",
			CommonName:         "sunet@KF",
		},
		NotBefore:  time.Now(),
		NotAfter:   time.Now().AddDate(10, 0, 0),
		IsCA:       true,
		KeyUsage:   x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		OCSPServer: []string{"URI:http://ca01.rome08.led.ladok.se:8080/ca/ocsp"},
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &privateKey.PublicKey, privateKey)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	signRootCA, err := x509.ParseCertificates(caBytes)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	return ca, privateKey, signRootCA
}

// MockCertificateAndKey return mock certificate template
//func MockCertificateAndKey(t *testing.T, env, schoolName, tempdir string, notBefore, notAfter int) (*x509.Certificate, *rsa.PrivateKey, []*x509.Certificate) {
func MockCertificateAndKey(t *testing.T, env string, notBefore, notAfter int) ([]byte, *x509.Certificate, []byte, *rsa.PrivateKey) {
	ca, caPrivateKey, _ := mockCACertificateAndKey(t)

	certTemplate := &x509.Certificate{
		SignatureAlgorithm: x509.SHA256WithRSA,
		PublicKeyAlgorithm: x509.RSA,
		Version:            3,
		SerialNumber:       big.NewInt(2300),
		Issuer: pkix.Name{
			SerialNumber: "",
			CommonName:   "Ladok3 LED MIT API CA",
		},
		Subject: pkix.Name{
			Country:            []string{"SE"},
			Organization:       []string{"Ladok"},
			OrganizationalUnit: []string{"LED", env},
			Locality:           []string{"Stockholm"},
			SerialNumber:       "",
			CommonName:         "sunet@KF",
		},
		NotBefore:   time.Now().AddDate(0, 0, notBefore),
		NotAfter:    time.Now().AddDate(0, 0, notAfter),
		KeyUsage:    x509.KeyUsageDataEncipherment | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		OCSPServer:  []string{"URI:http://ca01.rome08.led.ladok.se:8080/ca/ocsp"},
	}

	clientPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(clientPrivateKey),
	}

	certDERByte, err := x509.CreateCertificate(rand.Reader, certTemplate, ca, &clientPrivateKey.PublicKey, caPrivateKey)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	clientCert, err := x509.ParseCertificate(certDERByte)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	certPEM := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: clientCert.Raw,
	}
	//certPEMFile, err := os.Create(filepath.Join(tempdir, fmt.Sprintf("%s.pem", schoolName)))
	//if err != nil {
	//	t.FailNow()
	//}
	//if err := pem.Encode(certPEMFile, certPEM); err != nil {
	//	t.FailNow()
	//}
	//certPEMFile.Close()

	//privateKeyPEMFile, err := os.Create(filepath.Join(tempdir, fmt.Sprintf("%s.key", schoolName)))
	//if err != nil {
	//	t.FailNow()
	//}
	//if err := pem.Encode(privateKeyPEMFile, privateKeyPEM); err != nil {
	//	t.FailNow()
	//}
	//privateKeyPEMFile.Close()
	assert.NotNil(t, pem.EncodeToMemory(certPEM))
	assert.NotNil(t, clientCert)
	assert.NotNil(t, pem.EncodeToMemory(privateKeyPEM))
	assert.NotNil(t, clientPrivateKey)

	if !assert.NotEmpty(t, clientCert.Subject.OrganizationalUnit[1]) {
		t.FailNow()
	}

	return pem.EncodeToMemory(certPEM), clientCert, pem.EncodeToMemory(privateKeyPEM), clientPrivateKey
	//	chainPEM := &pem.Block{}
	//
	//	c.chain = x509.NewCertPool()
	//	for _, cert := range chainCerts {
	//		c.chain.AddCert(cert)
	//	}
	//return clientCert, certPrivateKey, signRootCA
}
