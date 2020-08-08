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
	} {
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
