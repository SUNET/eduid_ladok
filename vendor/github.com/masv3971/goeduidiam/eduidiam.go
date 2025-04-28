package goeduidiam

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/masv3971/goeduidiam/internal/sunetjwt"
)

// Client holds the object
type Client struct {
	httpClient *http.Client
	url        string
	//jwtLock    *sync.RWMutex
	//jwtObject  *sunetjwt.JWT

	SunetJWT *sunetjwt.Client
	Events   *EventsService
	Groups   *GroupsService
	Invites  *InvitesService
	Status   *StatusService
	Users    *UsersService
}

// Config holds the configuration for New
type Config struct {
	URL   string
	Token TokenConfig
}

// TokenConfig configs token renew
type TokenConfig struct {
	Certificate []byte
	PrivateKey  []byte
	Password    string
	Scope       string
	Type        string
	URL         string
	Key         string
	Client      string
}

// New creates a new instance of goeduidiam
func New(config Config) *Client {
	c := &Client{
		url: config.URL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	c.SunetJWT = sunetjwt.New(sunetjwt.Config{
		Certificate: config.Token.Certificate,
		PrivateKey:  config.Token.PrivateKey,
		Password:    config.Token.Password,
		Scope:       config.Token.Scope,
		Type:        config.Token.Type,
		URL:         config.Token.URL,
		Key:         config.Token.Key,
		Client:      config.Token.Client,
	})
	c.Events = &EventsService{client: c, path: "events"}
	c.Groups = &GroupsService{client: c, path: "groups"}
	c.Invites = &InvitesService{client: c, path: "invites"}
	c.Status = &StatusService{client: c, path: "status"}
	c.Users = &UsersService{client: c, path: "users"}

	return c
}

func (c *Client) newRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
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

	if err := c.SunetJWT.EnsureJWT(ctx); err != nil {
		return nil, err
	}
	// Obtain lock for jwt token
	c.SunetJWT.JWT.RLock()
	defer c.SunetJWT.JWT.RUnlock()

	req, err := http.NewRequestWithContext(ctx, method, url.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "goeduidiam") // TODO(masv): add version
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", c.SunetJWT.JWT.RAW))
	return req, nil
}

func (c *Client) do(req *http.Request, value interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := checkResponse(resp); err != nil {
		errorReply := &Errors{}
		buf := &bytes.Buffer{}
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(buf.Bytes(), errorReply); err != nil {
			return nil, err
		}
		return nil, errorReply
	}

	if err := json.NewDecoder(resp.Body).Decode(value); err != nil {
		return nil, err
	}

	return resp, nil
}

func checkResponse(r *http.Response) error {
	serviceName := "goeduidiam"

	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	case 500:
		return fmt.Errorf("%s: not allowed", serviceName)
	default:
		return fmt.Errorf("%s: invalid request", serviceName)
	}
}

func (c *Client) call(ctx context.Context, verb, path, param string, req, value interface{}) (*http.Response, error) {
	request, err := c.newRequest(
		ctx,
		verb,
		fmt.Sprintf("/%s/%s", path, param),
		req,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(request, value)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
