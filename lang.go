package lodash

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// Clone returns a shallow copy of value. For maps and slices a new top-level
// container is allocated whose elements are copied by assignment; nested
// containers are shared with the original. For all other types the value is
// returned unchanged (Go value semantics already copy it).
func Clone[T any](value T) T {
	rv := reflect.ValueOf(&value).Elem()
	switch rv.Kind() {
	case reflect.Map, reflect.Slice:
		out := reflect.New(rv.Type()).Elem()
		out.Set(shallowCopy(rv))
		return out.Interface().(T)
	default:
		return value
	}
}

func shallowCopy(rv reflect.Value) reflect.Value {
	switch rv.Kind() {
	case reflect.Map:
		if rv.IsNil() {
			return rv
		}
		out := reflect.MakeMapWithSize(rv.Type(), rv.Len())
		for _, k := range rv.MapKeys() {
			out.SetMapIndex(k, rv.MapIndex(k))
		}
		return out
	case reflect.Slice:
		if rv.IsNil() {
			return rv
		}
		out := reflect.MakeSlice(rv.Type(), rv.Len(), rv.Len())
		reflect.Copy(out, rv)
		return out
	default:
		return rv
	}
}

// CloneDeep returns a deep copy of value. Maps, slices, arrays, pointers and
// structs are recursively cloned so the result shares no mutable state with the
// original. Channels, functions and unsafe pointers are copied by reference.
func CloneDeep[T any](value T) T {
	rv := reflect.ValueOf(&value).Elem()
	out := reflect.New(rv.Type()).Elem()
	out.Set(cloneValue(rv))
	return out.Interface().(T)
}

func cloneValue(rv reflect.Value) reflect.Value {
	switch rv.Kind() {
	case reflect.Map:
		if rv.IsNil() {
			return rv
		}
		out := reflect.MakeMapWithSize(rv.Type(), rv.Len())
		for _, k := range rv.MapKeys() {
			out.SetMapIndex(cloneValue(k), cloneValue(rv.MapIndex(k)))
		}
		return out
	case reflect.Slice:
		if rv.IsNil() {
			return rv
		}
		out := reflect.MakeSlice(rv.Type(), rv.Len(), rv.Len())
		for i := 0; i < rv.Len(); i++ {
			out.Index(i).Set(cloneValue(rv.Index(i)))
		}
		return out
	case reflect.Array:
		out := reflect.New(rv.Type()).Elem()
		for i := 0; i < rv.Len(); i++ {
			out.Index(i).Set(cloneValue(rv.Index(i)))
		}
		return out
	case reflect.Ptr:
		if rv.IsNil() {
			return rv
		}
		out := reflect.New(rv.Type().Elem())
		out.Elem().Set(cloneValue(rv.Elem()))
		return out
	case reflect.Struct:
		out := reflect.New(rv.Type()).Elem()
		for i := 0; i < rv.NumField(); i++ {
			if out.Field(i).CanSet() {
				out.Field(i).Set(cloneValue(rv.Field(i)))
			}
		}
		return out
	case reflect.Interface:
		if rv.IsNil() {
			return rv
		}
		out := reflect.New(rv.Type()).Elem()
		out.Set(cloneValue(rv.Elem()))
		return out
	default:
		return rv
	}
}

// IsEqual reports whether a and b are deeply equal, recursively comparing the
// contents of maps, slices, arrays, structs and pointers. It is a thin,
// well-defined wrapper over reflect.DeepEqual.
func IsEqual(a, b any) bool {
	return reflect.DeepEqual(a, b)
}

// Eq reports whether two comparable values are equal. It is the typed
// counterpart to IsEqual for values whose type is known at compile time.
func Eq[T comparable](a, b T) bool {
	return a == b
}

// IsEmpty reports whether value is empty. Nil, empty strings, and empty maps,
// slices, arrays and channels are considered empty. Following lodash semantics,
// numbers, booleans and other scalar values are always reported as empty.
func IsEmpty(value any) bool {
	if value == nil {
		return true
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.String, reflect.Map, reflect.Slice, reflect.Array, reflect.Chan:
		return rv.Len() == 0
	case reflect.Ptr, reflect.Interface:
		if rv.IsNil() {
			return true
		}
		return IsEmpty(rv.Elem().Interface())
	case reflect.Struct:
		return rv.NumField() == 0
	default:
		return true
	}
}

// IsNil reports whether value is nil, including typed nil pointers, maps,
// slices, channels, functions and interfaces.
func IsNil(value any) bool {
	if value == nil {
		return true
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func, reflect.Interface:
		return rv.IsNil()
	default:
		return false
	}
}

// IsPlainObject reports whether value is a plain map keyed by strings, i.e. the
// Go analogue of a JavaScript plain object such as map[string]any.
func IsPlainObject(value any) bool {
	if value == nil {
		return false
	}
	rv := reflect.ValueOf(value)
	return rv.Kind() == reflect.Map && rv.Type().Key().Kind() == reflect.String
}

// IsMatch reports whether object contains all of the key/value pairs in source,
// comparing nested map[string]any values recursively. Extra keys in object are
// ignored.
func IsMatch(object, source map[string]any) bool {
	for k, sv := range source {
		ov, ok := object[k]
		if !ok {
			return false
		}
		if sm, ok := sv.(map[string]any); ok {
			om, ok := ov.(map[string]any)
			if !ok || !IsMatch(om, sm) {
				return false
			}
			continue
		}
		if !reflect.DeepEqual(ov, sv) {
			return false
		}
	}
	return true
}

// IsString reports whether value's dynamic type is a string.
func IsString(value any) bool { return reflect.ValueOf(value).Kind() == reflect.String }

// IsBool reports whether value's dynamic type is a boolean.
func IsBool(value any) bool { return reflect.ValueOf(value).Kind() == reflect.Bool }

// IsNumber reports whether value's dynamic type is any built-in numeric type
// (integer, unsigned, or floating point).
func IsNumber(value any) bool {
	switch reflect.ValueOf(value).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// IsInteger reports whether value is an integer-valued number: an integer type,
// or a float with no fractional part.
func IsInteger(value any) bool {
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return true
	case reflect.Float32, reflect.Float64:
		f := rv.Float()
		return !math.IsInf(f, 0) && !math.IsNaN(f) && f == math.Trunc(f)
	default:
		return false
	}
}

// IsSlice reports whether value's dynamic type is a slice.
func IsSlice(value any) bool { return reflect.ValueOf(value).Kind() == reflect.Slice }

// IsMap reports whether value's dynamic type is a map.
func IsMap(value any) bool { return reflect.ValueOf(value).Kind() == reflect.Map }

// IsNaN reports whether value is the floating-point Not-a-Number value.
func IsNaN(value any) bool {
	f, ok := value.(float64)
	if ok {
		return math.IsNaN(f)
	}
	if f32, ok := value.(float32); ok {
		return math.IsNaN(float64(f32))
	}
	return false
}

// IsFinite reports whether value is a finite number (not NaN and not an
// infinity). Non-numeric values are not finite.
func IsFinite(value any) bool {
	if !IsNumber(value) {
		return false
	}
	f := ToNumber(value)
	return !math.IsNaN(f) && !math.IsInf(f, 0)
}

// IsObjectLike reports whether value is non-nil and of a composite kind: a map,
// slice, array, struct or pointer.
func IsObjectLike(value any) bool {
	if value == nil {
		return false
	}
	switch reflect.ValueOf(value).Kind() {
	case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct, reflect.Ptr:
		return true
	default:
		return false
	}
}

// ToSafeInteger clamps ToInteger(value) to the range of integers that float64
// can represent exactly, i.e. [-(2^53-1), 2^53-1].
func ToSafeInteger(value any) int {
	const maxSafe = 1<<53 - 1
	n := ToInteger(value)
	if n > maxSafe {
		return maxSafe
	}
	if n < -maxSafe {
		return -maxSafe
	}
	return n
}

// IsError reports whether value implements the error interface and is non-nil.
func IsError(value any) bool {
	if value == nil {
		return false
	}
	_, ok := value.(error)
	return ok
}

// CastArray wraps the provided values into a slice. Called with a single value
// it yields a one-element slice; called with none it yields an empty slice.
func CastArray[T any](values ...T) []T {
	out := make([]T, len(values))
	copy(out, values)
	return out
}

// ToArray returns a copy of the given slice. It is the slice-typed analogue of
// lodash's toArray for the common case of converting/copying a collection.
func ToArray[T any](s []T) []T {
	out := make([]T, len(s))
	copy(out, s)
	return out
}

// DefaultTo returns value unless it is the zero value for its type, in which
// case it returns fallback.
func DefaultTo[T comparable](value, fallback T) T {
	var zero T
	if value == zero {
		return fallback
	}
	return value
}

// Gt reports whether a is greater than b.
func Gt[T Ordered](a, b T) bool { return a > b }

// Gte reports whether a is greater than or equal to b.
func Gte[T Ordered](a, b T) bool { return a >= b }

// Lt reports whether a is less than b.
func Lt[T Ordered](a, b T) bool { return a < b }

// Lte reports whether a is less than or equal to b.
func Lte[T Ordered](a, b T) bool { return a <= b }

// Ordered constrains the types that support the relational operators used by
// Gt, Gte, Lt and Lte.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

// ToNumber converts value to a float64. Booleans become 0 or 1, numeric types
// are converted directly, and strings are parsed (with surrounding whitespace
// trimmed). Anything that cannot be interpreted as a number yields NaN.
func ToNumber(value any) float64 {
	switch v := value.(type) {
	case nil:
		return 0
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		f, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if err != nil {
			return math.NaN()
		}
		return f
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return rv.Float()
	default:
		return math.NaN()
	}
}

// ToInteger converts value to an int by truncating toward zero. Non-numeric and
// non-finite inputs yield 0.
func ToInteger(value any) int {
	f := ToNumber(value)
	if math.IsNaN(f) {
		return 0
	}
	if math.IsInf(f, 1) {
		return math.MaxInt
	}
	if math.IsInf(f, -1) {
		return math.MinInt
	}
	return int(math.Trunc(f))
}

// ToFinite converts value to a finite float64, mapping NaN to 0 and infinities
// to the largest representable magnitudes.
func ToFinite(value any) float64 {
	f := ToNumber(value)
	if math.IsNaN(f) {
		return 0
	}
	if math.IsInf(f, 1) {
		return math.MaxFloat64
	}
	if math.IsInf(f, -1) {
		return -math.MaxFloat64
	}
	return f
}

// ToString converts value to its string representation. Nil becomes the empty
// string; all other values use fmt.Sprint.
func ToString(value any) string {
	if value == nil {
		return ""
	}
	if s, ok := value.(string); ok {
		return s
	}
	return fmt.Sprint(value)
}

// Conforms builds a predicate that reports whether a map[string]any satisfies
// every predicate in spec. Each spec entry names a key that must be present and
// whose value must satisfy the associated predicate.
func Conforms(spec map[string]func(any) bool) func(map[string]any) bool {
	return func(m map[string]any) bool {
		return ConformsTo(m, spec)
	}
}

// ConformsTo reports whether object satisfies every predicate in spec.
func ConformsTo(object map[string]any, spec map[string]func(any) bool) bool {
	for k, pred := range spec {
		v, ok := object[k]
		if !ok || !pred(v) {
			return false
		}
	}
	return true
}
