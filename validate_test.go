package scv

import "testing"

func TestValidate(t *testing.T) {
	type args struct {
		value      string
		stringCase StringCase
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"snake case",
			args{
				value:      "snale_case",
				stringCase: Snake,
			},
			false,
		},
		{
			"kebab case",
			args{
				value:      "kabel-case",
				stringCase: Kebab,
			},
			false,
		},
		{
			"upper camel case",
			args{
				value:      "UpperCamelCase",
				stringCase: UpperCamel,
			},
			false,
		},
		{
			"lower camel case",
			args{
				value:      "lowerCamel",
				stringCase: LowerCamel,
			},
			false,
		},
		{
			"invalid case",
			args{
				value:      "invalid_case",
				stringCase: 1_000_000,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.args.value, tt.args.stringCase); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
