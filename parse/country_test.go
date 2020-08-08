package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseCountry(t *testing.T) {
	// given
	testcases := map[string]struct {
		vin string
		country string
	} {
		"german vin": {
			vin:    "WDBUH87J46X202412",
			country: "DE",
		},
		"US vin": {
			vin:    "1J4GZB8S1TY103658",
			country: "US",
		},
		"invalid vin": {
			vin:    "00BUH87J46I202412",
			country: "",
		},
	}

	for testcase, testdata := range testcases {
		// when
		country := ParseCountry(testdata.vin)

		// then
		assert.Equal(t, testdata.country, country, testcase)
	}
}
