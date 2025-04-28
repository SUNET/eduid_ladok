package goladok3

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/masv3971/goladok3/ladoktypes"
	"golang.org/x/time/rate"
)

var (
	ErrInvalidRequest    = errors.New("Invalid request")
	ErrNotAllowedRequest = errors.New("Not allowed request")
)

// X509Config configures new function
type X509Config struct {
	URL            string            `validate:"required"`
	Certificate    *x509.Certificate `validate:"required"`
	CertificatePEM []byte            `validate:"required"`
	//PrivateKey     *rsa.PrivateKey   `validate:"required"`
	PrivateKeyPEM []byte `validate:"required"`
	ProxyURL      string
}

// OidcConfig configures NewOIDC function
type OidcConfig struct {
}

// Client holds the ladok object
type Client struct {
	HTTPClient     *http.Client
	rateLimit      *rate.Limiter
	format         string
	url            string
	certificate    *x509.Certificate
	certificatePEM []byte
	chain          *x509.CertPool
	chainPEM       []byte
	privateKey     *rsa.PrivateKey
	privateKeyPEM  []byte
	proxyURL       string

	Kataloginformation *kataloginformationService
	Studentinformation *studentinformationService
	Studentdeltagande  *studentdeltagandeService
	Feed               *feedService
}

// NewX509 create a new x509 instance of ladok
func NewX509(config X509Config) (*Client, error) {
	if err := Check(config); err != nil {
		return nil, err
	}
	c := &Client{
		format:         "json",
		url:            config.URL,
		proxyURL:       config.ProxyURL,
		privateKeyPEM:  config.PrivateKeyPEM,
		certificatePEM: config.CertificatePEM,
		certificate:    config.Certificate,
		//privateKey:     config.PrivateKey,
		rateLimit: rate.NewLimiter(rate.Every(1*time.Second), 30),
	}

	if err := c.httpConfigure(); err != nil {
		return nil, err
	}

	c.Studentinformation = &studentinformationService{client: c, service: "studentinformation"}
	c.Kataloginformation = &kataloginformationService{client: c, service: "kataloginformation"}
	c.Studentdeltagande = &studentdeltagandeService{client: c, service: "studentdeltagande"}
	c.Feed = &feedService{client: c, service: "feed"}

	return c, nil
}

// NewOIDC create a new OIDC instance of ladok
func NewOIDC(config OidcConfig) (*Client, error) {
	return nil, nil
}

func (c *Client) httpConfigure() error {
	keyPair, err := tls.X509KeyPair(c.certificatePEM, c.privateKeyPEM)
	if err != nil {
		return err
	}

	tlsCfg := &tls.Config{
		Rand:         rand.Reader,
		Certificates: []tls.Certificate{keyPair},
		NextProtos:   []string{},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		//ClientCAs:          c.chainDER,
		InsecureSkipVerify: false,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		},
		PreferServerCipherSuites:    true,
		SessionTicketsDisabled:      false,
		DynamicRecordSizingDisabled: false,
		Renegotiation:               0,
		KeyLogWriter:                nil,
	}

	//	tlsCfg.BuildNameToCertificate()

	c.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:     tlsCfg,
			DialContext:         nil,
			TLSHandshakeTimeout: 30 * time.Second,
			Proxy:               http.ProxyFromEnvironment,
		},
	}

	return nil
}

var ladokAcceptHeader = map[string]map[string]string{
	"studentinformation": {
		"json": "application/vnd.ladok-studentinformation+json",
		"xml":  "application/vnd.ladok-studentinformation+xml",
	},
	"studentdeltagande": {
		"json": "application/vnd.ladok-studiedeltagande+json",
		"xml":  "application/vnd.ladok-studiedeltagande+xml",
	},
	"resultat": {
		"json": "application/vnd.ladok-resultat+json",
		"xml":  "application/vnd.ladok-resultat+xml",
	},
	"uppfoljning": {
		"json": "application/vnd.ladok-uppfoljning+json",
		"xml":  "application/vnd.ladok-uppfoljning+xml",
	},
	"examen": {
		"json": "application/vnd.ladok-examen+json",
		"xml":  "application/vnd.ladok-examen+xml",
	},
	"utbildningsinformation": {
		"json": "application/vnd.ladok-utbildningsinformation+json",
		"xml":  "application/vnd.ladok-utbildningsinformation+xml",
	},
	"kataloginformation": {
		"json": "application/vnd.ladok-kataloginformation+json",
		"xml":  "application/vnd.ladok-kataloginformation+xml",
	},
	"feed": {
		"xml": "application/atom+xml",
	},
}

// NewRequest make a new request
func (c *Client) newRequest(ctx context.Context, acceptHeader string, method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(c.url)
	if err != nil {
		return nil, err
	}
	url := u.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		payload := struct {
			Data interface{} `json:"data"`
		}{
			Data: body,
		}
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", acceptHeader)
	req.Header.Set("User-Agent", "goladok3/0.0.15")

	return req, nil
}

// Do does the new request
func (c *Client) do(ctx context.Context, req *http.Request, value interface{}) (*http.Response, error) {
	if err := c.rateLimit.Wait(ctx); err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := checkResponse(resp); err != nil {
		buf := &bytes.Buffer{}
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			return nil, err
		}
		ladokError := &ladoktypes.LadokError{}
		if err := json.Unmarshal(buf.Bytes(), ladokError); err != nil { // TODO(masv): Fix xml error parsing into Errors.
			return nil, err
		}
		return nil, ladokError
	}

	switch resp.Header.Get("Content-Type") {
	case ContentTypeAtomXML:
		if err := xml.NewDecoder(resp.Body).Decode(value); err != nil {
			return nil, err
		}
	case ContentTypeKataloginformationJSON, ContentTypeStudiedeltagandeJSON, ContentTypeStudentinformationJSON:
		if err := json.NewDecoder(resp.Body).Decode(value); err != nil {
			return nil, err
		}
	default:
		return nil, ladoktypes.ErrNoValidContentType
	}

	return resp, nil
}

func checkResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	case 500:
		return ErrInvalidRequest
	case 401:
		return ErrNotAllowedRequest
	}

	return ErrInvalidRequest
}

func (c *Client) call(ctx context.Context, acceptHeader, method, url string, body, reply interface{}) (*http.Response, error) {
	request, err := c.newRequest(
		ctx,
		acceptHeader,
		method,
		url,
		body,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(ctx, request, reply)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
