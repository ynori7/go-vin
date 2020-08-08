package parse

import "regexp"

type Region string

const (
	Region_Africa       Region = "AF"
	Region_Asia                = "AS"
	Region_Europe              = "EU"
	Region_NorthAmerica        = "NA"
	Region_Oceania             = "OC"
	Region_SouthAmerica        = "SA"
	Region_Antarctica          = "AN"
	Region_Unknown             = "UN"
)

func ParseRegion(vin string) Region {
	for region, matcher := range regionMatchers {
		if matcher.MatchString(vin) {
			return region
		}
	}
	return Region_Unknown
}

var regionMatchers = map[Region]*regexp.Regexp{
	Region_Africa:       regexp.MustCompile("^[A-H]"),
	Region_Asia:         regexp.MustCompile("^[J-R]"),
	Region_Europe:       regexp.MustCompile("^[S-Z]"),
	Region_NorthAmerica: regexp.MustCompile("^[1-5]"),
	Region_Oceania:      regexp.MustCompile("^[6-7]"),
	Region_SouthAmerica: regexp.MustCompile("^[8-9]"),
}
