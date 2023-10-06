package healthcrm

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/savannahghi/serverutils"
)

var (
	// BaseURL represents the health CRM's base URL
	BaseURL = serverutils.MustGetEnvVar("HEALTH_CRM_BASE_URL")
)

// HealthCRMLib interacts with the healthcrm APIs
type HealthCRMLib struct {
	client *client
}

// NewHealthCRMLib initializes a new instance of healthCRM SDK
func NewHealthCRMLib() (*HealthCRMLib, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}

	return &HealthCRMLib{
		client: client,
	}, nil
}

// CreateFacility is used to create facility in health CRM service
func (h *HealthCRMLib) CreateFacility(ctx context.Context, facility *Facility) (*FacilityOutput, error) {
	path := "/v1/facilities/facilities/"
	response, err := h.client.MakeRequest(ctx, http.MethodPost, path, nil, facility)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unable to create facility in the registry with status code: %v", response.StatusCode)
	}

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %w", err)
	}

	var facilityResponse *FacilityOutput

	err = json.Unmarshal(respBytes, &facilityResponse)
	if err != nil {
		return nil, err
	}

	return facilityResponse, nil
}

// GetFacilities is used to fetch facilities from health crm facility registry
func (h *HealthCRMLib) GetFacilities(ctx context.Context) (*FacilityPage, error) {
	path := "/v1/facilities/facilities/"
	response, err := h.client.MakeRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to fetch facility(ies) in the registry with status code: %v", response.StatusCode)
	}

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %w", err)
	}

	var facilityPage *FacilityPage

	err = json.Unmarshal(respBytes, &facilityPage)
	if err != nil {
		return nil, err
	}

	return facilityPage, nil
}

// GetFacilityByID is used to fetch facilities from health crm facility registry using its ID
func (h *HealthCRMLib) GetFacilityByID(ctx context.Context, id string) (*FacilityOutput, error) {
	path := fmt.Sprintf("/v1/facilities/facilities/%s/", id)
	response, err := h.client.MakeRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to fetch facility from the registry with status code: %v", response.StatusCode)
	}

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %w", err)
	}

	var facilityOutput *FacilityOutput

	err = json.Unmarshal(respBytes, &facilityOutput)
	if err != nil {
		return nil, err
	}

	return facilityOutput, nil
}

// GetFacilityContact is used to fetch facilities contacts
func (h *HealthCRMLib) GetFacilityContact(ctx context.Context, facility string) (*FacilityContactOutput, error) {
	path := "/v1/facilities/contacts/"
	response, err := h.client.MakeRequest(ctx, http.MethodGet, path, nil, facility)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to fetch facility contact from the registry with status code: %v", response.StatusCode)
	}

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %w", err)
	}

	var facilityContactOutput *FacilityContactOutput

	err = json.Unmarshal(respBytes, &facilityContactOutput)
	if err != nil {
		return nil, err
	}

	return facilityContactOutput, nil
}
