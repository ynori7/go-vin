package parse

import (
	"fmt"
	"time"
)

func ParseYear(vin string, region Region) string {
	mapping, ok := yearMapping[string(vin[9])]
	if !ok {
		return ""
	}

	possibleYears, ok := mapping[region]
	if !ok {
		return ""
	}

	// take the newest year which is not in the future
	currentYear := getCurrentYear()
	for i := len(possibleYears) - 1; i > 0; i-- {
		if possibleYears[i] <= currentYear {
			return possibleYears[i]
		}
	}
	return possibleYears[0]
}

var getCurrentYear = func() string {
	return fmt.Sprintf("%d", time.Now().Year())
}

// map of code to possible years
var yearMapping = map[string]map[Region][]string{
	"A": {
		Region_NorthAmerica: {"1980", "2010"},
	},
	"B": {
		Region_NorthAmerica: {"1981", "2011"},
	},
	"C": {
		Region_NorthAmerica: {"1982", "2012"},
	},
	"D": {
		Region_NorthAmerica: {"1983", "2013"},
	},
	"E": {
		Region_NorthAmerica: {"1984", "2014"},
	},
	"F": {
		Region_NorthAmerica: {"1985", "2015"},
	},
	"G": {
		Region_NorthAmerica: {"1986", "2016"},
	},
	"H": {
		Region_NorthAmerica: {"1987", "2017"},
	},
	"I": {
		Region_NorthAmerica: {}, // not used
	},
	"J": {
		Region_NorthAmerica: {"1988", "2018"},
	},
	"K": {
		Region_NorthAmerica: {"1989", "2019"},
	},
	"L": {
		Region_NorthAmerica: {"1990", "2020"},
	},
	"M": {
		Region_NorthAmerica: {"1991", "2021"},
	},
	"N": {
		Region_NorthAmerica: {"1992", "2022"},
	},
	"O": {
		Region_NorthAmerica: {}, // not used
	},
	"P": {
		Region_NorthAmerica: {"1993", "2023"},
	},
	"Q": {
		Region_NorthAmerica: {}, // not used
	},
	"R": {
		Region_NorthAmerica: {"1994", "2024"},
	},
	"S": {
		Region_NorthAmerica: {"1995", "2025"},
	},
	"T": {
		Region_NorthAmerica: {"1996", "2026"},
	},
	"U": {
		Region_NorthAmerica: {}, // not used
	},
	"V": {
		Region_NorthAmerica: {"1997", "2027"},
	},
	"W": {
		Region_NorthAmerica: {"1998", "2028"},
	},
	"X": {
		Region_NorthAmerica: {"1999", "2029"},
	},
	"Y": {
		Region_NorthAmerica: {"2000", "2030"},
	},
	"Z": {
		Region_NorthAmerica: {}, // not used
	},
	"0": {
		Region_NorthAmerica: {}, // not used
	},
	"1": {
		Region_NorthAmerica: {"2001", "2031"},
	},
	"2": {
		Region_NorthAmerica: {"2002", "2032"},
	},
	"3": {
		Region_NorthAmerica: {"2003", "2033"},
	},
	"4": {
		Region_NorthAmerica: {"2004", "2034"},
	},
	"5": {
		Region_NorthAmerica: {"2005", "2035"},
	},
	"6": {
		Region_NorthAmerica: {"2006", "2036"},
	},
	"7": {
		Region_NorthAmerica: {"2007", "2037"},
	},
	"8": {
		Region_NorthAmerica: {"2008", "2038"},
	},
	"9": {
		Region_NorthAmerica: {"2009", "2039"},
	},
}
