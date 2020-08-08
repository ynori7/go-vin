package parse

import "regexp"

type Region string

const (
	Africa       Region = "AF"
	Asia                = "AS"
	Europe              = "EU"
	NorthAmerica        = "NA"
	Oceania             = "OC"
	SouthAmerica        = "SA"
	Antarctica          = "AN"
	Unknown             = "UN"
)

func ParseRegion(vin string) Region {
	for region, matcher := range regionMatchers {
		if matcher.MatchString(vin) {
			return region
		}
	}
	return Unknown
}

var regionMatchers = map[Region]*regexp.Regexp{
	Africa:       regexp.MustCompile("^[A-H].*"),
	Asia:         regexp.MustCompile("^[J-R].*"),
	Europe:       regexp.MustCompile("^[S-Z].*"),
	NorthAmerica: regexp.MustCompile("^[1-5].*"),
	Oceania:      regexp.MustCompile("^6-7].*"),
	SouthAmerica: regexp.MustCompile("^[8-9].*"),
}
