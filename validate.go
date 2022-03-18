package govin

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	// ErrInvalidVinLength is returned if a supplied VIN is too long or too short
	ErrInvalidVinLength = errors.New("invalid vin length")
	// ErrInvalidCharacters is returned if a VIN contains special characters or certain letters which are never contained in a VIN
	ErrInvalidCharacters = errors.New("invalid characters")
)

// Transliteration involves removing all of letters and substituting them with their appropriate numerical counterparts
var checksumTransliterations = map[string]int{
	"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8,
	"J": 1, "K": 2, "L": 3, "M": 4, "N": 5, "P": 7, "R": 9,
	"S": 2, "T": 3, "U": 4, "V": 5, "W": 6, "X": 7, "Y": 8, "Z": 9,
}

// The weight factor for each position in the VIN. The 9th position is that of the check digit, so the multiplier is zero
var checksumWeights = []int{8, 7, 6, 5, 4, 3, 2, 10, 0, 9, 8, 7, 6, 5, 4, 3, 2}

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

// HasValidChecksum returns true if the vin contains the correct value in its checksum digit
func HasValidChecksum(vin string) bool {
	if len(vin) != 17 {
		return false
	}
	checkDigit := string(vin[8])

	sum := 0
	for k, v := range vin {
		i, err := strconv.Atoi(string(v))
		if err != nil {
			i = checksumTransliterations[string(v)]
		}
		sum += i * checksumWeights[k]
	}

	expected := sum % 11
	if expected == 10 {
		return checkDigit == "X"
	}

	return checkDigit == strconv.Itoa(expected)
}
