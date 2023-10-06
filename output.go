package healthcrm

// Facility is the hospitals model used to show facility details
type FacilityOutput struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	FacilityType  string            `json:"facility_type"`
	County        string            `json:"county"`
	Country       string            `json:"country"`
	Address       string            `json:"address"`
	Coordinates   CoordinatesOutput `json:"coordinates"`
	Contacts      []Contacts        `json:"contacts"`
	Identifiers   []Identifiers     `json:"identifiers"`
	BusinessHours []any             `json:"businesshours"`
}

// CoordinatesOutput is used to show the geographical's location data class of a facility
type CoordinatesOutput struct {
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
