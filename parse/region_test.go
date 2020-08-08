package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseRegion(t *testing.T) {
	// given
	testcases := map[string]struct {
		vin string
		region Region
	} {
		"german vin": {
			vin:    "WDBUH87J46X202412",
			region: Region_Europe,
		},
		"US vin": {
			vin:    "1J4GZB8S1TY103658",
			region: Region_NorthAmerica,
		},
		"invalid vin": {
			vin:    "00BUH87J46I202412",
			region: Region_Unknown,
		},
	}

	for testcase, testdata := range testcases {
		// when
		region := ParseRegion(testdata.vin)

		// then
		assert.Equal(t, testdata.region, region, testcase)
	}
}