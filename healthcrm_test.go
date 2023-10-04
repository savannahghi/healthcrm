package healthcrm

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/jarcoal/httpmock"
	"github.com/savannahghi/authutils"
	"github.com/savannahghi/serverutils"
)

// MockAuthenticate mocks a mock login request to obtain a token
func MockAuthenticate() {
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/oauth2/token/", serverutils.MustGetEnvVar("HEALTH_CRM_AUTH_SERVER_ENDPOINT")), func(r *http.Request) (*http.Response, error) {
		resp := authutils.OAUTHResponse{
			Scope:        "",
			ExpiresIn:    3600,
			AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			TokenType:    "Bearer",
		}
		return httpmock.NewJsonResponse(http.StatusCreated, resp)
	})
}

func TestHealthCRMLib_CreateFacility(t *testing.T) {
	type args struct {
		ctx      context.Context
		facility *Facility
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: create facility",
			args: args{
				ctx: context.Background(),
				facility: &Facility{
					Name:         "Test Facility",
					Description:  "A test facility",
					FacilityType: "HOSPITAL",
					County:       "Meru",
					Country:      "KE",
					Address:      "1200-Meru",
					Coordinates: Coordinates{
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
					BusinessHours: []any{},
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to create facility",
			args: args{
				ctx: context.Background(),
				facility: &Facility{
					ID:   gofakeit.UUID(),
					Name: gofakeit.Name(),
				},
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to make request",
			args: args{
				ctx: context.Background(),
				facility: &Facility{
					ID:   gofakeit.UUID(),
					Name: gofakeit.Name(),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Happy case: create facility" {
				path := fmt.Sprintf("%s/v1/facilities/facilities/", BaseURL)
				httpmock.RegisterResponder(http.MethodPost, path, func(r *http.Request) (*http.Response, error) {
					resp := &FacilityOutput{
						ID:           gofakeit.UUID(),
						Name:         gofakeit.BeerName(),
						Description:  gofakeit.HipsterSentence(50),
						FacilityType: "HOSPITAL",
						County:       "Baringo",
						Country:      "KE",
						Address:      "",
						Coordinates: CoordinatesOutput{
							ID:        gofakeit.UUID(),
							Latitude:  30.4556,
							Longitude: 4.54556,
						},
						Contacts:      []Contacts{},
						Identifiers:   []Identifiers{},
						BusinessHours: []any{},
					}
					return httpmock.NewJsonResponse(http.StatusCreated, resp)
				})
			}
			if tt.name == "Sad case: unable to create facility" {
				path := fmt.Sprintf("%s/v1/facilities/facilities/", BaseURL)
				httpmock.RegisterResponder(http.MethodPost, path, func(r *http.Request) (*http.Response, error) {
					resp := &Facility{
						ID:            gofakeit.UUID(),
						Name:          gofakeit.BeerName(),
						Description:   gofakeit.HipsterSentence(50),
						FacilityType:  "HOSPITAL",
						County:        "Baringo",
						Country:       "KE",
						Address:       "",
						Coordinates:   Coordinates{},
						Contacts:      []Contacts{},
						Identifiers:   []Identifiers{},
						BusinessHours: []any{},
					}
					return httpmock.NewJsonResponse(http.StatusBadRequest, resp)
				})
			}
			if tt.name == "Sad case: unable to make request" {
				httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/oauth2/token/", serverutils.MustGetEnvVar("HEALTH_CRM_AUTH_SERVER_ENDPOINT")), func(r *http.Request) (*http.Response, error) {
					resp := authutils.OAUTHResponse{
						Scope:        "",
						ExpiresIn:    3600,
						AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
						RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
						TokenType:    "Bearer",
					}
					return httpmock.NewJsonResponse(http.StatusBadRequest, resp)
				})
			}

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			MockAuthenticate()
			h, err := NewHealthCRMLib()
			if err != nil {
				t.Errorf("unable to initiLize sdk: %v", err)
			}
			_, err = h.CreateFacility(tt.args.ctx, tt.args.facility)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.CreateFacility() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
