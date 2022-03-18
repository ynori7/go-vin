package govin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidateVin(t *testing.T) {
	// given
	testcases := map[string]struct {
		vin string
		err error
	}{
		"valid": {
			vin: "WDBUH87J46X202412",
			err: nil,
		},
		"valid after normalizing": {
			vin: " WDbuH87J46X202412  ",
			err: nil,
		},
		"with an I in it": {
			vin: "WDBUH87J46I202412",
			err: ErrInvalidCharacters,
		},
		"with special chars": {
			vin: "WDBUH87J46#202412",
			err: ErrInvalidCharacters,
		},
		"too long": {
			vin: "WDBUH87J46X202412333",
			err: ErrInvalidVinLength,
		},
		"too short": {
			vin: "WDBUH87J46X2024",
			err: ErrInvalidVinLength,
		},
	}

	for testcase, testdata := range testcases {
		// when
		err := ValidateVin(testdata.vin)

		// then
		assert.Equal(t, err, testdata.err, testcase)
	}
}

func Test_HasValidChecksum(t *testing.T) {
	// given
	testcases := map[string]struct {
		vin      string
		expected bool
	}{
		"Valid German VIN": {
			vin:      "WDBUH87J46X202412",
			expected: true,
		},
		"Valid US VIN": {
			vin:      "1J4GZB8S1TY103658",
			expected: true,
		},
		"Invalid US VIN": {
			vin:      "1J4GZB8S2TY103658", //2 instead of 1
			expected: false,
		},
		"Invalid length": {
			vin:      "1J4GZB8S2T",
			expected: false,
		},
	}

	for testcase, testdata := range testcases {
		// when
		actual := HasValidChecksum(testdata.vin)

		// then
		assert.Equal(t, testdata.expected, actual, testcase)
	}
}
