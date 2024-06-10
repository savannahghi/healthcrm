package healthcrm

import "time"

// FacilityPage is the hospitals model used to show facility details
type FacilityPage struct {
	Count       int              `json:"count"`
	Next        string           `json:"next"`
	Previous    any              `json:"previous"`
	PageSize    int              `json:"page_size"`
	CurrentPage int              `json:"current_page"`
	TotalPages  int              `json:"total_pages"`
	StartIndex  int              `json:"start_index"`
	EndIndex    int              `json:"end_index"`
	Results     []FacilityOutput `json:"results"`
}

// CoordinatesOutput is used to show geographical coordinates
type CoordinatesOutput struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// ContactsOutput is used to show facility contacts
type ContactsOutput struct {
	ID           string `json:"id"`
	ContactType  string `json:"contact_type"`
	ContactValue string `json:"contact_value"`
	Active       bool   `json:"active"`
	Role         string `json:"role"`
	FacilityID   string `json:"facility_id"`
}

// IdentifiersOutput is used to display facility identifiers
type IdentifiersOutput struct {
	ID              string `json:"id"`
	IdentifierType  string `json:"identifier_type"`
	IdentifierValue string `json:"identifier_value"`
	ValidFrom       string `json:"valid_from"`
	ValidTo         string `json:"valid_to"`
	FacilityID      string `json:"facility_id"`
}

// FacilityOutput is used to display facility(ies)
type FacilityOutput struct {
	ID            string                `json:"id,omitempty"`
	Created       time.Time             `json:"created,omitempty"`
	Name          string                `json:"name,omitempty"`
	Description   string                `json:"description,omitempty"`
	FacilityType  string                `json:"facility_type,omitempty"`
	County        string                `json:"county,omitempty"`
	Country       string                `json:"country,omitempty"`
	Coordinates   CoordinatesOutput     `json:"coordinates,omitempty"`
	Distance      float64               `json:"distance,omitempty"`
	Status        string                `json:"status,omitempty"`
	Address       string                `json:"address,omitempty"`
	Contacts      []ContactsOutput      `json:"contacts,omitempty"`
	Identifiers   []IdentifiersOutput   `json:"identifiers,omitempty"`
	BusinessHours []BusinessHoursOutput `json:"businesshours,omitempty"`
	Services      []FacilityService     `json:"services,omitempty"`
}

// BusinessHoursOutput models data that show facility's operational hours
type BusinessHoursOutput struct {
	ID          string `json:"id"`
	Day         string `json:"day"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
	FacilityID  string `json:"facility_id"`
}

// FacilityServicePage models the services offered in a facility
type FacilityServicePage struct {
	Results     []FacilityService `json:"results"`
	Count       int               `json:"count"`
	Next        string            `json:"next"`
	Previous    string            `json:"previous"`
	PageSize    int               `json:"page_size"`
	CurrentPage int               `json:"current_page"`
	TotalPages  int               `json:"total_pages"`
	StartIndex  int               `json:"start_index"`
	EndIndex    int               `json:"end_index"`
}

// FacilityService models the data class that is used to show facility services
type FacilityService struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Identifiers []*ServiceIdentifier `json:"identifiers"`
}

// ServiceIdentifier models the structure of facility's service identifiers
type ServiceIdentifier struct {
	ID              string `json:"id"`
	IdentifierType  string `json:"identifier_type"`
	IdentifierValue string `json:"identifier_value"`
	ServiceID       string `json:"service_id"`
}

// ProfileOutput is used to display profile(s)
type ProfileOutput struct {
	ID        string `json:"id"`
	ProfileID string `json:"profile_id"`
	HealthID  string `json:"health_id,omitempty"`
	SladeCode string `json:"slade_code"`
}

// FacilityServices is used to get a list of Facility Services
type FacilityServices struct {
	Results []*FacilityService `json:"results"`
}
