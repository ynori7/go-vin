package govin

import (
	"github.com/ynori7/go-vin/parse"
)

// Vin contains human-readable data derived from the VIN as well as the raw vin parts
type Vin struct {
	Vin       string
	Region    parse.Region // The content derived from vin parts
	Country   string       // The country derived from vin parts
	ModelYear int          // The actual year, derived from the year field in the vin parts. Will be zero if it cannot be parsed reliably
	VinParts  VinParts
}

// VinParts contains the raw pieces of the VIN. Certain sequences within the ID represent certain things.
type VinParts struct {
	Wmi           string // World Manufacturer Identifier
	Vds           string // Vehicle Descriptor Section
	Vis           string // Vehicle Identifier Section
	ModelYear     string
	SerialNumber  string
	Checksum      string // Empty if region does not support it (i.e. GB)
	ChecksumValid bool   // True if the checksum is valid
}

// ParseVin accepts a VIN string and then validates it and returns structured data derived from the VIN or an error if it is invalid
func ParseVin(v string) (*Vin, error) {
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

	country := parse.ParseCountry(v)

	// Checksums aren't mandatory in Region_Europe (see ISO 3779 / ISO 3780), but are typically used still except in Great Britain
	if country != "GB" {
		vinParts.Checksum = string(v[8])
		vinParts.ChecksumValid = HasValidChecksum(v)
	}

	vin := &Vin{
		Vin:       v,
		ModelYear: parse.ParseYear(v),
		Region:    vinRegion,
		Country:   country,
		VinParts:  vinParts,
	}

	return vin, nil
}
