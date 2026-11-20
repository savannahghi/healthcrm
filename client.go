package healthcrm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const requestTimeout = 15 * time.Second

type Config struct {
	// KeycloakURL is the Keycloak server's base URL, e.g. https://auth.example.com
	KeycloakURL  string
	Realm        string
	ClientID     string
	ClientSecret string // nolint: gosec
	BaseURL      string
}

func (c Config) validate() error {
	fields := []struct{ name, value string }{
		{"KeycloakURL", c.KeycloakURL},
		{"Realm", c.Realm},
		{"ClientID", c.ClientID},
		{"ClientSecret", c.ClientSecret},
		{"BaseURL", c.BaseURL},
	}

	var missing []string

	for _, f := range fields {
		if strings.TrimSpace(f.value) == "" {
			missing = append(missing, f.name)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required config: %s", strings.Join(missing, ", "))
	}

	return nil
}

// tokenURL is the realm's OpenID Connect token endpoint
func (c Config) tokenURL() string {
	return fmt.Sprintf(
		"%s/realms/%s/protocol/openid-connect/token",
		strings.TrimSuffix(c.KeycloakURL, "/"),
		url.PathEscape(c.Realm),
	)
}

// client is the library's client used to make requests
type client struct {
	httpClient *http.Client
	baseURL    string
}

// newClient authenticates with Keycloak's client-credentials grant; tokens are cached and renewed before they expire.
func newClient(cfg Config) (*client, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	baseTransport := otelhttp.NewTransport(http.DefaultTransport)

	// oauth2 uses this client for token requests; the context must never be cancelled.
	tokenCtx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{
		Timeout:   requestTimeout,
		Transport: baseTransport,
	})

	credentials := clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		TokenURL:     cfg.tokenURL(),
		AuthStyle:    oauth2.AuthStyleInParams,
	}

	tokenSource := credentials.TokenSource(tokenCtx)

	// Fail fast on bad credentials or a wrong realm.
	if _, err := tokenSource.Token(); err != nil {
		return nil, fmt.Errorf("failed to fetch Keycloak token: %w", err)
	}

	return &client{
		httpClient: &http.Client{
			Timeout: requestTimeout,
			Transport: &oauth2.Transport{
				Source: tokenSource,
				Base:   baseTransport,
			},
		},
		baseURL: cfg.BaseURL,
	}, nil
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

	if queryParams != nil {
		request.URL.RawQuery = queryParams.Encode()
	}

	// oauth2.Transport adds the Bearer token.
	return c.httpClient.Do(request) // nolint: gosec
}
