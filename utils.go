package healthcrm

import "github.com/savannahghi/enumutils"

// ConvertEnumutilsGenderToCRMGender converts an enumutils Gender to a CRM gender type
func ConvertEnumutilsGenderToCRMGender(gender enumutils.Gender) GenderType {
	switch gender {
	case enumutils.GenderMale:
		return GenderTypeMale

	case enumutils.GenderFemale:
		return GenderTypeFemale

	case enumutils.GenderAgender, enumutils.GenderBigender, enumutils.GenderGenderQueer,
		enumutils.GenderNonBinary, enumutils.GenderTransGender, enumutils.GenderTwoSpirit, enumutils.GenderOther:
		return GenderTypeOther

	case enumutils.GenderPreferNotToSay:
		return GenderTypeASKU

	case enumutils.GenderUnknown:
		return GenderTypeUNK
	default:
		return GenderTypeUNK
	}
}
