package healthcrm

import (
	"fmt"

	"github.com/savannahghi/scalarutils"
)

// Facility is the hospitals data class
type Facility struct {
	ID            string          `json:"id,omitempty"`
	Name          string          `json:"name,omitempty"`
	Description   string          `json:"description,omitempty"`
	FacilityType  string          `json:"facility_type,omitempty"`
	County        string          `json:"county,omitempty"`
	Country       string          `json:"country,omitempty"`
	Address       string          `json:"address,omitempty"`
	Coordinates   *Coordinates    `json:"coordinates,omitempty"`
	Contacts      []Contacts      `json:"contacts,omitempty"`
	Identifiers   []Identifiers   `json:"identifiers,omitempty"`
	BusinessHours []BusinessHours `json:"businesshours,omitempty"`
}

// Coordinates represents geographical coordinates using latitude and longitude.
// Latitude measures the north-south position, while longitude measures
// the east-west position.
type Coordinates struct {
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Radius    string `json:"radius,omitempty"`
}

// ToString returns the location in comma-separated values format.
// The order of values in the string is longitude,latitude.
// The latitude and longitude are formatted up to 5 decimal places.
// For example, if the Location has Latitude 36.79 and Longitude -1.29,
// the returned string will be "-1.29, 36.79".
func (c Coordinates) ToString() (string, error) {
	if c.Latitude == "" || c.Longitude == "" {
		return "", fmt.Errorf("both Latitude and Longitude must be provided to generate the coordinates string")
	}
	return fmt.Sprintf("%v, %v", c.Longitude, c.Latitude), nil
}

// Contacts models facility's model data class
type Contacts struct {
	ContactType  string `json:"contact_type,omitempty"`
	ContactValue string `json:"contact_value,omitempty"`
	Role         string `json:"role,omitempty"`
}

// Identifiers models facility's identifiers; can be MFL Code, Slade Code etc...
type Identifiers struct {
	IdentifierType  string `json:"identifier_type,omitempty"`
	IdentifierValue string `json:"identifier_value,omitempty"`
	ValidFrom       string `json:"valid_from,omitempty"`
	ValidTo         string `json:"valid_to,omitempty"`
}

// BusinessHours models data to store business hours
type BusinessHours struct {
	Day         string `json:"day"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

// Pagination is used to hold pagination values
type Pagination struct {
	Page     string `json:"page"`
	PageSize string `json:"page_size"`
}

// FacilityServiceInput models is used to create a new service
type FacilityServiceInput struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Identifiers []*ServiceIdentifierInput `json:"identifiers"`
}

// ServiceIdentifierInput is used to create an identifier
type ServiceIdentifierInput struct {
	IdentifierType  string `json:"identifier_type"`
	IdentifierValue string `json:"identifier_value"`
}

// ProfileInput is the host of users data or a brief description of a person
type ProfileInput struct {
	ProfileID     string                    `json:"profile_id"`
	FirstName     string                    `json:"first_name"`
	LastName      string                    `json:"last_name"`
	OtherName     string                    `json:"other_name,omitempty"`
	DateOfBirth   string                    `json:"date_of_birth,omitempty"`
	Gender        GenderType                `json:"gender"`
	EnrolmentDate string                    `json:"enrolment_date,omitempty"`
	SladeCode     string                    `json:"slade_code"`
	ServiceCode   string                    `json:"service_code"`
	Contacts      []*ProfileContactInput    `json:"contacts,omitempty"`
	Identifiers   []*ProfileIdentifierInput `json:"identifiers,omitempty"`
}

// ProfileIdentifierInput is used to create profile(s) identifier(s)
type ProfileIdentifierInput struct {
	IdentifierType  IdentifierType    `json:"identifier_type"`
	IdentifierValue string            `json:"identifier_value"`
	Verified        bool              `json:"verified"`
	ValidFrom       *scalarutils.Date `json:"valid_from,omitempty"`
	ValidTo         *scalarutils.Date `json:"valid_to,omitempty"`
}

// ProfileContanctInput is used to create profile(s) contact(s)
type ProfileContactInput struct {
	ContactType  ContactType       `json:"contact_type"`
	ContactValue string            `json:"contact_value"`
	Verified     bool              `json:"verified"`
	ValidFrom    *scalarutils.Date `json:"valid_from,omitempty"`
	ValidTo      *scalarutils.Date `json:"valid_to,omitempty"`
}
