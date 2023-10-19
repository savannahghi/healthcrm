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
							Latitude:  30.4556,
							Longitude: 4.54556,
						},
						Contacts:    []ContactsOutput{},
						Identifiers: []IdentifiersOutput{},
						BusinessHours: []BusinessHoursOutput{
							{
								ID:          gofakeit.UUID(),
								Day:         "MONDAY",
								OpeningTime: "08:00:01",
								ClosingTime: "18:00:01",
								FacilityID:  gofakeit.UUID(),
							},
						},
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
						Coordinates:   &Coordinates{},
						Contacts:      []Contacts{},
						Identifiers:   []Identifiers{},
						BusinessHours: []BusinessHours{},
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
				t.Errorf("unable to initialize sdk: %v", err)
			}
			_, err = h.CreateFacility(tt.args.ctx, tt.args.facility)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.CreateFacility() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHealthCRMLib_GetFacilities(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: fetch facility(ies)",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to fetch facility(ies)",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to make request",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Happy case: fetch facility(ies)" {
				path := fmt.Sprintf("%s/v1/facilities/facilities/", BaseURL)
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					resp := &FacilityOutput{
						ID:           gofakeit.UUID(),
						Name:         gofakeit.BeerName(),
						Description:  gofakeit.HipsterSentence(50),
						FacilityType: "HOSPITAL",
						County:       "Baringo",
						Country:      "KE",
						Address:      "",
						Coordinates: CoordinatesOutput{
							Latitude:  30.4556,
							Longitude: 4.54556,
						},
						Contacts:    []ContactsOutput{},
						Identifiers: []IdentifiersOutput{},
						BusinessHours: []BusinessHoursOutput{
							{
								ID:          gofakeit.UUID(),
								Day:         "MONDAY",
								OpeningTime: "08:00:01",
								ClosingTime: "18:00:01",
								FacilityID:  gofakeit.UUID(),
							},
						},
					}
					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			}

			if tt.name == "Sad case: unable to fetch facility(ies)" {
				path := fmt.Sprintf("%s/v1/facilities/facilities/", BaseURL)
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					return httpmock.NewJsonResponse(http.StatusBadRequest, nil)
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
				t.Errorf("unable to initialize sdk: %v", err)
			}

			_, err = h.GetFacilities(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.GetFacilities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHealthCRMLib_GetFacilityByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: get facility",
			args: args{
				ctx: context.Background(),
				id:  "123",
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to get facility",
			args: args{
				ctx: context.Background(),
				id:  "123",
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to make request",
			args: args{
				ctx: context.Background(),
				id:  "123",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Happy case: get facility" {
				path := fmt.Sprintf("%s/v1/facilities/facilities/%s/", BaseURL, "123")
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					resp := &FacilityOutput{
						ID:           gofakeit.UUID(),
						Name:         gofakeit.BeerName(),
						Description:  gofakeit.HipsterSentence(50),
						FacilityType: "HOSPITAL",
						County:       "Baringo",
						Country:      "KE",
						Address:      "",
						Coordinates: CoordinatesOutput{
							Latitude:  30.4556,
							Longitude: 4.54556,
						},
						Contacts:    []ContactsOutput{},
						Identifiers: []IdentifiersOutput{},
						BusinessHours: []BusinessHoursOutput{
							{
								ID:          gofakeit.UUID(),
								Day:         "MONDAY",
								OpeningTime: "08:00:01",
								ClosingTime: "18:00:01",
								FacilityID:  gofakeit.UUID(),
							},
						},
					}
					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			}

			if tt.name == "Sad case: unable to get facility" {
				path := fmt.Sprintf("%s/v1/facilities/facilities/%s/", BaseURL, "123")
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					return httpmock.NewJsonResponse(http.StatusBadRequest, nil)
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
				t.Errorf("unable to initialize sdk: %v", err)
			}

			_, err = h.GetFacilityByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.GetFacilities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHealthCRMLib_UpdateFacility(t *testing.T) {
	type args struct {
		ctx           context.Context
		id            string
		updatePayload *Facility
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: update facility",
			args: args{
				ctx: context.Background(),
				id:  "123",
				updatePayload: &Facility{
					Name: "Makuyu Level 5 Hospital",
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to update facility",
			args: args{
				ctx: context.Background(),
				id:  "123",
				updatePayload: &Facility{
					Name: "Makuyu Level 5 Hospital",
				},
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to make request",
			args: args{
				ctx: context.Background(),
				id:  gofakeit.UUID(),
				updatePayload: &Facility{
					Name: "Makuyu Level 5 Hospital",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Happy case: update facility" {
				path := fmt.Sprintf("%s/v1/facilities/facilities/%s/", BaseURL, "123")
				httpmock.RegisterResponder(http.MethodPatch, path, func(r *http.Request) (*http.Response, error) {
					resp := &FacilityOutput{
						ID:           gofakeit.UUID(),
						Name:         gofakeit.BeerName(),
						Description:  gofakeit.HipsterSentence(50),
						FacilityType: "HOSPITAL",
						County:       "Baringo",
						Country:      "KE",
						Address:      "",
						Coordinates: CoordinatesOutput{
							Latitude:  30.4556,
							Longitude: 4.54556,
						},
						Contacts:    []ContactsOutput{},
						Identifiers: []IdentifiersOutput{},
						BusinessHours: []BusinessHoursOutput{
							{
								ID:          gofakeit.UUID(),
								Day:         "MONDAY",
								OpeningTime: "08:00:01",
								ClosingTime: "18:00:01",
								FacilityID:  gofakeit.UUID(),
							},
						},
					}
					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			}
			if tt.name == "Sad case: unable to update facility" {
				path := fmt.Sprintf("%s/v1/facilities/facilities/%s/", BaseURL, "123")
				httpmock.RegisterResponder(http.MethodPatch, path, func(r *http.Request) (*http.Response, error) {
					return httpmock.NewJsonResponse(http.StatusBadGateway, nil)
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
				t.Errorf("unable to initialize sdk: %v", err)
			}

			_, err = h.UpdateFacility(tt.args.ctx, tt.args.id, tt.args.updatePayload)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.UpdateFacility() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHealthCRMLib_GetFacilityServices(t *testing.T) {
	type args struct {
		ctx        context.Context
		facilityID string
		pagination *Pagination
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: get all services",
			args: args{
				ctx: context.Background(),
				pagination: &Pagination{
					Page:     "2",
					PageSize: "5",
				},
			},
			wantErr: false,
		},
		{
			name: "Happy case: get services in a facility",
			args: args{
				ctx:        context.Background(),
				facilityID: "1b5baf1a-1aec-48bd-951c-01896e5fe5a8",
				pagination: &Pagination{
					Page:     "2",
					PageSize: "5",
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to get all services",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to get facility services",
			args: args{
				ctx:        context.Background(),
				facilityID: gofakeit.UUID(),
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to make request",
			args: args{
				ctx:        context.Background(),
				facilityID: gofakeit.UUID(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Happy case: get all services" {
				path := fmt.Sprintf("%s/v1/facilities/services/", BaseURL)
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					resp := &FacilityServicePage{
						Results: []FacilityService{
							{
								ID:          gofakeit.UUID(),
								Name:        gofakeit.BeerName(),
								Description: gofakeit.HipsterSentence(56),
								Identifiers: []*ServiceIdentifier{},
							},
						},
					}
					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			}
			if tt.name == "Happy case: get services in a facility" {
				path := fmt.Sprintf("%s/v1/facilities/services/?facility=1b5baf1a-1aec-48bd-951c-01896e5fe5a8&page=2&page_size=5", BaseURL)
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					resp := &FacilityServicePage{
						Results: []FacilityService{
							{
								ID:          gofakeit.UUID(),
								Name:        gofakeit.BeerName(),
								Description: gofakeit.HipsterSentence(56),
								Identifiers: []*ServiceIdentifier{},
							},
						},
					}
					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			}
			if tt.name == "Sad case: unable to get all services" {
				path := fmt.Sprintf("%s/v1/facilities/services/", BaseURL)
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					return httpmock.NewJsonResponse(http.StatusBadGateway, nil)
				})
			}
			if tt.name == "Sad case: unable to get facility services" {
				path := fmt.Sprintf("%s/v1/facilities/services/?facility=1b5baf1a-1aec-48bd-951c-01896e5fe5a8", BaseURL)
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					return httpmock.NewJsonResponse(http.StatusBadGateway, nil)
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
				t.Errorf("unable to initialize sdk: %v", err)
			}

			_, err = h.GetFacilityServices(tt.args.ctx, tt.args.facilityID, tt.args.pagination)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.GetFacilityServices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHealthCRMLib_GetFacilitiesOfferingAService(t *testing.T) {
	type args struct {
		ctx        context.Context
		serviceID  string
		pagination *Pagination
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: get facilities offering a service",
			args: args{
				ctx:       context.Background(),
				serviceID: "227305a7-b9a5-4ca7-a211-71210d68206c",
				pagination: &Pagination{
					Page:     "1",
					PageSize: "20",
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to get facilities offering a service",
			args: args{
				ctx: context.Background(),
				pagination: &Pagination{
					Page:     "2",
					PageSize: "20",
				},
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to make request",
			args: args{
				ctx:       context.Background(),
				serviceID: gofakeit.UUID(),
				pagination: &Pagination{
					Page:     "2",
					PageSize: "20",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Happy case: get facilities offering a service" {
				path := fmt.Sprintf("%s/v1/facilities/facilities/", BaseURL)
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					resp := &FacilityOutput{
						ID:           gofakeit.UUID(),
						Name:         gofakeit.BeerName(),
						Description:  gofakeit.HipsterSentence(50),
						FacilityType: "HOSPITAL",
						County:       "Baringo",
						Country:      "KE",
						Address:      "",
						Coordinates: CoordinatesOutput{
							Latitude:  30.4556,
							Longitude: 4.54556,
						},
						Contacts:    []ContactsOutput{},
						Identifiers: []IdentifiersOutput{},
						BusinessHours: []BusinessHoursOutput{
							{
								ID:          gofakeit.UUID(),
								Day:         "MONDAY",
								OpeningTime: "08:00:01",
								ClosingTime: "18:00:01",
								FacilityID:  gofakeit.UUID(),
							},
						},
					}
					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			}
			if tt.name == "Sad case: unable to get facilities offering a service" {
				path := fmt.Sprintf("%s/v1/facilities/facilities/", BaseURL)
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					return httpmock.NewJsonResponse(http.StatusBadGateway, nil)
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
				t.Errorf("unable to initialize sdk: %v", err)
			}

			_, err = h.GetFacilitiesOfferingAService(tt.args.ctx, tt.args.serviceID, tt.args.pagination)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.GetFacilitiesOfferingAService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHealthCRMLib_CreateService(t *testing.T) {
	type args struct {
		ctx   context.Context
		input FacilityServiceInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: create a service",
			args: args{
				ctx: context.Background(),
				input: FacilityServiceInput{
					Name:        "Oxygen Desaturation",
					Description: "Oxygen desaturation",
					Identifiers: []*ServiceIdentifierInput{
						{
							IdentifierType:  "CIEL",
							IdentifierValue: "158211",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to create a service",
			args: args{
				ctx: context.Background(),
				input: FacilityServiceInput{
					Name:        "Oxygen Desaturation",
					Description: "Oxygen desaturation",
					Identifiers: []*ServiceIdentifierInput{
						{
							IdentifierType:  "CIEL",
							IdentifierValue: "158211",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to make request",
			args: args{
				ctx: context.Background(),
				input: FacilityServiceInput{
					Name:        "Oxygen Desaturation",
					Description: "Oxygen desaturation",
					Identifiers: []*ServiceIdentifierInput{
						{
							IdentifierType:  "CIEL",
							IdentifierValue: "158211",
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Happy case: create a service" {
				path := fmt.Sprintf("%s/v1/facilities/services/", BaseURL)
				httpmock.RegisterResponder(http.MethodPost, path, func(r *http.Request) (*http.Response, error) {
					resp := &FacilityService{
						ID:          gofakeit.UUID(),
						Name:        "Oxygen",
						Description: "158211",
						Identifiers: []*ServiceIdentifier{
							{
								ID:              gofakeit.UUID(),
								IdentifierType:  "CIEL",
								IdentifierValue: "158211",
								ServiceID:       gofakeit.UUID(),
							},
						},
					}
					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			}
			if tt.name == "Sad case: unable to create a service" {
				path := fmt.Sprintf("%s/v1/facilities/services/", BaseURL)
				httpmock.RegisterResponder(http.MethodPost, path, func(r *http.Request) (*http.Response, error) {
					return httpmock.NewJsonResponse(http.StatusBadGateway, nil)
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
				t.Errorf("unable to initialize sdk: %v", err)
			}
			_, err = h.CreateService(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.CreateService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHealthCRMLib_LinkServiceToFacility(t *testing.T) {
	type args struct {
		ctx        context.Context
		facilityID string
		input      []*FacilityServiceInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: link facility to service",
			args: args{
				ctx:        context.Background(),
				facilityID: "b6792568-564f-41ca-b951-69fae05e6ca1",
				input: []*FacilityServiceInput{
					{
						Name:        "Renal Pain",
						Description: "Renal Pain Description",
						Identifiers: []*ServiceIdentifierInput{
							{
								IdentifierType:  "CIEL",
								IdentifierValue: "127681",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to link facility to service",
			args: args{
				ctx:        context.Background(),
				facilityID: "b6792568-564f-41ca-b951-69fae05e6ca1",
				input: []*FacilityServiceInput{
					{
						Name:        "Oxygen Desaturation",
						Description: "Oxygen desaturation",
						Identifiers: []*ServiceIdentifierInput{
							{
								IdentifierType:  "CIEL",
								IdentifierValue: "158211",
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to make request",
			args: args{
				ctx:        context.Background(),
				facilityID: gofakeit.UUID(),
				input: []*FacilityServiceInput{
					{
						Name:        "Oxygen Desaturation",
						Description: "Oxygen desaturation",
						Identifiers: []*ServiceIdentifierInput{
							{
								IdentifierType:  "CIEL",
								IdentifierValue: "158211",
							},
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Happy case: link facility to service" {
				path := fmt.Sprintf("%s/v1/facilities/facilities/b6792568-564f-41ca-b951-69fae05e6ca1/add_services/", BaseURL)
				httpmock.RegisterResponder(http.MethodPost, path, func(r *http.Request) (*http.Response, error) {
					resp := &FacilityService{
						ID:          gofakeit.UUID(),
						Name:        "Oxygen",
						Description: "158211",
						Identifiers: []*ServiceIdentifier{
							{
								ID:              gofakeit.UUID(),
								IdentifierType:  "CIEL",
								IdentifierValue: "158211",
								ServiceID:       gofakeit.UUID(),
							},
						},
					}
					return httpmock.NewJsonResponse(http.StatusCreated, resp)
				})
			}
			if tt.name == "Sad case: unable to link facility to service" {
				path := fmt.Sprintf("%s/v1/facilities/facilities/b6792568-564f-41ca-b951-69fae05e6ca1/add_services/", BaseURL)
				httpmock.RegisterResponder(http.MethodPost, path, func(r *http.Request) (*http.Response, error) {
					return httpmock.NewJsonResponse(http.StatusBadGateway, nil)
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
				t.Errorf("unable to initialize sdk: %v", err)
			}

			_, err = h.LinkServiceToFacility(tt.args.ctx, tt.args.facilityID, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.LinkServiceToFacility() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
