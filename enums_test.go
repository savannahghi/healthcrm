package healthcrm

import (
	"bytes"
	"strconv"
	"testing"
)

func TestIdentifierType_String(t *testing.T) {
	tests := []struct {
		name string
		e    IdentifierType
		want string
	}{
		{
			name: "happy case: enum to string",
			e:    IdentifierTypeNationalID,
			want: "NATIONAL_ID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("IdentifierType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdentifierType_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    IdentifierType
		want bool
	}{
		{
			name: "valid type",
			e:    IdentifierTypeNationalID,
			want: true,
		},
		{
			name: "invalid type",
			e:    IdentifierType("invalid"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("IdentifierType.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdentifierType_UnmarshalGQL(t *testing.T) {
	value := IdentifierTypeNationalID
	invalid := IdentifierType("invalid")

	type args struct {
		v interface{}
	}

	tests := []struct {
		name    string
		e       *IdentifierType
		args    args
		wantErr bool
	}{
		{
			name: "valid type",
			e:    &value,
			args: args{
				v: "NATIONAL_ID",
			},
			wantErr: false,
		},
		{
			name: "invalid type",
			e:    &invalid,
			args: args{
				v: "this is not a valid type",
			},
			wantErr: true,
		},
		{
			name: "non string type",
			e:    &invalid,
			args: args{
				v: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.UnmarshalGQL(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("IdentifierType.UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIdentifierType_MarshalGQL(t *testing.T) {
	w := &bytes.Buffer{}

	tests := []struct {
		name  string
		e     IdentifierType
		b     *bytes.Buffer
		wantW string
		panic bool
	}{
		{
			name:  "valid type enums",
			e:     IdentifierTypeNationalID,
			b:     w,
			wantW: strconv.Quote("NATIONAL_ID"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.MarshalGQL(tt.b)

			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("IdentifierType.MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestContactType_String(t *testing.T) {
	tests := []struct {
		name string
		e    ContactType
		want string
	}{
		{
			name: "happy case: enum to string",
			e:    ContactTypePhoneNumber,
			want: "PHONE_NUMBER",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("ContactType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContactType_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    ContactType
		want bool
	}{
		{
			name: "valid type",
			e:    ContactTypePhoneNumber,
			want: true,
		},
		{
			name: "invalid type",
			e:    ContactType("invalid"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("ContactType.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContactType_UnmarshalGQL(t *testing.T) {
	value := ContactTypePhoneNumber
	invalid := ContactType("invalid")

	type args struct {
		v interface{}
	}

	tests := []struct {
		name    string
		e       *ContactType
		args    args
		wantErr bool
	}{
		{
			name: "valid type",
			e:    &value,
			args: args{
				v: "PHONE_NUMBER",
			},
			wantErr: false,
		},
		{
			name: "invalid type",
			e:    &invalid,
			args: args{
				v: "this is not a valid type",
			},
			wantErr: true,
		},
		{
			name: "non string type",
			e:    &invalid,
			args: args{
				v: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.UnmarshalGQL(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ContactType.UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContactType_MarshalGQL(t *testing.T) {
	w := &bytes.Buffer{}

	tests := []struct {
		name  string
		e     ContactType
		b     *bytes.Buffer
		wantW string
		panic bool
	}{
		{
			name:  "valid type enums",
			e:     ContactTypePhoneNumber,
			b:     w,
			wantW: strconv.Quote("PHONE_NUMBER"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.MarshalGQL(tt.b)

			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("ContactType.MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestGenderType_String(t *testing.T) {
	tests := []struct {
		name string
		e    GenderType
		want string
	}{
		{
			name: "happy case: enum to string",
			e:    GenderTypeMale,
			want: "MALE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("GenderType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenderType_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    GenderType
		want bool
	}{
		{
			name: "valid type",
			e:    GenderTypeMale,
			want: true,
		},
		{
			name: "invalid type",
			e:    GenderType("invalid"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("GenderType.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenderType_UnmarshalGQL(t *testing.T) {
	value := GenderTypeFemale
	invalid := GenderType("invalid")

	type args struct {
		v interface{}
	}

	tests := []struct {
		name    string
		e       *GenderType
		args    args
		wantErr bool
	}{
		{
			name: "valid type",
			e:    &value,
			args: args{
				v: "FEMALE",
			},
			wantErr: false,
		},
		{
			name: "invalid type",
			e:    &invalid,
			args: args{
				v: "this is not a valid type",
			},
			wantErr: true,
		},
		{
			name: "non string type",
			e:    &invalid,
			args: args{
				v: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.UnmarshalGQL(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("GenderType.UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenderType_MarshalGQL(t *testing.T) {
	w := &bytes.Buffer{}

	tests := []struct {
		name  string
		e     GenderType
		b     *bytes.Buffer
		wantW string
		panic bool
	}{
		{
			name:  "valid type enums",
			e:     GenderTypeOther,
			b:     w,
			wantW: strconv.Quote("OTHER"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.MarshalGQL(tt.b)

			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("GenderType.MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
