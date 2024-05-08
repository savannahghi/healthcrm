package healthcrm

import (
	"fmt"
	"io"
	"strconv"
)

type IdentifierType string

type ContactType string

const (
	// Identifier types
	IdentifierTypeNationalID    IdentifierType = "NATIONAL_ID"
	IdentifierTypePassportNo    IdentifierType = "PASSPORT_NO"
	IdentifierTypeMilitaryID    IdentifierType = "MILITARY_ID"
	IdentifierTypeAlienID       IdentifierType = "ALIEN_ID"
	IdentifierTypeNHIFNo        IdentifierType = "NHIF_NO"
	IdentifierTypePatientNo     IdentifierType = "PATIENT_NO"
	IdentifierTypePayerMemberNo IdentifierType = "PAYER_MEMBER_NO"
	IdentifierTypeSmartMemberNo IdentifierType = "SMART_MEMBER_NO"
	IdentifierTypeFHIRPatientID IdentifierType = "FHIR_PATIENT_ID"
	IdentifierTypeERPCustomerID IdentifierType = "ERP_CUSTOMER_ID"
	IdentifierTypeCCCNumber     IdentifierType = "CCC_NUMBER"
)

const (
	ContactTypePhoneNumber ContactType = "PHONE_NUMBER"
	ContactTypeEmail       ContactType = "EMAIL"
)

// IsValid returns true if a contact type is valid
func (f ContactType) IsValid() bool {
	switch f {
	case ContactTypePhoneNumber, ContactTypeEmail:
		return true
	default:
		return false
	}
}

// String converts the contact type enum to a string
func (f ContactType) String() string {
	return string(f)
}

// UnmarshalGQL converts the supplied value to a contact type.
func (f *ContactType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*f = ContactType(str)
	if !f.IsValid() {
		return fmt.Errorf("%s is not a valid ContactType type", str)
	}

	return nil
}

// MarshalGQL writes the contact type to the supplied writer
func (f ContactType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(f.String()))
}

// IsValid returns true if an identifier type is valid
func (f IdentifierType) IsValid() bool {
	switch f {
	case
		IdentifierTypeNationalID,
		IdentifierTypePassportNo,
		IdentifierTypeMilitaryID,
		IdentifierTypeAlienID,
		IdentifierTypeNHIFNo,
		IdentifierTypePatientNo,
		IdentifierTypePayerMemberNo,
		IdentifierTypeSmartMemberNo,
		IdentifierTypeFHIRPatientID,
		IdentifierTypeERPCustomerID,
		IdentifierTypeCCCNumber:
		return true
	default:
		return false
	}
}

// String converts the identifier type enum to a string
func (f IdentifierType) String() string {
	return string(f)
}

// UnmarshalGQL converts the supplied value to an identifier type.
func (f *IdentifierType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*f = IdentifierType(str)
	if !f.IsValid() {
		return fmt.Errorf("%s is not a valid IdentifierType type", str)
	}

	return nil
}

// MarshalGQL writes the identifier type to the supplied writer
func (f IdentifierType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(f.String()))
}
