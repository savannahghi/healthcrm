package healthcrm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestMakeRequest(t *testing.T) {
	ctx := context.Background()
	queryParam := url.Values{}
	queryParam.Add("param1", "value1")

	tests := []struct {
		name        string
		method      string
		path        string
		queryParams url.Values
		body        interface{}
		want        int
	}{
		{
			name:        "Happy case: GET Request",
			method:      http.MethodGet,
			path:        "/v1/facilities/facilities/",
			queryParams: queryParam,
			body: &Facility{
				Name:         "Test Facility",
				Description:  "A test facility",
				FacilityType: "HOSPITAL",
				County:       "Meru",
				Country:      "KE",
				Address:      "1200-Meru",
				Coordinates: &Coordinates{
					Latitude:  "30.40338",
					Longitude: "5.17403",
				},
				Contacts: []Contacts{
					{
						ContactType:  "PHONE_NUMBER",
						ContactValue: "+254788223223",
						Role:         "PRIMARY_CONTACT",
					},
				},
				Identifiers: []Identifiers{
					{
						IdentifierType:  "SLADE_CODE",
						IdentifierValue: "3243",
						ValidFrom:       "2022-09-01",
						ValidTo:         "2023-09-01",
					},
				},
				BusinessHours: []BusinessHours{},
			},
			want: http.StatusOK,
		},
		{
			name:        "Happy case: POST Request",
			method:      http.MethodPost,
			path:        "/v1/facilities/facilities/",
			queryParams: nil,
			body: &Facility{
				Name:         "Test Facility",
				Description:  "A test facility",
				FacilityType: "HOSPITAL",
				County:       "Meru",
				Country:      "KE",
				Address:      "1200-Meru",
				Coordinates: &Coordinates{
					Latitude:  "30.40338",
					Longitude: "5.17403",
				},
				Contacts: []Contacts{
					{
						ContactType:  "PHONE_NUMBER",
						ContactValue: "+254788223223",
						Role:         "PRIMARY_CONTACT",
					},
				},
				Identifiers: []Identifiers{
					{
						IdentifierType:  "SLADE_CODE",
						IdentifierValue: "3243",
						ValidFrom:       "2022-09-01",
						ValidTo:         "2023-09-01",
					},
				},
				BusinessHours: []BusinessHours{},
			},
			want: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder(tt.method, tt.path, func(req *http.Request) (*http.Response, error) {
				responseData := map[string]string{"message": "mockResponse"}
				responseJSON, _ := json.Marshal(responseData)

				return httpmock.NewStringResponse(tt.want, string(responseJSON)), nil
			})

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			mockClient := &client{
				httpClient: &http.Client{},
			}

			response, err := mockClient.MakeRequest(ctx, tt.method, tt.path, tt.queryParams, tt.body)
			if err != nil {
				t.Errorf("Error making request: %v", err)
				return
			}

			defer response.Body.Close()

			if response.StatusCode != tt.want {
				t.Errorf("Expected status code %d, got %d", tt.want, response.StatusCode)
			}
		})
	}
}

func TestConfig_tokenURL(t *testing.T) {
	tests := []struct {
		name string
		cfg  Config
		want string
	}{
		{
			name: "builds the realm token endpoint",
			cfg:  Config{KeycloakURL: "https://auth.example.com", Realm: "slade360"},
			want: "https://auth.example.com/realms/slade360/protocol/openid-connect/token",
		},
		{
			name: "tolerates a trailing slash on the server URL",
			cfg:  Config{KeycloakURL: "https://auth.example.com/", Realm: "slade360"},
			want: "https://auth.example.com/realms/slade360/protocol/openid-connect/token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cfg.tokenURL(); got != tt.want {
				t.Errorf("tokenURL() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestNewClient_ValidatesConfig(t *testing.T) {
	cfg := mockConfig()
	cfg.Realm = ""
	cfg.ClientSecret = " "

	_, err := newClient(cfg)
	if err == nil {
		t.Fatal("expected an error for missing config, got nil")
	}

	for _, field := range []string{"Realm", "ClientSecret"} {
		if !strings.Contains(err.Error(), field) {
			t.Errorf("error %q does not name missing field %s", err, field)
		}
	}
}

func TestNewClient_SendsClientCredentials(t *testing.T) {
	cfg := mockConfig()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, cfg.tokenURL(), func(r *http.Request) (*http.Response, error) {
		if err := r.ParseForm(); err != nil {
			t.Fatalf("could not parse token request: %v", err)
		}

		want := map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     cfg.ClientID,
			"client_secret": cfg.ClientSecret,
		}
		for key, value := range want {
			if got := r.PostForm.Get(key); got != value {
				t.Errorf("token request %s = %q, want %q", key, got, value)
			}
		}

		for _, key := range []string{"username", "password"} {
			if r.PostForm.Has(key) {
				t.Errorf("token request should not send %s", key)
			}
		}

		resp := httpmock.NewStringResponse(http.StatusOK, mockTokenResponse)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	if _, err := newClient(cfg); err != nil {
		t.Fatalf("newClient() error = %v", err)
	}
}

func TestNewClient_FailsWhenKeycloakRejectsCredentials(t *testing.T) {
	cfg := mockConfig()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, cfg.tokenURL(), httpmock.NewStringResponder(
		http.StatusUnauthorized,
		`{"error":"unauthorized_client","error_description":"Invalid client or Invalid client credentials"}`,
	))

	_, err := newClient(cfg)
	if err == nil {
		t.Fatal("expected an authentication error, got nil")
	}

	if !strings.Contains(err.Error(), "failed to fetch Keycloak token") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestMakeRequest_SendsBearerToken(t *testing.T) {
	cfg := mockConfig()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	MockAuthenticate(cfg)

	var gotAuthorization string

	httpmock.RegisterResponder(http.MethodGet, cfg.BaseURL+facilitiesPath, func(r *http.Request) (*http.Response, error) {
		gotAuthorization = r.Header.Get("Authorization")
		return httpmock.NewStringResponse(http.StatusOK, `{}`), nil
	})

	c, err := newClient(cfg)
	if err != nil {
		t.Fatalf("newClient() error = %v", err)
	}

	resp, err := c.MakeRequest(context.Background(), http.MethodGet, facilitiesPath, nil, nil)
	if err != nil {
		t.Fatalf("MakeRequest() error = %v", err)
	}
	defer resp.Body.Close()

	if !strings.HasPrefix(gotAuthorization, "Bearer eyJ") {
		t.Errorf("Authorization header = %q, want a Bearer access token", gotAuthorization)
	}
}

func TestMakeRequest_FetchesNewTokenAfterExpiry(t *testing.T) {
	cfg := mockConfig()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// A token that is already inside the renewal window forces a new fetch
	// on every request, the way an expired 5-minute Keycloak token would.
	httpmock.RegisterResponder(http.MethodPost, cfg.tokenURL(), func(r *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(http.StatusOK,
			`{"access_token":"short-lived","expires_in":1,"token_type":"Bearer"}`)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterResponder(http.MethodGet, cfg.BaseURL+facilitiesPath,
		httpmock.NewStringResponder(http.StatusOK, `{}`))

	c, err := newClient(cfg)
	if err != nil {
		t.Fatalf("newClient() error = %v", err)
	}

	resp, err := c.MakeRequest(context.Background(), http.MethodGet, facilitiesPath, nil, nil)
	if err != nil {
		t.Fatalf("MakeRequest() error = %v", err)
	}
	resp.Body.Close()

	tokenCalls := httpmock.GetCallCountInfo()["POST "+cfg.tokenURL()]
	if tokenCalls < 2 {
		t.Errorf("token endpoint called %d times, want at least 2 (initial + renewal)", tokenCalls)
	}
}
