package scv

import (
	"errors"
	"fmt"

	"github.com/iancoleman/strcase"
)

type StringCase int

const (
	Snake StringCase = iota
	Kebab
	UpperCamel
	LowerCamel
)

func Validate(value string, stringCase StringCase) error {
	converter := func() func(string) string {
		switch stringCase {
		case Snake:
			return strcase.ToSnake
		case Kebab:
			return strcase.ToKebab
		case UpperCamel:
			return strcase.ToCamel
		case LowerCamel:
			return strcase.ToLowerCamel
		default:
			return nil
		}
	}()

	if converter == nil {
		return errors.New("invalid string case")
	}

	expect := converter(value)
	if value != expect {
		return fmt.Errorf("expect string case is %s, actual is %s", expect, value)
	}

	return nil
}
