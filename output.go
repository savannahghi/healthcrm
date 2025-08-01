package healthcrm

import (
	"time"

	"github.com/savannahghi/scalarutils"
)

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
	IsAIResult    bool                  `json:"is_ai_result,omitempty"`
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
	ID             string      `json:"id"`
	ProfileID      string      `json:"profile_id"`
	HealthID       string      `json:"health_id,omitempty"`
	Classification MatchResult `json:"classification,omitempty"`
	SladeCode      string      `json:"slade_code"`
}

// FacilityServices is used to get a list of facility Services
type FacilityServices struct {
	Results []*FacilityService `json:"results"`
}

// FacilityOutputs is used to get a list of facilities
type FacilityOutputs struct {
	Results []*FacilityOutput `json:"results"`
}

// Service is used get a service from HealthCRM. These
// are services such as advantage, edi , consumer etc
type Service struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// Profile models how a profile from a service is modelled.
type Profile struct {
	Service    Service `json:"service"`
	Name       string  `json:"full_name"`
	SladeCode  string  `json:"slade_code"`
	ExternalID string  `json:"external_id"`
}

// ProfileIdentifierOutput is used to display profile(s) identifier(s)
type ProfileIdentifierOutput struct {
	IdentifierType  IdentifierType    `json:"identifier_type"`
	IdentifierValue string            `json:"identifier_value"`
	Verified        bool              `json:"verified"`
	ValidFrom       *scalarutils.Date `json:"valid_from,omitempty"`
	ValidTo         *scalarutils.Date `json:"valid_to,omitempty"`
	Profile         Profile           `json:"profile"`
}

// ProfileIdentifierOutputs is used to get a list of identifiers
type ProfileIdentifierOutputs struct {
	Results []*ProfileIdentifierOutput `json:"results"`
}

// ProfileContactOutput is used to display profile(s) contact(s)
type ProfileContactOutput struct {
	ContactType  ContactType       `json:"contact_type"`
	ContactValue string            `json:"contact_value"`
	Verified     bool              `json:"verified"`
	ValidFrom    *scalarutils.Date `json:"valid_from,omitempty"`
	ValidTo      *scalarutils.Date `json:"valid_to,omitempty"`
	Profile      Profile           `json:"profile"`
}

// ProfileContactOutputs is used to get a list of contacts
type ProfileContactOutputs struct {
	Results []*ProfileContactOutput `json:"results"`
}

type UserDetails struct {
	IDNumber    string     `json:"id_number"`
	FullNames   string     `json:"full_names"`
	DateOfBirth string     `json:"date_of_birth"`
	Gender      GenderType `json:"gender"`
}

type RegistryDetails struct {
	ID    string `json:"id"`
	Phone string `json:"phone"`
}

// IDVerificationResult is the result of a verification request
type IDVerificationResult struct {
	ConfidenceScore float64         `json:"confidence_score"`
	UserDetails     UserDetails     `json:"patient_details"`
	RegistryDetails RegistryDetails `json:"client_registry_details"`
}

type PractitionerBusinessHours struct {
	ID             string `json:"id"`
	Day            string `json:"day"`
	OpeningTime    string `json:"opening_time"`
	ClosingTime    string `json:"closing_time"`
	PractitionerID string `json:"practitioner_id"`
}

// IdentifiersOutput is used to display practitioners identifiers
type PractitionerIdentifier struct {
	ID              string                     `json:"id"`
	IdentifierType  PractitionerIdentifierType `json:"identifier_type"`
	IdentifierValue string                     `json:"identifier_value"`
	ValidFrom       string                     `json:"valid_from"`
	ValidTo         string                     `json:"valid_to"`
}

type Practitioners struct {
	Count       int            `json:"count"`
	Next        *string        `json:"next"`
	Previous    *string        `json:"previous"`
	PageSize    int            `json:"page_size"`
	CurrentPage int            `json:"current_page"`
	TotalPages  int            `json:"total_pages"`
	StartIndex  int            `json:"start_index"`
	EndIndex    int            `json:"end_index"`
	Results     []Practitioner `json:"results"`
}

// ContactsOutput is used to show practitioners contacts
type PractitionerContact struct {
	ID           string `json:"id"`
	ContactType  string `json:"contact_type"`
	ContactValue string `json:"contact_value"`
	Role         string `json:"role"`
}

// specialty Identifier
type SpecialtyIdentifier struct {
	ID              string `json:"id"`
	IdentifierType  string `json:"identifier_type"`
	IdentifierValue string `json:"identifier_value"`
	SpecialtyID     string `json:"specialty_id"`
}

type PractitionerSpecialty struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Identifiers []SpecialtyIdentifier `json:"identifiers,omitempty"`
}

type Practitioner struct {
	ID             string                      `json:"id,omitempty"`
	Slug           string                      `json:"slug"`
	Title          string                      `json:"title"`
	FullName       string                      `json:"full_name"`
	FirstName      string                      `json:"first_name"`
	LastName       string                      `json:"last_name"`
	OtherName      string                      `json:"other_name"`
	DateOfBirth    string                      `json:"date_of_birth"`
	Gender         GenderType                  `json:"gender"`
	Country        string                      `json:"country,omitempty"`
	Status         PractitionerStatus          `json:"status,omitempty"`
	Address        string                      `json:"address,omitempty"`
	BusinessHours  []PractitionerBusinessHours `json:"businesshours,omitempty"`
	Coordinates    CoordinatesOutput           `json:"coordinates,omitempty"`
	Distance       *float64                    `json:"distance,omitempty"`
	Contacts       []PractitionerContact       `json:"contacts,omitempty"`
	Identifiers    []PractitionerIdentifier    `json:"identifiers,omitempty"`
	Specialties    []PractitionerSpecialty     `json:"specialties,omitempty"`
	Services       []FacilityService           `json:"services,omitempty"`
	Qualifications string                      `json:"qualifications"`
}
