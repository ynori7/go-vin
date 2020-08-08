package govin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ynori7/go-vin/parse"
)

func Test_NewVin(t *testing.T) {
	// given
	testcases := map[string]struct {
		vin          string
		expectedResp *Vin
		expectedErr  error
	}{
		"Valid German VIN": {
			vin:         "WDBUH87J46X202412",
			expectedErr: nil,
			expectedResp: &Vin{
				Vin:       "WDBUH87J46X202412",
				Region:    parse.Region_Europe,
				Country:   "DE",
				ModelYear: "",
			},
		},
		"Valid US VIN": {
			vin:         "1J4GZB8S1TY103658",
			expectedErr: nil,
			expectedResp: &Vin{
				Vin:       "1J4GZB8S1TY103658",
				Region:    parse.Region_NorthAmerica,
				Country:   "US",
				ModelYear: "1996",
			},
		},
		"Invalid VIN": {
			vin:          "1J4GZB8S1TY103",
			expectedErr:  ErrInvalidVinLength,
			expectedResp: nil,
		},
	}

	for testcase, testdata := range testcases {
		// when
		parsedVin, err := NewVin(testdata.vin)
		if testdata.expectedErr != nil {
			require.Error(t, err, testcase)
			assert.EqualError(t, err, testdata.expectedErr.Error(), testcase)
			continue
		}

		// then
		require.NoError(t, err, testcase)
		assert.Equal(t, testdata.expectedResp.Vin, parsedVin.Vin)
		assert.Equal(t, testdata.expectedResp.ModelYear, parsedVin.ModelYear)
		assert.Equal(t, testdata.expectedResp.Region, parsedVin.Region)
		assert.Equal(t, testdata.expectedResp.Country, parsedVin.Country)

	}
}
