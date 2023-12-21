package scv

import (
	"testing"
)

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

func TestValidateMapKeys(t *testing.T) {
	type args struct {
		value      map[string]string
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
				value:      map[string]string{"snale_case": "value"},
				stringCase: Snake,
			},
			false,
		},
		{
			"kebab case",
			args{
				value:      map[string]string{"kabel-case": "value"},
				stringCase: Kebab,
			},
			false,
		},
		{
			"upper camel case",
			args{
				value:      map[string]string{"UpperCamelCase": "value"},
				stringCase: UpperCamel,
			},
			false,
		},
		{
			"lower camel case",
			args{
				value:      map[string]string{"lowerCamel": "value"},
				stringCase: LowerCamel,
			},
			false,
		},
		{
			"invalid case",
			args{
				value:      map[string]string{"invalid_case": "value"},
				stringCase: 1_000_000,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateMapKeys(tt.args.value, tt.args.stringCase); (err != nil) != tt.wantErr {
				t.Errorf("ValidateMapKeys() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateMapValues(t *testing.T) {
	type args struct {
		value      map[string]any
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
				value:      map[string]any{"key": "snale_case"},
				stringCase: Snake,
			},
			false,
		},
		{
			"kebab case",
			args{
				value:      map[string]any{"key": "kabel-case"},
				stringCase: Kebab,
			},
			false,
		},
		{
			"upper camel case",
			args{
				value:      map[string]any{"key": "UpperCamelCase"},
				stringCase: UpperCamel,
			},
			false,
		},
		{
			"lower camel case",
			args{
				value:      map[string]any{"key": "lowerCamel"},
				stringCase: LowerCamel,
			},
			false,
		},
		{
			"skip not string",
			args{
				value:      map[string]any{"key": 1},
				stringCase: Snake,
			},
			false,
		},
		{
			"invalid case",
			args{
				value:      map[string]any{"key": "invalid_case"},
				stringCase: 1_000_000,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateMapValues(tt.args.value, tt.args.stringCase); (err != nil) != tt.wantErr {
				t.Errorf("ValidateMapValues() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
