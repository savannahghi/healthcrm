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

const (
	facilitiesPath = "/v1/facilities/facilities/"
	// TODO: use an environment variable
	crmServiceCode = "05"
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

// GetServices retrieves a list of healthcare services provided by facilities
// that are owned by a specific SIL service, such as Mycarehub or Advantage.
func (h *HealthCRMLib) GetServices(ctx context.Context, pagination *Pagination) (*FacilityServicePage, error) {
	path := "/v1/facilities/services/"

	queryParams := make(map[string]string)

	if pagination != nil {
		queryParams["page_size"] = pagination.PageSize
		queryParams["page"] = pagination.Page
	}

	queryParams["crm_service_code"] = crmServiceCode

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
func (h *HealthCRMLib) GetFacilitiesOfferingAService(ctx context.Context, serviceID string, pagination *Pagination) (*FacilityPage, error) {
	path := "/v1/facilities/facilities/"

	queryParams := make(map[string]string)
	queryParams["service"] = serviceID

	if pagination != nil {
		queryParams["page_size"] = pagination.PageSize
		queryParams["page"] = pagination.Page
	}

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

// CreateService is used to create a new service in health crm
func (h *HealthCRMLib) CreateService(ctx context.Context, input FacilityServiceInput) (*FacilityService, error) {
	path := "/v1/facilities/services/"

	response, err := h.client.MakeRequest(ctx, http.MethodPost, path, nil, input)
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

	var output *FacilityService

	err = json.Unmarshal(respBytes, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// LinkServiceToFacility is used to link a service to a facility
func (h *HealthCRMLib) LinkServiceToFacility(ctx context.Context, facilityID string, input []*FacilityServiceInput) (*FacilityService, error) {
	path := fmt.Sprintf("/v1/facilities/facilities/%s/add_services/", facilityID)

	response, err := h.client.MakeRequest(ctx, http.MethodPost, path, nil, input)
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

	var output *FacilityService

	err = json.Unmarshal(respBytes, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// GetFacilities retrieves a list of facilities associated with MyCareHub
// stored in HealthCRM. The method allows for filtering facilities by location proximity and services offered.
//
// Parameters:
//   - location: A Location struct that represents the reference location.
//     If provided, facilities will be filtered based on proximity
//     to this location.
//   - pagination: A Pagination struct containing options for paginating
//     the results.
//   - serviceIDs: A parameter that allows specifying one or more
//     service IDs. Facilities offering these services will be
//     included in the results. You can pass multiple service
//     IDs as separate arguments (e.g., GetFacilities(ctx, location, pagination, []string{"1234", "178"})).
//
// Usage:
// Example 1: Retrieve facilities by location and service IDs:
// --> E.g If we are searching with service ID that represents Chemotherapy, the response
// will be a list of facilities that offer Chemotherapy and it will be ordered with proximity
//
// Example 2: Retrieve facilities by location without specifying services:
// This will return a list of all facilities ordered by the proximity
//
// Example 3: Retrieve all facilities without specifying location or services:
func (h *HealthCRMLib) GetFacilities(ctx context.Context, location *Coordinates, serviceIDs []string, pagination *Pagination) (*FacilityPage, error) {
	queryParams := make(map[string]string)

	if pagination != nil {
		queryParams["page_size"] = pagination.PageSize
		queryParams["page"] = pagination.Page
	}

	if location != nil {
		queryParams["ref_location"] = location.ToString()
	}

	for _, id := range serviceIDs {
		queryParams["service"] = id
	}

	queryParams["crm_service_code"] = crmServiceCode

	response, err := h.client.MakeRequest(ctx, http.MethodGet, facilitiesPath, queryParams, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		// TODO: Get the exact error message (should be formatted well)
		return nil, errors.New(string(respBytes))
	}

	var facilityPage *FacilityPage

	err = json.Unmarshal(respBytes, &facilityPage)
	if err != nil {
		return nil, err
	}

	return facilityPage, nil
}
