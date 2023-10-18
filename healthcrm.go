package healthcrm

import (
	"context"
	"encoding/json"
	"errors"
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

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %w", err)
	}

	if response.StatusCode != http.StatusCreated {
		return nil, errors.New(string(respBytes))
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

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(string(respBytes))
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

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(string(respBytes))
	}

	var facilityOutput *FacilityOutput

	err = json.Unmarshal(respBytes, &facilityOutput)
	if err != nil {
		return nil, err
	}

	return facilityOutput, nil
}

// UpdateFacility is used to update facility's data
func (h *HealthCRMLib) UpdateFacility(ctx context.Context, id string, updatePayload *Facility) (*FacilityOutput, error) {
	path := fmt.Sprintf("/v1/facilities/facilities/%s/", id)
	response, err := h.client.MakeRequest(ctx, http.MethodPatch, path, nil, updatePayload)
	if err != nil {
		return nil, err
	}

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(string(respBytes))
	}

	var facilityOutput *FacilityOutput

	err = json.Unmarshal(respBytes, &facilityOutput)
	if err != nil {
		return nil, err
	}

	return facilityOutput, nil
}

// GetFacilityServices fetches services associated with facility
func (h *HealthCRMLib) GetFacilityServices(ctx context.Context, facilityID string, pagination *Pagination) (*FacilityServicePage, error) {
	path := "/v1/facilities/services/"

	queryParams := make(map[string]string)

	if pagination != nil {
		queryParams["page_size"] = pagination.PageSize
		queryParams["page"] = pagination.Page
	}

	if facilityID != "" {
		queryParams["facility"] = facilityID
	}

	response, err := h.client.MakeRequest(ctx, http.MethodGet, path, queryParams, nil)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(string(respBytes))
	}

	var facilityServicePage FacilityServicePage
	err = json.Unmarshal(respBytes, &facilityServicePage)
	if err != nil {
		return nil, err
	}

	return &facilityServicePage, nil
}

// GetFacilitiesOfferingAService fetches the facilities that offer a particular service
func (h *HealthCRMLib) GetFacilitiesOfferingAService(ctx context.Context, serviceID string) (*FacilityPage, error) {
	path := "/v1/facilities/facilities/"

	queryParams := make(map[string]string)
	queryParams["service"] = serviceID
	response, err := h.client.MakeRequest(ctx, http.MethodGet, path, queryParams, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(string(respBytes))
	}

	var output *FacilityPage

	err = json.Unmarshal(respBytes, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
