package healthcrm

// Facility is the hospitals data class
type Facility struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	FacilityType  string        `json:"facility_type"`
	County        string        `json:"county"`
	Country       string        `json:"country"`
	Address       string        `json:"address"`
	Coordinates   Coordinates   `json:"coordinates"`
	Contacts      []Contacts    `json:"contacts"`
	Identifiers   []Identifiers `json:"identifiers"`
	BusinessHours []any         `json:"businesshours"`
}

// Coordinates models the geographical's location data class of a facility
type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Contacts models facility's model data class
type Contacts struct {
	ContactType  string `json:"contact_type"`
	ContactValue string `json:"contact_value"`
	Role         string `json:"role"`
}

// Identifiers models facility's identifiers; can be MFL Code, Slade Code etc...
type Identifiers struct {
	IdentifierType  string `json:"identifier_type"`
	IdentifierValue string `json:"identifier_value"`
	ValidFrom       string `json:"valid_from"`
	ValidTo         string `json:"valid_to"`
}
