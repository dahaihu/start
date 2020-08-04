package limit

import "time"

func newTokensFromDuration(d time.Duration, limit float64) float64 {
	// Split the integer and fractional parts ourself to minimize rounding errors.
	// See golang.org/issues/34861.
	sec := float64(d/time.Second) * limit
	nsec := float64(d%time.Second) * limit
	return sec + nsec/1e9
}

func oldTokensFromDuration(d time.Duration, limit float64) float64 {
	return d.Seconds() * limit
}