package govin

import (
	"github.com/ynori7/go-vin/parse"
)

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
	Checksum     string // Empty if region does not support it (i.e. EU)
}

func NewVin(v string) (*Vin, error) {
	v = NormalizeVin(v)

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

	if vinRegion != parse.Region_Europe {
		// Checksums aren't used in Region_Europe (see ISO 3779 / ISO 3780)
		vinParts.Checksum = string(v[8])
		// TODO: validate the checksum
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
