package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseYear(t *testing.T) {
	// given
	testcases := map[string]struct {
		vin             string
		year            int
		region          Region
		currentYearFunc func() int
	}{
		"German vin": {
			vin:             "WDBUH87J46X202412",
			year:            0,
			region:          Region_Europe,
			currentYearFunc: func() int { return 2020 },
		},
		"US vin": {
			vin:             "1J4GZB8S1TY103658",
			year:            1996,
			region:          Region_NorthAmerica,
			currentYearFunc: func() int { return 2020 },
		},
		"US vin in the future": {
			vin:             "1J4GZB8S1TY103658",
			year:            2026,
			region:          Region_NorthAmerica,
			currentYearFunc: func() int { return 2030 },
		},
		"US vin from next year": {
			vin:             "1J4GZB8S11Y103658",
			year:            2001,
			region:          Region_NorthAmerica,
			currentYearFunc: func() int { return 2020 },
		},
	}

	for testcase, testdata := range testcases {
		getCurrentYear = testdata.currentYearFunc

		// when
		year := ParseYear(testdata.vin, testdata.region)

		// then
		assert.Equal(t, testdata.year, year, testcase)
	}
}
