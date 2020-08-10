package govin

import (
	"errors"
	"regexp"
	"strings"
)

var (
	// ErrInvalidVinLength is returned if a supplied VIN is too long or too short
	ErrInvalidVinLength  = errors.New("invalid vin length")
	// ErrInvalidCharacters is returned if a VIN contains special characters or certain letters which are never contained in a VIN
	ErrInvalidCharacters = errors.New("invalid characters")
)

var vinValidator = regexp.MustCompile("[A-HJ-NPR-Z0-9]{17}") //alphanumeric with no I, O, or Q

// ValidateVin accepts a VIN string and returns nil if it is valid or an error explaining why it is invalid
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

// NormalizeVin normalizes a VIN string so that it can be parsed
func NormalizeVin(v string) string {
	return strings.ToUpper(strings.TrimSpace(v))
}
