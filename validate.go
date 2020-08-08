package govin

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidVinLength  = errors.New("invalid vin length")
	ErrInvalidCharacters = errors.New("invalid characters")
)

var vinValidator = regexp.MustCompile("[A-HJ-NPR-Z0-9]{17}") //alphanumeric with no I, O, or Q

func ValidateVin(v string) error {
	v = NormalizeVin(v)

	if len(v) != 17 {
		return ErrInvalidVinLength
	}

	if !vinValidator.MatchString(v) {
		return ErrInvalidCharacters
	}

	return nil
}

func NormalizeVin(v string) string {
	return strings.ToUpper(strings.TrimSpace(v))
}
