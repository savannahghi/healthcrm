package healthcrm

import (
	"fmt"
	"io"
	"strconv"
)

type IdentifierType string

type ContactType string

type GenderType string

type MatchResult string

type PractitionerStatus string

type PractitionerIdentifierType string

const (
	// Identifier types
	IdentifierTypeNationalID           IdentifierType = "NATIONAL_ID"
	IdentifierTypePassportNo           IdentifierType = "PASSPORT_NO"
	IdentifierTypeMilitaryID           IdentifierType = "MILITARY_ID"
	IdentifierTypeAlienID              IdentifierType = "ALIEN_ID"
	IdentifierTypeNHIFNo               IdentifierType = "NHIF_NO"
	IdentifierTypePatientNo            IdentifierType = "PATIENT_NO"
	IdentifierTypePayerMemberNo        IdentifierType = "PAYER_MEMBER_NO"
	IdentifierTypeSmartMemberNo        IdentifierType = "SMART_MEMBER_NO"
	IdentifierTypeFHIRPatientID        IdentifierType = "FHIR_PATIENT_ID"
	IdentifierTypeERPCustomerID        IdentifierType = "ERP_CUSTOMER_ID"
	IdentifierTypeCCCNumber            IdentifierType = "CCC_NUMBER"
	IdentifierTypeRefugeeID            IdentifierType = "REFUGEE_ID"
	IdentifierTypeBirthCertificateNo   IdentifierType = "BIRTH_CERTIFICATE_NO"
	IdentifierTypeMandateNo            IdentifierType = "MANDATE_NO"
	IdentifierTypeClientRegistryNo     IdentifierType = "CLIENT_REGISTRY_NO"
	IdentifierTypeDRChronoChartID      IdentifierType = "DR_CHRONO_CHART_ID"
	IdentifierTypeBirthNotificationNo  IdentifierType = "BIRTH_NOTIFICATION_NO"
	IdentifierTypeSHANumber            IdentifierType = "SHA_NUMBER"
	IdentifierTypeHouseholdNumber      IdentifierType = "HOUSEHOLD_NUMBER"
	IdentifierTypeTemporaryID          IdentifierType = "TEMPORARY_ID"
	IdentifierTypeTemporaryDependentID IdentifierType = "TEMPORARY_DEPENDENT_ID"
)

const (
	ContactTypePhoneNumber ContactType = "PHONE_NUMBER"
	ContactTypeEmail       ContactType = "EMAIL"
)

const (
	GenderTypeMale   GenderType = "MALE"
	GenderTypeFemale GenderType = "FEMALE"
	GenderTypeOther  GenderType = "OTHER"
	// GenderTypeASKU stands for Asked but Unknown
	GenderTypeASKU GenderType = "ASKU"
	// GenderTypeUNK stands for Unknown
	GenderTypeUNK GenderType = "UNK"
)

const (
	MatchResultMatch         MatchResult = "MATCH"
	MatchResultPossibleMatch MatchResult = "POSSIBLE_MATCH"
	MatchResultNoMatch       MatchResult = "NO_MATCH"
)

const (
	PractitionerStatusDraft     PractitionerStatus = "DRAFT"
	PractitionerStatusPublished PractitionerStatus = "PUBLISHED"
)

const (
	//nolint:all
	SLADE_CODE                PractitionerIdentifierType = "SLADE_CODE"//nolint:all
	SHA_SLADE_CODE            PractitionerIdentifierType = "SHA_SLADE_CODE"//nolint:all
	NATIONAL_ID               PractitionerIdentifierType = "NATIONAL_ID"//nolint:all
	PASSPORT                  PractitionerIdentifierType = "PASSPORT"//nolint:all
	KMPDC_REGISTRATION_NUMBER PractitionerIdentifierType = "KMPDC_REGISTRATION_NUMBER"//nolint:all
	KMPDC_LICENSE_NUMBER      PractitionerIdentifierType = "KMPDC_LICENSE_NUMBER"//nolint:all
	ALIEN_ID                  PractitionerIdentifierType = "ALIEN_ID"//nolint:all
	REFUGEE_ID                PractitionerIdentifierType = "REFUGEE_ID"//nolint:all
	CLIENT_REGISTRY_ID        PractitionerIdentifierType = "CLIENT_REGISTRY_ID"//nolint:all
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
		IdentifierTypeCCCNumber,
		IdentifierTypeRefugeeID,
		IdentifierTypeBirthCertificateNo,
		IdentifierTypeMandateNo,
		IdentifierTypeClientRegistryNo,
		IdentifierTypeDRChronoChartID,
		IdentifierTypeBirthNotificationNo,
		IdentifierTypeSHANumber,
		IdentifierTypeHouseholdNumber,
		IdentifierTypeTemporaryID,
		IdentifierTypeTemporaryDependentID:
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

// IsValid returns true if a Gender type is valid
func (f GenderType) IsValid() bool {
	switch f {
	case GenderTypeMale, GenderTypeFemale, GenderTypeOther, GenderTypeASKU, GenderTypeUNK:
		return true
	default:
		return false
	}
}

// String converts the Gender type enum to a string
func (f GenderType) String() string {
	return string(f)
}

// UnmarshalGQL converts the supplied value to a Gender type.
func (f *GenderType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*f = GenderType(str)
	if !f.IsValid() {
		return fmt.Errorf("%s is not a valid ContactType type", str)
	}

	return nil
}

// MarshalGQL writes the gender type to the supplied writer
func (f GenderType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(f.String()))
}

// IsValid returns true if a match result is valid
func (m MatchResult) IsValid() bool {
	switch m {
	case MatchResultMatch, MatchResultPossibleMatch, MatchResultNoMatch:
		return true
	default:
		return false
	}
}

// String converts the match result enum to a string
func (m MatchResult) String() string {
	return string(m)
}
