package healthcrm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/savannahghi/authutils"
	"github.com/sirupsen/logrus"
)

var (
	accessTokenTimeout = 59 * time.Minute
)

// AuthUtilsLib holds the method defined in authutils library
type authUtilsLib interface {
	Authenticate() (*authutils.OAUTHResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*authutils.OAUTHResponse, error)
}

// client is the library's client used to make requests
type client struct {
	authClient        authUtilsLib
	httpClient        *http.Client
	refreshToken      string
	accessToken       string
	accessTokenTicker *time.Ticker
	authFailed        bool
	baseURL           string
}

// Config contains the settings required to initialize a health crm client
type Config struct {
	AuthServerEndpoint string
	ClientID           string
	ClientSecret       string
	GrantType          string
	Username           string
	Password           string
	BaseURL            string
}

// newClient is the constructor which initializes health crm's authentication mechanism
func newClient(cfg Config) (*client, error) {
	authCfg := authutils.Config{
		AuthServerEndpoint: cfg.AuthServerEndpoint,
		ClientID:           cfg.ClientID,
		ClientSecret:       cfg.ClientSecret,
		GrantType:          cfg.GrantType,
		Username:           cfg.Username,
		Password:           cfg.Password,
	}
	slade360AuthClient, err := authutils.NewClient(authCfg)
	if err != nil {
		return nil, err
	}

	c := client{
		authClient: slade360AuthClient,
		httpClient: &http.Client{
			Timeout: time.Minute * 1,
		},
		accessToken:  "",
		refreshToken: "",
		authFailed:   false,
		baseURL:      cfg.BaseURL,
	}

	err = c.login()
	if err != nil {
		return nil, err
	}

	// set up background routine to update tokens
	go c.background()

	return &c, nil
}

// executed as a go routine to update access and refresh token
func (c *client) background() {
	for t := range c.accessTokenTicker.C {
		logrus.Println("HealthCRM Access Token updated at: ", t)

		err := c.refreshAccessToken()
		if err != nil {
			c.authFailed = true
		} else {
			c.authFailed = false
		}
	}
}

// setAccessToken sets the access token and updates the ticker timer
func (c *client) setRefreshAndAccessToken(token *authutils.OAUTHResponse) {
	c.accessToken = token.AccessToken
	c.refreshToken = token.RefreshToken

	if c.accessTokenTicker != nil {
		c.accessTokenTicker.Reset(accessTokenTimeout)
	} else {
		c.accessTokenTicker = time.NewTicker(accessTokenTimeout)
	}
}

// login uses the provided credentials to login to the authserver backend
// It obtains the necessary tokens required to make authenticated requests
func (c *client) login() error {
	token, err := c.authClient.Authenticate()
	if err != nil {
		return err
	}

	c.setRefreshAndAccessToken(token)

	return nil
}

// refreshAccessToken makes a request to get
// new access and refresh tokens
func (c *client) refreshAccessToken() error {
	ctx := context.Background()

	token, err := c.authClient.RefreshToken(ctx, c.refreshToken)
	if err != nil {
		return err
	}

	c.setRefreshAndAccessToken(token)

	return nil
}

// MakeRequest performs a HTTP request to the provided path and parameters
func (c *client) MakeRequest(ctx context.Context, method, path string, queryParams url.Values, body interface{}) (*http.Response, error) {
	urlPath := fmt.Sprintf("%s%s", c.baseURL, path)

	var request *http.Request
	switch method {
	case http.MethodGet:
		req, err := http.NewRequestWithContext(ctx, method, urlPath, nil)
		if err != nil {
			return nil, err
		}
		request = req

	case http.MethodPost, http.MethodPatch:
		encoded, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		payload := bytes.NewBuffer(encoded)

		req, err := http.NewRequestWithContext(ctx, method, urlPath, payload)
		if err != nil {
			return nil, err
		}

		request = req

	default:
		return nil, fmt.Errorf("s.MakeRequest() unsupported http method: %s", method)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	if queryParams != nil {
		request.URL.RawQuery = queryParams.Encode()
	}

	return c.httpClient.Do(request)
}
