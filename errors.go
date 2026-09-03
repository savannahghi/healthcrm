package healthcrm

import (
	"errors"
	"fmt"
)

// ErrIdentifierConflict is the sentinel for an identifier write that clashes
// with one the record already holds. Match it with errors.Is.
//
// Only the facility identifier endpoint produces this today. A facility holds
// at most one identifier per type, so writing a second value for a type it
// already carries is a conflict rather than a repeat. A practitioner may hold
// several values for one type, so its endpoint has no equivalent case.
//
// Retrying will not clear it: the existing identifier has to be updated, which
// this SDK does not currently expose.
var ErrIdentifierConflict = errors.New("identifier conflict")

// ConflictError carries the upstream response body for a conflicting
// identifier write. It unwraps to ErrIdentifierConflict, so callers can branch
// with errors.Is and still read the detail off the body.
type ConflictError struct {
	// Body is the response body returned by health CRM, unmodified.
	Body string
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("identifier conflict: %s", e.Body)
}

// Unwrap allows errors.Is(err, ErrIdentifierConflict) to match.
func (e *ConflictError) Unwrap() error {
	return ErrIdentifierConflict
}
