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
		currentYearFunc func() int
	}{
		"German vin": {
			vin:             "WDBUH87J46X202412",
			year:            2006,
			currentYearFunc: func() int { return 2020 },
		},
		"US vin": {
			vin:             "1J4GZB8S1TY103658",
			year:            1996,
			currentYearFunc: func() int { return 2020 },
		},
		"US vin in the future": {
			vin:             "1J4GZB8S1TY103658",
			year:            2026,
			currentYearFunc: func() int { return 2030 },
		},
		"US vin from next year": {
			vin:             "1J4GZB8S11Y103658",
			year:            2001,
			currentYearFunc: func() int { return 2020 },
		},
	}

	for testcase, testdata := range testcases {
		getCurrentYear = testdata.currentYearFunc

		// when
		year := ParseYear(testdata.vin)

		// then
		assert.Equal(t, testdata.year, year, testcase)
	}
}
