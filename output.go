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
	ID            string              `json:"id"`
	Created       time.Time           `json:"created"`
	Name          string              `json:"name"`
	Description   string              `json:"description"`
	FacilityType  string              `json:"facility_type"`
	County        string              `json:"county"`
	Country       string              `json:"country"`
	Coordinates   CoordinatesOutput   `json:"coordinates"`
	Status        string              `json:"status"`
	Address       string              `json:"address"`
	Contacts      []ContactsOutput    `json:"contacts"`
	Identifiers   []IdentifiersOutput `json:"identifiers"`
	BusinessHours []any               `json:"businesshours"`
}

// FacilityContactOutput is used to show facility's contact
type FacilityContactOutput struct {
	Count       int              `json:"count"`
	Next        string           `json:"next"`
	Previous    interface{}      `json:"previous"`
	PageSize    int              `json:"page_size"`
	CurrentPage int              `json:"current_page"`
	TotalPages  int              `json:"total_pages"`
	StartIndex  int              `json:"start_index"`
	EndIndex    int              `json:"end_index"`
	Results     []ContactsOutput `json:"results"`
}
