package govin

import (
	"errors"
	"strings"

	"github.com/ynori7/go-vin/parse"
)

var ErrInvalidVinLength = errors.New("invalid vin length")

type Vin struct {
	Vin       string
	Region    parse.Region // The content derived from vin parts
	Country   string       // The country derived from vin parts
	ModelYear string       // The actual year, derived from the year field in the vin parts
	VinParts  VinParts
}

type VinParts struct {
	Wmi          string // World Manufacturer Identifier
	Vds          string // Vehicle Descriptor Section
	Vis          string // Vehicle Identifier Section
	ModelYear    string
	SerialNumber string
}

func NewVin(v string) (*Vin, error) {
	v = strings.ToUpper(v) //normalize it

	if err := ValidateVin(v); err != nil {
		return nil, err
	}

	vinRegion := parse.ParseRegion(v)

	vinParts := VinParts{
		Wmi:          v[0:3],
		Vds:          v[3:9],
		Vis:          v[9:],
		SerialNumber: v[12:],
		ModelYear:    string(v[9]),
	}

	vin := &Vin{
		Vin:       v,
		ModelYear: parse.ParseYear(v, vinRegion),
		Region:    vinRegion,
		Country:   parse.ParseCountry(v),
		VinParts:  vinParts,
	}

	return vin, nil
}

func ValidateVin(v string) error {
	if len(v) != 17 {
		return ErrInvalidVinLength
	}

	//TODO verify checksum

	return nil
}
