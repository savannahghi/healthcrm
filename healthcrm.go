package healthcrm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/savannahghi/serverutils"
)

var (
	// BaseURL represents the health CRM's base URL
	BaseURL = serverutils.MustGetEnvVar("HEALTH_CRM_BASE_URL")
)

const (
	facilitiesPath = "/v1/facilities/facilities/"
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

	defer response.Body.Close()

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

	defer response.Body.Close()

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

	defer response.Body.Close()

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
func (h *HealthCRMLib) GetServices(ctx context.Context, pagination *Pagination, crmServiceCode string) (*FacilityServicePage, error) {
	path := "/v1/facilities/services/"

	queryParams := url.Values{}

	if pagination != nil {
		queryParams.Add("page_size", pagination.PageSize)
		queryParams.Add("page", pagination.Page)
	}

	queryParams.Add("crm_service_code", crmServiceCode)

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

// GetPractitioners retrieves a list of practitioners associated with a specific CRM service code.
func (h *HealthCRMLib) GetPractitioners(ctx context.Context, filters FilterPractitionersInput) (*Practitioners, error) {
	path := "/v1/practitioners/practitioners/"

	queryParams := url.Values{}

	if filters.CrmServiceCode == "" {
		return nil, errors.New("CRM service code must be provided")
	}

	if filters.Pagination != nil {
		queryParams.Add("page_size", filters.Pagination.PageSize)
		queryParams.Add("page", filters.Pagination.Page)
	}

	if len(filters.Specialty) > 0 && filters.SearchParameter != "" {
		return nil, errors.New("cannot filter by both specialty and search parameter")
	}

	if len(filters.Service) > 0 && filters.SearchParameter != "" {
		return nil, errors.New("cannot filter by both service and search parameter")
	}

	if len(filters.Specialty) > 0 {
		for _, id := range filters.Specialty {
			queryParams.Add("specialty", id)
		}
	}

	if len(filters.Service) > 0 {
		for _, id := range filters.Service {
			queryParams.Add("service", id)
		}
	}

	if filters.IdentifierType != "" && filters.IdentifierValue != "" {
		queryParams.Add("identifier_type", filters.IdentifierType)
		queryParams.Add("identifier_value", filters.IdentifierValue)
	}

	queryParams.Add("crm_service_code", filters.CrmServiceCode)

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

	var practitioners Practitioners
	err = json.Unmarshal(respBytes, &practitioners)
	if err != nil {
		return nil, err
	}

	return &practitioners, nil
}

// GetPractitionerByID retrieves a practitioner by their ID or slug.
func (h *HealthCRMLib) GetPractitionerByID(ctx context.Context, practitionerID string) (*Practitioner, error) {
	path := fmt.Sprintf("/v1/practitioners/practitioners/%s/", practitionerID)
	response, err := h.client.MakeRequest(ctx, http.MethodGet, path, nil, nil)
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

	var practitioner *Practitioner

	err = json.Unmarshal(respBytes, &practitioner)
	if err != nil {
		return nil, err
	}

	return practitioner, nil
}

// GetSpecialties retrieves a list of specialties associated with a specific CRM service code.
func (h *HealthCRMLib) GetSpecialties(ctx context.Context, pagination *Pagination, crmServiceCode string) (*Specialties, error) {
	path := "/v1/practitioners/specialties/"

	queryParams := url.Values{}

	if pagination != nil {
		queryParams.Add("page_size", pagination.PageSize)
		queryParams.Add("page", pagination.Page)
	}

	queryParams.Add("crm_service_code", crmServiceCode)

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

	var specialties Specialties

	err = json.Unmarshal(respBytes, &specialties)
	if err != nil {
		return nil, err
	}

	return &specialties, nil
}

// GetFacilitiesOfferingAService fetches the facilities that offer a particular service
func (h *HealthCRMLib) GetFacilitiesOfferingAService(ctx context.Context, serviceID string, pagination *Pagination) (*FacilityPage, error) {
	path := "/v1/facilities/facilities/"

	queryParams := url.Values{}
	queryParams.Add("service", serviceID)

	if pagination != nil {
		queryParams.Add("page_size", pagination.PageSize)
		queryParams.Add("page", pagination.Page)
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

	defer response.Body.Close()

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

	defer response.Body.Close()

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
//   - searchParameter: A parameter used to search a facility by the facility name or a service name
//     Note that this parameter cannot be passed together with the serviceIDs
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
func (h *HealthCRMLib) GetFacilities(ctx context.Context, filters FilterFacilitiesInput) (*FacilityPage, error) {
	queryParams := url.Values{}

	pagination := filters.Pagination
	location := filters.Location
	serviceIDs := filters.ServiceIDs
	searchParameter := filters.SearchParameter
	crmServiceCode := filters.CrmServiceCode
	identifierType := filters.IdentifierType
	identifierValue := filters.IdentifierValue

	if crmServiceCode == "" {
		return nil, errors.New("CRM service code must be provided")
	}

	if pagination != nil {
		queryParams.Add("page_size", pagination.PageSize)
		queryParams.Add("page", pagination.Page)
	}

	if location != nil {
		coordinateString, err := location.ToString()
		if err != nil {
			return nil, err
		}

		queryParams.Add("ref_location", coordinateString)

		if location.Radius != "" {
			queryParams.Add("distance", location.Radius)
		}
	}

	if len(serviceIDs) > 0 && searchParameter != "" {
		return nil, errors.New("both service IDs and search parameter cannot be provided simultaneously")
	}

	if len(serviceIDs) > 0 {
		for _, id := range serviceIDs {
			queryParams.Add("service", id)
		}
	}

	if searchParameter != "" {
		queryParams.Add("search", searchParameter)
	}

	// Only add identifier type and value  params if both are provided
	if identifierType != "" && identifierValue != "" {
		queryParams.Add("identifier_type", identifierType.String())
		queryParams.Add("identifier_value", identifierValue)
	}

	// If identifier type is provided but value is not, return an error
	if identifierType != "" && identifierValue == "" {
		return nil, errors.New("identifier value must be provided if identifier type is specified")
	}

	queryParams.Add("crm_service_code", crmServiceCode)

	response, err := h.client.MakeRequest(ctx, http.MethodGet, facilitiesPath, queryParams, nil)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

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

// GetService is used to fetch a single service given its ID
func (h *HealthCRMLib) GetService(ctx context.Context, serviceID string) (*FacilityService, error) {
	path := fmt.Sprintf("/v1/facilities/services/%s", serviceID)

	response, err := h.client.MakeRequest(ctx, http.MethodGet, path, nil, nil)
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

	var service FacilityService
	err = json.Unmarshal(respBytes, &service)
	if err != nil {
		return nil, err
	}

	return &service, nil
}

// CreateProfile is used to create profile in health CRM service
func (h *HealthCRMLib) CreateProfile(ctx context.Context, profile *ProfileInput) (*ProfileOutput, error) {
	path := "/v1/identities/profiles/"
	response, err := h.client.MakeRequest(ctx, http.MethodPost, path, nil, profile)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %w", err)
	}

	if response.StatusCode != http.StatusAccepted {
		return nil, errors.New(string(respBytes))
	}

	var profileResponse *ProfileOutput

	err = json.Unmarshal(respBytes, &profileResponse)
	if err != nil {
		return nil, err
	}

	return profileResponse, nil
}

// MatchProfile is used to create profile in health CRM service
func (h *HealthCRMLib) MatchProfile(ctx context.Context, profile *ProfileInput) (*ProfileOutput, error) {
	path := "/v1/identities/profiles/match_profile/"
	response, err := h.client.MakeRequest(ctx, http.MethodPost, path, nil, profile)
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

	var profileResponse *ProfileOutput

	err = json.Unmarshal(respBytes, &profileResponse)
	if err != nil {
		return nil, err
	}

	return profileResponse, nil
}

// GetMultipleServices is used to fetch multiple services
//
// Parameters:
//   - serviceIDs: A parameter that is a list of IDs specifying one or more
//     service IDs. Service identifiers identifying these services will be
//     included in the results. You can **ONLY** pass a single or multiple service
//     IDs which should be of type **UUID** (e.g., GetMultipleServices(ctx, []string{"0fee2792-dffc-40d3-a744-2a70732b1053",
//     "56c62083-c7b4-4055-8d44-6cc7446ac1d0", "8474ea55-8ede-4bc6-aa67-f53ed5456a03"})).
func (h *HealthCRMLib) GetMultipleServices(ctx context.Context, servicesIDs []string) ([]*FacilityService, error) {
	if len(servicesIDs) < 1 || servicesIDs == nil {
		return nil, fmt.Errorf("no service IDs provided")
	}

	searchParameter := servicesIDs[0]
	for idx, id := range servicesIDs {
		if idx == 0 {
			continue
		}

		searchParameter += fmt.Sprintf(",%s", id)
	}

	path := fmt.Sprintf("/v1/facilities/services?service_ids=%s", searchParameter)

	response, err := h.client.MakeRequest(ctx, http.MethodGet, path, nil, nil)
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

	var services *FacilityServices
	err = json.Unmarshal(respBytes, &services)
	if err != nil {
		return nil, err
	}

	return services.Results, nil
}

// GetMultipleFacilities is used to fetch multiple facilities
//
// Parameters:
//   - facilityIDs: A parameter that is a list of IDs specifying one or more
//     facility IDs. Facility identifiers, contacts, services and business hours linked to a facility will be
//     included in the results. You can **ONLY** pass a single or multiple facility
//     IDs which should be of type **UUID** (e.g., GetMultipleFacilities(ctx, []string{"0fee2792-dffc-40d3-a744-2a70732b1053",
//     "56c62083-c7b4-4055-8d44-6cc7446ac1d0", "8474ea55-8ede-4bc6-aa67-f53ed5456a03"})).
func (h *HealthCRMLib) GetMultipleFacilities(ctx context.Context, facilityIDs []string) ([]*FacilityOutput, error) {
	if len(facilityIDs) < 1 || facilityIDs == nil {
		return nil, fmt.Errorf("no facility IDs provided")
	}

	searchParameter := facilityIDs[0]
	for idx, id := range facilityIDs {
		if idx == 0 {
			continue
		}

		searchParameter += fmt.Sprintf(",%s", id)
	}

	path := fmt.Sprintf("/v1/facilities/facilities?facility_ids=%s", searchParameter)

	response, err := h.client.MakeRequest(ctx, http.MethodGet, path, nil, nil)
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

	var facilities *FacilityOutputs
	err = json.Unmarshal(respBytes, &facilities)
	if err != nil {
		return nil, err
	}

	return facilities.Results, nil
}

// GetPersonIdentifiers fetches a persons identifiers using their HealthID, a
// filter for identifier_type can be passed
func (h *HealthCRMLib) GetPersonIdentifiers(ctx context.Context, healthID string, identifierTypes []*IdentifierType) ([]*ProfileIdentifierOutput, error) {
	if healthID == "" {
		return nil, errors.New("no health ID provided")
	}

	path := fmt.Sprintf("/v1/identities/persons/%s/identifiers/", healthID)

	queryParams := url.Values{}

	for _, identifier := range identifierTypes {
		if !identifier.IsValid() {
			return nil, fmt.Errorf("invalid identifier type provided: %s", identifier)
		}

		queryParams.Add("identifier_type", identifier.String())
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

	var identifiers *ProfileIdentifierOutputs
	err = json.Unmarshal(respBytes, &identifiers)
	if err != nil {
		return nil, err
	}

	return identifiers.Results, nil
}

// GetPersonContacts fetches a persons Contacts using their HealthID
func (h *HealthCRMLib) GetPersonContacts(ctx context.Context, healthID string) ([]*ProfileContactOutput, error) {
	if healthID == "" {
		return nil, errors.New("no health ID provided")
	}

	path := fmt.Sprintf("/v1/identities/persons/%s/contacts/", healthID)

	response, err := h.client.MakeRequest(ctx, http.MethodGet, path, nil, nil)
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

	var identifiers *ProfileContactOutputs
	err = json.Unmarshal(respBytes, &identifiers)
	if err != nil {
		return nil, err
	}

	return identifiers.Results, nil
}

func (h *HealthCRMLib) VerifyIdentifierDocument(ctx context.Context, input IDVerificationInput) (*IDVerificationResult, error) {
	path := "/v1/identities/identifiers/verify/"

	response, err := h.client.MakeRequest(ctx, http.MethodPost, path, nil, input)
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

	var result IDVerificationResult

	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
