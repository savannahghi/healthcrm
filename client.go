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
	"github.com/savannahghi/serverutils"
	"github.com/sirupsen/logrus"
)

var (
	accessTokenTimeout = 59 * time.Minute
)

// IAuthUtilsLib holds the method defined in authutils library
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
}

// newClient is the constructor which initializes health crm's authentication mechanism
func newClient() (*client, error) {
	config := authutils.Config{
		AuthServerEndpoint: serverutils.MustGetEnvVar("HEALTH_CRM_AUTH_SERVER_ENDPOINT"),
		ClientID:           serverutils.MustGetEnvVar("HEALTH_CRM_CLIENT_ID"),
		ClientSecret:       serverutils.MustGetEnvVar("HEALTH_CRM_CLIENT_SECRET"),
		GrantType:          serverutils.MustGetEnvVar("HEALTH_CRM_GRANT_TYPE"),
		Username:           serverutils.MustGetEnvVar("HEALTH_CRM_USERNAME"),
		Password:           serverutils.MustGetEnvVar("HEALTH_CRM_PASSWORD"),
	}
	slade360AuthClient, err := authutils.NewClient(config)
	if err != nil {
		return nil, err
	}

	c := client{
		authClient: slade360AuthClient,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		accessToken:  "",
		refreshToken: "",
		authFailed:   false,
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
func (cl *client) MakeRequest(ctx context.Context, method, path string, queryParams url.Values, body interface{}) (*http.Response, error) {
	oauthResponse, err := cl.authClient.Authenticate()
	if err != nil {
		return nil, err
	}

	urlPath := fmt.Sprintf("%s%s", BaseURL, path)

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
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", oauthResponse.AccessToken))

	if queryParams != nil {
		request.URL.RawQuery = queryParams.Encode()
	}

	return cl.httpClient.Do(request)
}
