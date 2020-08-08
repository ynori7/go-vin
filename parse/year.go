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
	currentYear := fmt.Sprintf("%d", time.Now().Year())
	for i := len(possibleYears) - 1; i > 0; i-- {
		if possibleYears[i] <= currentYear {
			return possibleYears[i]
		}
	}
	return possibleYears[0]
}

// map of code to possible years
var yearMapping = map[string]map[Region][]string {
	"A": {
		NorthAmerica: {"1980", "2010"},
	},
	"B": {
		NorthAmerica: {"1981", "2011"},
	},
	"C": {
		NorthAmerica: {"1982", "2012"},
	},
	"D": {
		NorthAmerica: {"1983", "2013"},
	},
	"E": {
		NorthAmerica: {"1984", "2014"},
	},
	"F": {
		NorthAmerica: {"1985", "2015"},
	},
	"G": {
		NorthAmerica: {"1986", "2016"},
	},
	"H": {
		NorthAmerica: {"1987", "2017"},
	},
	"I": {
		NorthAmerica: {}, // not used
	},
	"J": {
		NorthAmerica: {"1988", "2018"},
	},
	"K": {
		NorthAmerica: {"1989", "2019"},
	},
	"L": {
		NorthAmerica: {"1990", "2020"},
	},
	"M": {
		NorthAmerica: {"1991", "2021"},
	},
	"N": {
		NorthAmerica: {"1992", "2022"},
	},
	"O": {
		NorthAmerica: {}, // not used
	},
	"P": {
		NorthAmerica: {"1993", "2023"},
	},
	"Q": {
		NorthAmerica: {}, // not used
	},
	"R": {
		NorthAmerica: {"1994", "2024"},
	},
	"S": {
		NorthAmerica: {"1995", "2025"},
	},
	"T": {
		NorthAmerica: {"1996", "2026"},
	},
	"U": {
		NorthAmerica: {}, // not used
	},
	"V": {
		NorthAmerica: {"1997", "2027"},
	},
	"W": {
		NorthAmerica: {"1998", "2028"},
	},
	"X": {
		NorthAmerica: {"1999", "2029"},
	},
	"Y": {
		NorthAmerica: {"2000", "2030"},
	},
	"Z": {
		NorthAmerica: {}, // not used
	},
	"0": {
		NorthAmerica: {}, // not used
	},
	"1": {
		NorthAmerica: {"2001", "2031"},
	},
	"2": {
		NorthAmerica: {"2002", "2032"},
	},
	"3": {
		NorthAmerica: {"2003", "2033"},
	},
	"4": {
		NorthAmerica: {"2004", "2034"},
	},
	"5": {
		NorthAmerica: {"2005", "2035"},
	},
	"6": {
		NorthAmerica: {"2006", "2036"},
	},
	"7": {
		NorthAmerica: {"2007", "2037"},
	},
	"8": {
		NorthAmerica: {"2008", "2038"},
	},
	"9": {
		NorthAmerica: {"2009", "2039"},
	},
}
