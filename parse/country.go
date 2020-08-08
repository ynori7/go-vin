package parse

import "regexp"

func ParseCountry(vin string) string {
	for c, m := range countryMatchers {
		if m.MatchString(vin) {
			return c
		}
	}
	return ""
}

var countryMatchers = map[string]*regexp.Regexp {
	"ZA": regexp.MustCompile("^A[A-H]"),
	"CI": regexp.MustCompile("^A[J-N]"),
	"AO": regexp.MustCompile("^B[A-E]"),
	"KE": regexp.MustCompile("^B[F-K]"),
	"TZ": regexp.MustCompile("^B[L-R]"),
	"BJ": regexp.MustCompile("^C[A-E]"),
	"MG": regexp.MustCompile("^C[F-K]"),
	"TN": regexp.MustCompile("^C[L-R]"),
	"EG": regexp.MustCompile("^D[A-E]"),
	"MA": regexp.MustCompile("^D[F-K]"),
	"ZM": regexp.MustCompile("^D[L-R]"),
	"ET": regexp.MustCompile("^E[A-E]"),
	"NZ": regexp.MustCompile("^(E[F-K]|7[A-E])"),
	"GH": regexp.MustCompile("^F[A-E]"),
	"NG": regexp.MustCompile("^F[F-K]"),
	"JP": regexp.MustCompile("^J[A-Z0]"),
	"LK": regexp.MustCompile("^K[A-E]"),
	"IL": regexp.MustCompile("^K[F-K]"),
	"KR": regexp.MustCompile("^K[L-R]"),
	"KZ": regexp.MustCompile("^K[S-Z0]"),
	"CN": regexp.MustCompile("^L[A-Z0]"),
	"IN": regexp.MustCompile("^M[A-E]"),
	"ID": regexp.MustCompile("^M[F-K]"),
	"TH": regexp.MustCompile("^M[L-R]"),
	"IR": regexp.MustCompile("^N[A-E]"),
	"PK": regexp.MustCompile("^N[F-K]"),
	"TR": regexp.MustCompile("^N[L-R]"),
	"PH": regexp.MustCompile("^P[A-E]"),
	"SG": regexp.MustCompile("^P[F-K]"),
	"MY": regexp.MustCompile("^P[L-R]"),
	"AE": regexp.MustCompile("^R[A-E]"),
	"TW": regexp.MustCompile("^R[F-K]"),
	"VN": regexp.MustCompile("^R[L-R]"),
	"SA": regexp.MustCompile("^R[S-Z0]"),
	"GB": regexp.MustCompile("^S[A-M]"),
	"DE": regexp.MustCompile("^(S[N-T])|(W[A-Z0])"), // East & West Germany
	"PL": regexp.MustCompile("^S[U-Z]"),
	"LV": regexp.MustCompile("^S[1-4]"),
	"CH": regexp.MustCompile("^T[A-H]"),
	"CZ": regexp.MustCompile("^T[J-P]"),
	"HU": regexp.MustCompile("^T[R-V]"),
	"PT": regexp.MustCompile("^T[W-Z0-1]"),
	"DK": regexp.MustCompile("^U[H-M]"),
	"IE": regexp.MustCompile("^U[N-T]"),
	"RO": regexp.MustCompile("^U[U-Z]"),
	"SK": regexp.MustCompile("^U[5-7]"),
	"AT": regexp.MustCompile("^V[A-E]"),
	"FR": regexp.MustCompile("^V[F-R]"),
	"ES": regexp.MustCompile("^V[S-W]"),
	"RS": regexp.MustCompile("^V[X-Z0-2]"),
	"HR": regexp.MustCompile("^V[3-5]"),
	"EE": regexp.MustCompile("^V[6-Z0]"),
	"BG": regexp.MustCompile("^X[A-E]"),
	"GR": regexp.MustCompile("^X[F-K]"),
	"NL": regexp.MustCompile("^X[L-R]"),
	"RU": regexp.MustCompile("^(X[S-W])|(X[3-Z0])"), // Russia & USSR
	"LU": regexp.MustCompile("^X[X-Z0-2]"),
	"BE": regexp.MustCompile("^Y[A-E]"),
	"FI": regexp.MustCompile("^Y[F-K]"),
	"MT": regexp.MustCompile("^Y[L-R]"),
	"SE": regexp.MustCompile("^Y[S-W]"),
	"NO": regexp.MustCompile("^Y[X-Z0-2]"),
	"BY": regexp.MustCompile("^Y[3-5]"),
	"UA": regexp.MustCompile("^Y[6-Z0]"),
	"IT": regexp.MustCompile("^Z[A-R]"),
	"SI": regexp.MustCompile("^Z[X-Z0-2]"),
	"LT": regexp.MustCompile("^Z[3-5]"),
	"US": regexp.MustCompile("^(1[A-Z0])|(4[A-Z0])|(5[A-Z0])"),
	"CA": regexp.MustCompile("^2[A-Z0]"),
	"MX": regexp.MustCompile("^3[A-Z0-7]"),
	"KY": regexp.MustCompile("^3[8-Z0]"),
	"AU": regexp.MustCompile("^6[A-W]"),
	"AR": regexp.MustCompile("^8[A-E]"),
	"CL": regexp.MustCompile("^8[F-K]"),
	"EC": regexp.MustCompile("^8[L-R]"),
	"PE": regexp.MustCompile("^8[S-W]"),
	"VE": regexp.MustCompile("^8[X-Z0-2]"),
	"BR": regexp.MustCompile("^9[A-E,3-9]"),
	"CO": regexp.MustCompile("^9[F-K]"),
	"PY": regexp.MustCompile("^9[L-R]"),
	"UY": regexp.MustCompile("^9[S-W]"),
	"TT": regexp.MustCompile("^9[X-Z0-2]"),
}
