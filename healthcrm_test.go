package healthcrm

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/jarcoal/httpmock"
	"github.com/savannahghi/authutils"
	"github.com/savannahghi/enumutils"
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
		ctx             context.Context
		location        *Coordinates
		serviceIDs      []string
		pagination      *Pagination
		searchParameter string
		crmServiceCode  string
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
				location: &Coordinates{
					Latitude:  "-1.29",
					Longitude: "36.79",
				},
				serviceIDs: []string{"1234"},
				pagination: &Pagination{
					Page:     "1",
					PageSize: "10",
				},
				crmServiceCode: "05",
			},
			wantErr: false,
		},
		{
			name: "Happy case: fetch facilities",
			args: args{
				ctx: context.Background(),
				location: &Coordinates{
					Latitude:  "-1.29",
					Longitude: "36.79",
				},
				serviceIDs: []string{"1234", "4567"},
				pagination: &Pagination{
					Page:     "1",
					PageSize: "10",
				},
				crmServiceCode: "05",
			},
			wantErr: false,
		},
		{
			name: "Happy case: search facility by service name",
			args: args{
				ctx: context.Background(),
				location: &Coordinates{
					Latitude:  "-1.29",
					Longitude: "36.79",
				},
				searchParameter: "prep",
				pagination: &Pagination{
					Page:     "1",
					PageSize: "10",
				},
				crmServiceCode: "05",
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
		{
			name: "Sad Case: Pass both service IDs and search parameter",
			args: args{
				ctx:             context.Background(),
				serviceIDs:      []string{"1234"},
				searchParameter: "Nairobi",
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
			if tt.name == "Happy case: fetch facilities" {
				path := fmt.Sprintf("%s/v1/facilities/facilities/", BaseURL)
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					service1 := &FacilityOutput{
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
					service2 := &FacilityOutput{
						ID:           gofakeit.UUID(),
						Name:         gofakeit.BeerName(),
						Description:  gofakeit.HipsterSentence(50),
						FacilityType: "HOSPITAL",
						County:       "Nairobi",
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

					resp := &FacilityPage{
						Results: []FacilityOutput{*service1, *service2},
					}

					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			}

			if tt.name == "Happy case: search facility by service name" {
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

			_, err = h.GetFacilities(tt.args.ctx, tt.args.location, tt.args.serviceIDs, tt.args.searchParameter, tt.args.pagination, tt.args.crmServiceCode)
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

func TestHealthCRMLib_GetServices(t *testing.T) {
	type args struct {
		ctx            context.Context
		facilityID     string
		pagination     *Pagination
		crmServiceCode string
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
				crmServiceCode: "05",
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

			_, err = h.GetServices(tt.args.ctx, tt.args.pagination, tt.args.crmServiceCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.GetServices() error = %v, wantErr %v", err, tt.wantErr)
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

func TestHealthCRMLib_GetService(t *testing.T) {
	type args struct {
		ctx       context.Context
		serviceID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: get service",
			args: args{
				ctx:       context.Background(),
				serviceID: "b7142d0f-88a0-436b-976d-4ecc86482107",
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to get a service",
			args: args{
				ctx:       context.Background(),
				serviceID: "b7142d0f-88a0-436b-976d-4ecc86482107",
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to make request",
			args: args{
				ctx:       context.Background(),
				serviceID: "b7142d0f-88a0-436b-976d-4ecc86482107",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Happy case: get service" {
				path := fmt.Sprintf("%s/v1/facilities/services/b7142d0f-88a0-436b-976d-4ecc86482107", BaseURL)
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
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
			if tt.name == "Sad case: unable to get a service" {
				path := fmt.Sprintf("%s/v1/facilities/services/b7142d0f-88a0-436b-976d-4ecc86482107", BaseURL)
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

			_, err = h.GetService(tt.args.ctx, tt.args.serviceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.GetService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHealthCRMLib_CreateProfile(t *testing.T) {
	type args struct {
		ctx     context.Context
		profile *ProfileInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case: Create Profile",
			args: args{
				ctx: context.Background(),
				profile: &ProfileInput{
					ProfileID:     gofakeit.UUID(),
					FirstName:     "TestProfile",
					LastName:      "BikoTest",
					OtherName:     "SteveTest",
					DateOfBirth:   gofakeit.Date().String(),
					Gender:        "MALE",
					EnrolmentDate: "2023-09-01",
					SladeCode:     "6000",
					ServiceCode:   "50",
					Contacts: []*ProfileContactInput{
						{
							ContactType:  "PHONE_NUMBER",
							ContactValue: "+254788223223",
						},
					},
					Identifiers: []*ProfileIdentifierInput{
						{
							IdentifierType:  "SLADE_CODE",
							IdentifierValue: "3243",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Sad Case: Unable To Create Profile",
			args: args{
				ctx: context.Background(),
				profile: &ProfileInput{
					FirstName: gofakeit.FirstName(),
					LastName:  gofakeit.LastName(),
				},
			},
			wantErr: true,
		},
		{
			name: "Sad Case: Unable To Make Request",
			args: args{
				ctx: context.Background(),
				profile: &ProfileInput{
					FirstName: gofakeit.FirstName(),
					LastName:  gofakeit.LastName(),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Happy Case: Create Profile" {
				path := fmt.Sprintf("%s/v1/identities/profiles/", BaseURL)
				httpmock.RegisterResponder(http.MethodPost, path, func(r *http.Request) (*http.Response, error) {
					resp := &ProfileOutput{
						ID:        gofakeit.UUID(),
						ProfileID: gofakeit.UUID(),
						HealthID:  "50932",
						SladeCode: "50202",
					}
					return httpmock.NewJsonResponse(http.StatusAccepted, resp)
				})
			}
			if tt.name == "Sad Case: Unable To Create Profile" {
				path := fmt.Sprintf("%s/v1/identities/profiles/", BaseURL)
				httpmock.RegisterResponder(http.MethodPost, path, func(r *http.Request) (*http.Response, error) {
					resp := &ProfileInput{
						ProfileID:     gofakeit.UUID(),
						FirstName:     gofakeit.FirstName(),
						LastName:      gofakeit.LastName(),
						OtherName:     gofakeit.BeerName(),
						DateOfBirth:   gofakeit.Date().String(),
						Gender:        ConvertEnumutilsGenderToCRMGender(enumutils.GenderAgender),
						EnrolmentDate: gofakeit.Date().String(),
						SladeCode:     "50202",
						ServiceCode:   "50",
						Contacts:      []*ProfileContactInput{},
						Identifiers:   []*ProfileIdentifierInput{},
					}
					return httpmock.NewJsonResponse(http.StatusBadRequest, resp)
				})
			}
			if tt.name == "Sad Case: Unable To Make Request" {
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
			_, err = h.CreateProfile(tt.args.ctx, tt.args.profile)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.CreateProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHealthCRMLib_GetMultipleServices(t *testing.T) {
	type args struct {
		ctx         context.Context
		servicesIDs []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: get list of services",
			args: args{
				ctx: context.Background(),
				servicesIDs: []string{
					"0fee2792-dffc-40d3-a744-2a70732b1053",
					"56c62083-c7b4-4055-8d44-6cc7446ac1d0",
					"8474ea55-8ede-4bc6-aa67-f53ed5456a03",
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: no service IDs provided(empty list)",
			args: args{
				ctx:         context.Background(),
				servicesIDs: []string{},
			},
			wantErr: true,
		},
		{
			name: "Sad case: no service IDs provided(nil)",
			args: args{
				ctx:         context.Background(),
				servicesIDs: nil,
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to get services",
			args: args{
				ctx: context.Background(),
				servicesIDs: []string{
					"0fee2792-dffc-40d3-a744-2a70732b1053",
					"56c62083-c7b4-4055-8d44-6cc7446ac1d0",
					"8474ea55-8ede-4bc6-aa67-f53ed5456a03",
				},
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to make request",
			args: args{
				ctx: context.Background(),
				servicesIDs: []string{
					"0fee2792-dffc-40d3-a744-2a70732b1053",
					"56c62083-c7b4-4055-8d44-6cc7446ac1d0",
					"8474ea55-8ede-4bc6-aa67-f53ed5456a03",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("%s/v1/facilities/services?service_ids=0fee2792-dffc-40d3-a744-2a70732b1053,56c62083-c7b4-4055-8d44-6cc7446ac1d0,8474ea55-8ede-4bc6-aa67-f53ed5456a03", BaseURL)

			if tt.name == "Happy case: get list of services" {
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					resp := FacilityServices{
						Results: []*FacilityService{
							{
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
							},
							{
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
							},
						},
					}
					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			}

			if tt.name == "Sad case: unable to get a services" {
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

			_, err = h.GetMultipleServices(tt.args.ctx, tt.args.servicesIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.GetMultipleServices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHealthCRMLib_GetMultipleFacilities(t *testing.T) {
	type args struct {
		ctx         context.Context
		facilityIDs []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: get list of facilities",
			args: args{
				ctx: context.Background(),
				facilityIDs: []string{
					"556a1dd9-fbb5-40c2-a623-dde9a2335597",
					"7f59c528-8d9e-4a97-a9e5-bea7d7938c0e",
					"b8246d32-b9e7-422c-b3bb-a1066dec8561",
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: no facility IDs provided(empty list)",
			args: args{
				ctx:         context.Background(),
				facilityIDs: []string{},
			},
			wantErr: true,
		},
		{
			name: "Sad case: no facility IDs provided(nil)",
			args: args{
				ctx:         context.Background(),
				facilityIDs: nil,
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to get services",
			args: args{
				ctx: context.Background(),
				facilityIDs: []string{
					"556a1dd9-fbb5-40c2-a623-dde9a2335597",
					"7f59c528-8d9e-4a97-a9e5-bea7d7938c0e",
					"b8246d32-b9e7-422c-b3bb-a1066dec8561",
				},
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to make request",
			args: args{
				ctx: context.Background(),
				facilityIDs: []string{
					"556a1dd9-fbb5-40c2-a623-dde9a2335597",
					"7f59c528-8d9e-4a97-a9e5-bea7d7938c0e",
					"b8246d32-b9e7-422c-b3bb-a1066dec8561",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("%s/v1/facilities/facilities?facility_ids=556a1dd9-fbb5-40c2-a623-dde9a2335597,7f59c528-8d9e-4a97-a9e5-bea7d7938c0e,b8246d32-b9e7-422c-b3bb-a1066dec8561", BaseURL)

			if tt.name == "Happy case: get list of facilities" {
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					resp := FacilityOutputs{
						Results: []*FacilityOutput{
							{
								ID:          gofakeit.UUID(),
								Name:        "Oxygen",
								Description: "158211",
							},
							{
								ID:          gofakeit.UUID(),
								Name:        "Oxygen",
								Description: "158211",
							},
						},
					}
					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			}

			if tt.name == "Sad case: unable to get services" {
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

			_, err = h.GetMultipleFacilities(tt.args.ctx, tt.args.facilityIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.GetMultipleFacilities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHealthCRMLib_GetPersonIdentifiers(t *testing.T) {
	type args struct {
		ctx            context.Context
		healthID       string
		identifierType IdentifierType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: get list of identifiers",
			args: args{
				ctx:            context.Background(),
				healthID:       "5113010000018400",
				identifierType: "",
			},
			wantErr: false,
		},
		{
			name: "Sad case: no healthID",
			args: args{
				ctx:            context.Background(),
				healthID:       "",
				identifierType: "",
			},
			wantErr: true,
		},
		{
			name: "Sad case: invalid identifier",
			args: args{
				ctx:            context.Background(),
				healthID:       "5113010000018400",
				identifierType: "foo",
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to make request",
			args: args{
				ctx:            context.Background(),
				healthID:       "5113010000018400",
				identifierType: IdentifierTypePatientNo,
			},
			wantErr: true,
		},
		{
			name: "Sad case: invalid status code",
			args: args{
				ctx:            context.Background(),
				healthID:       "5113010000018400",
				identifierType: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("%s/v1/identities/persons/5113010000018400/identifiers/", BaseURL)

			if tt.name == "Happy case: get list of identifiers" {
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					resp := ProfileIdentifierOutputs{
						Results: []*ProfileIdentifierOutput{
							{
								IdentifierType:  IdentifierTypePayerMemberNo,
								IdentifierValue: "12345",
								SladeCode:       "1234",
							},
							{
								IdentifierType:  IdentifierTypeNationalID,
								IdentifierValue: "123456789",
								SladeCode:       "1234",
							},
						},
					}
					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			}

			if tt.name == "Sad case: invalid status code" {
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

			_, err = h.GetPersonIdentifiers(tt.args.ctx, tt.args.healthID, tt.args.identifierType)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.GetMultipleFacilities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHealthCRMLib_GetPersonContacts(t *testing.T) {
	type args struct {
		ctx      context.Context
		healthID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: get list of contacts",
			args: args{
				ctx:      context.Background(),
				healthID: "5113010000018400",
			},
			wantErr: false,
		},
		{
			name: "Sad case: no healthID",
			args: args{
				ctx:      context.Background(),
				healthID: "",
			},
			wantErr: true,
		},
		{
			name: "Sad case: invalid identifier",
			args: args{
				ctx:      context.Background(),
				healthID: "5113010000018400",
			},
			wantErr: true,
		},
		{
			name: "Sad case: unable to make request",
			args: args{
				ctx:      context.Background(),
				healthID: "5113010000018400",
			},
			wantErr: true,
		},
		{
			name: "Sad case: invalid status code",
			args: args{
				ctx:      context.Background(),
				healthID: "5113010000018400",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("%s/v1/identities/persons/5113010000018400/contacts/", BaseURL)

			if tt.name == "Happy case: get list of contacts" {
				httpmock.RegisterResponder(http.MethodGet, path, func(r *http.Request) (*http.Response, error) {
					resp := ProfileContactOutputs{
						Results: []*ProfileContactOutput{
							{
								ContactType:  ContactTypePhoneNumber,
								ContactValue: "+254711234567",
								SladeCode:    "1234",
							},
							{
								ContactType:  ContactTypeEmail,
								ContactValue: "foo@bar.com",
								SladeCode:    "1234",
							},
						},
					}
					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			}

			if tt.name == "Sad case: invalid status code" {
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

			_, err = h.GetPersonContacts(tt.args.ctx, tt.args.healthID)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCRMLib.GetMultipleFacilities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
