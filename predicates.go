package lodash

import (
	"reflect"
	"regexp"
	"time"
)

// maxSafeInt is the largest integer that a float64 can represent exactly,
// matching JavaScript's Number.MAX_SAFE_INTEGER. It bounds the "safe integer"
// and "length" predicates below.
const maxSafeInt = 9007199254740991

// IsFunction reports whether value is a function.
func IsFunction(value any) bool {
	if value == nil {
		return false
	}
	return reflect.ValueOf(value).Kind() == reflect.Func
}

// IsRegExp reports whether value is a compiled *regexp.Regexp.
func IsRegExp(value any) bool {
	_, ok := value.(*regexp.Regexp)
	return ok
}

// IsDate reports whether value is a time.Time (or a non-nil *time.Time),
// the closest Go analogue of a JavaScript Date.
func IsDate(value any) bool {
	switch v := value.(type) {
	case time.Time:
		return true
	case *time.Time:
		return v != nil
	default:
		return false
	}
}

// IsBuffer reports whether value is a byte slice ([]byte), the closest Go
// analogue of a Node.js Buffer.
func IsBuffer(value any) bool {
	_, ok := value.([]byte)
	return ok
}

// IsSafeInteger reports whether value is an integer-valued number that can be
// represented exactly as a float64, i.e. within the inclusive range
// [-(2^53-1), 2^53-1]. It mirrors lodash's isSafeInteger.
func IsSafeInteger(value any) bool {
	if !IsInteger(value) {
		return false
	}
	n := ToNumber(value)
	return n >= -maxSafeInt && n <= maxSafeInt
}

// IsLength reports whether value is a valid array-like length: an integer-valued
// number that is non-negative and no greater than 2^53-1.
func IsLength(value any) bool {
	if !IsInteger(value) {
		return false
	}
	n := ToNumber(value)
	return n >= 0 && n <= maxSafeInt
}

// IsNull reports whether value is nil, either as an untyped nil interface or as
// a typed nil pointer, map, slice, channel or function. It is the Go analogue
// of a JavaScript null check.
func IsNull(value any) bool {
	if value == nil {
		return true
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan,
		reflect.Func, reflect.Interface:
		return rv.IsNil()
	default:
		return false
	}
}

// IsUndefined reports whether value is an untyped nil interface. Unlike IsNull
// it does not consider typed nil values (such as a nil *T) to be undefined,
// mirroring the distinction JavaScript draws between undefined and null.
func IsUndefined(value any) bool {
	return value == nil
}

// IsObject reports whether value is a non-primitive, non-nil value: a map,
// slice, array, struct, pointer, channel or function. Strings, booleans and
// numbers return false. It mirrors lodash's isObject.
func IsObject(value any) bool {
	if value == nil {
		return false
	}
	switch reflect.ValueOf(value).Kind() {
	case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct,
		reflect.Ptr, reflect.Chan, reflect.Func:
		return true
	default:
		return false
	}
}

// IsArrayLike reports whether value has a length: a slice, array or string.
// Functions and maps are not array-like. It mirrors lodash's isArrayLike.
func IsArrayLike(value any) bool {
	if value == nil {
		return false
	}
	switch reflect.ValueOf(value).Kind() {
	case reflect.Slice, reflect.Array, reflect.String:
		return true
	default:
		return false
	}
}

// IsArrayLikeObject reports whether value is array-like and also an object,
// i.e. a slice or array but not a string. It mirrors lodash's
// isArrayLikeObject.
func IsArrayLikeObject(value any) bool {
	if value == nil {
		return false
	}
	switch reflect.ValueOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		return true
	default:
		return false
	}
}

// ToLength converts value to a valid array-like length by truncating it to an
// integer and clamping the result to the inclusive range [0, 2^53-1]. It
// mirrors lodash's toLength.
func ToLength(value any) int {
	n := ToInteger(value)
	if n < 0 {
		return 0
	}
	if n > maxSafeInt {
		return maxSafeInt
	}
	return n
}

// IsEqualWith performs a deep equality comparison like IsEqual, but first
// consults customizer. When customizer's second return value (handled) is true,
// its first return value is used as the comparison result; otherwise the
// default deep comparison is applied.
func IsEqualWith(a, b any, customizer func(a, b any) (equal bool, handled bool)) bool {
	if customizer != nil {
		if equal, handled := customizer(a, b); handled {
			return equal
		}
	}
	return IsEqual(a, b)
}

// IsMatchWith reports whether object matches the shape of source using a custom
// comparator for each source key. Every key in source must be present in object
// and satisfy customizer(objectValue, sourceValue). An empty source always
// matches.
func IsMatchWith(object, source map[string]any, customizer func(objValue, srcValue any) bool) bool {
	for k, sv := range source {
		ov, ok := object[k]
		if !ok {
			return false
		}
		if customizer != nil {
			if !customizer(ov, sv) {
				return false
			}
		} else if !IsEqual(ov, sv) {
			return false
		}
	}
	return true
}
