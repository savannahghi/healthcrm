package healthcrm

import (
	"testing"

	"github.com/savannahghi/enumutils"
)

func TestConvertEnumutilsGenderToCRMGender(t *testing.T) {
	type args struct {
		gender enumutils.Gender
	}

	tests := []struct {
		name string
		args args
		want GenderType
	}{
		{
			name: "success: get male gender",
			args: args{
				gender: enumutils.Gender("male"),
			},
			want: GenderTypeMale,
		},
		{
			name: "success: get female gender",
			args: args{
				gender: enumutils.GenderFemale,
			},
			want: GenderTypeFemale,
		},
		{
			name: "success: get other gender",
			args: args{
				gender: enumutils.GenderBigender,
			},
			want: GenderTypeOther,
		},
		{
			name: "success: get ask but not unknown",
			args: args{
				gender: enumutils.GenderPreferNotToSay,
			},
			want: GenderTypeASKU,
		},
		{
			name: "success: get unknown gender",
			args: args{
				gender: enumutils.GenderUnknown,
			},
			want: GenderTypeUNK,
		},
		{
			name: "success: get unknown gender",
			args: args{
				gender: enumutils.Gender("invalide"),
			},
			want: GenderTypeUNK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertEnumutilsGenderToCRMGender(tt.args.gender); got != tt.want {
				t.Errorf("ConvertEnumutilsGenderToCRMGender() = %v, want %v", got, tt.want)
			}
		})
	}
}
