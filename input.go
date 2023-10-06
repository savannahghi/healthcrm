package healthcrm

// Facility is the hospitals data class
type Facility struct {
	ID            string        `json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
	Description   string        `json:"description,omitempty"`
	FacilityType  string        `json:"facility_type,omitempty"`
	County        string        `json:"county,omitempty"`
	Country       string        `json:"country,omitempty"`
	Address       string        `json:"address,omitempty"`
	Coordinates   *Coordinates  `json:"coordinates,omitempty"`
	Contacts      []Contacts    `json:"contacts,omitempty"`
	Identifiers   []Identifiers `json:"identifiers,omitempty"`
	BusinessHours []any         `json:"businesshours,omitempty"`
}

// Coordinates models the geographical's location data class of a facility
type Coordinates struct {
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
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
