package lodash

import (
	"strconv"
	"strings"
)

// ParseInt parses s as an integer in the given base and returns the result as
// an int64. Surrounding whitespace is ignored. A base of 0 infers the base from
// any "0x", "0o" or "0b" prefix (defaulting to decimal). It mirrors lodash's
// parseInt, delegating to strconv.ParseInt for the conversion.
func ParseInt(s string, base int) (int64, error) {
	return strconv.ParseInt(strings.TrimSpace(s), base, 64)
}

// ParseFloat parses s as a floating-point number and returns the result as a
// float64. Surrounding whitespace is ignored. It complements ParseInt for
// fractional input, delegating to strconv.ParseFloat.
func ParseFloat(s string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}
