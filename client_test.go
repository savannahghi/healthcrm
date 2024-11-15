package healthcrm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/savannahghi/authutils"
)

// MockAuthUtilsLib is a mock implementation of the authUtilsLib interface
type MockAuthUtilsLib struct{}

// Authenticate mocks implementation of authutil's library
func (m *MockAuthUtilsLib) Authenticate() (*authutils.OAUTHResponse, error) {
	return &authutils.OAUTHResponse{
		AccessToken: "mockAccessToken",
	}, nil
}

// Authenticate mocks implementation of authutil's library
func (m *MockAuthUtilsLib) RefreshToken(ctx context.Context, refreshToken string) (*authutils.OAUTHResponse, error) {
	return &authutils.OAUTHResponse{
		AccessToken: "mockAccessToken",
	}, nil
}

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
				authClient: &MockAuthUtilsLib{},
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
