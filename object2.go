package lodash

import "reflect"

// Size returns the number of elements in collection, which may be a map, slice,
// array, string or channel. Any other type reports 0.
func Size(collection any) int {
	if collection == nil {
		return 0
	}
	rv := reflect.ValueOf(collection)
	switch rv.Kind() {
	case reflect.Map, reflect.Slice, reflect.Array, reflect.String, reflect.Chan:
		return rv.Len()
	default:
		return 0
	}
}

// ToPairs returns the key/value pairs of the map as a slice of Pair values in
// unspecified order.
func ToPairs[K comparable, V any](m map[K]V) []Pair[K, V] {
	r := make([]Pair[K, V], 0, len(m))
	for k, v := range m {
		r = append(r, Pair[K, V]{First: k, Second: v})
	}
	return r
}

// FromPairs builds a map from a slice of key/value Pair values. Later pairs
// overwrite earlier ones with the same key.
func FromPairs[K comparable, V any](pairs []Pair[K, V]) map[K]V {
	r := make(map[K]V, len(pairs))
	for _, p := range pairs {
		r[p.First] = p.Second
	}
	return r
}

// ZipObject builds a map by pairing each key with the value at the same index.
// Extra keys map to the zero value; extra values are ignored.
func ZipObject[K comparable, V any](keys []K, values []V) map[K]V {
	r := make(map[K]V, len(keys))
	for i, k := range keys {
		if i < len(values) {
			r[k] = values[i]
		} else {
			var zero V
			r[k] = zero
		}
	}
	return r
}

// FindKey returns the first key whose key/value pair satisfies pred, along with
// true. Because Go map iteration is unordered, "first" is not deterministic
// across runs when multiple keys match. It returns the zero key and false when
// nothing matches.
func FindKey[K comparable, V any](m map[K]V, pred func(K, V) bool) (K, bool) {
	for k, v := range m {
		if pred(k, v) {
			return k, true
		}
	}
	var zero K
	return zero, false
}

// InvertBy groups the keys of m by the group value produced by fn from each
// value, returning a map from group to the slice of keys that produced it.
func InvertBy[K comparable, V any, G comparable](m map[K]V, fn func(V) G) map[G][]K {
	r := make(map[G][]K)
	for k, v := range m {
		g := fn(v)
		r[g] = append(r[g], k)
	}
	return r
}

// Transform reduces the map into a single accumulated value by applying fn to
// each key/value pair, seeded with init.
func Transform[K comparable, V, A any](m map[K]V, fn func(acc A, v V, k K) A, init A) A {
	acc := init
	for k, v := range m {
		acc = fn(acc, v, k)
	}
	return acc
}

// ForOwn invokes fn for each key/value pair of the map. Iteration order is
// unspecified.
func ForOwn[K comparable, V any](m map[K]V, fn func(K, V)) {
	for k, v := range m {
		fn(k, v)
	}
}

// MergeWith combines all provided maps into a new map. When a key appears in
// more than one map, the resolver fn decides the merged value from the existing
// (accumulated) value and the incoming one.
func MergeWith[K comparable, V any](fn func(existing, incoming V) V, maps ...map[K]V) map[K]V {
	r := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			if existing, ok := r[k]; ok {
				r[k] = fn(existing, v)
			} else {
				r[k] = v
			}
		}
	}
	return r
}

// AssignWith copies key/value pairs from sources into a new map seeded with dst.
// When a key already has a value, the resolver fn decides the result from the
// destination value and the incoming source value. The input maps are not
// mutated.
func AssignWith[K comparable, V any](dst map[K]V, fn func(dstV, srcV V) V, sources ...map[K]V) map[K]V {
	r := make(map[K]V, len(dst))
	for k, v := range dst {
		r[k] = v
	}
	for _, m := range sources {
		for k, v := range m {
			if existing, ok := r[k]; ok {
				r[k] = fn(existing, v)
			} else {
				r[k] = v
			}
		}
	}
	return r
}

// Defaults returns a new map seeded with dst and filled in with entries from
// sources for any keys that are not already present. Earlier sources take
// precedence over later ones. The inputs are not mutated.
func Defaults[K comparable, V any](dst map[K]V, sources ...map[K]V) map[K]V {
	r := make(map[K]V, len(dst))
	for k, v := range dst {
		r[k] = v
	}
	for _, m := range sources {
		for k, v := range m {
			if _, ok := r[k]; !ok {
				r[k] = v
			}
		}
	}
	return r
}

// DefaultsDeep is like Defaults but recurses into nested map[string]any values,
// filling in missing keys at every level. The inputs are not mutated.
func DefaultsDeep(dst map[string]any, sources ...map[string]any) map[string]any {
	r := make(map[string]any, len(dst))
	for k, v := range dst {
		r[k] = v
	}
	for _, m := range sources {
		for k, v := range m {
			existing, ok := r[k]
			if !ok {
				r[k] = v
				continue
			}
			em, eok := existing.(map[string]any)
			vm, vok := v.(map[string]any)
			if eok && vok {
				r[k] = DefaultsDeep(em, vm)
			}
		}
	}
	return r
}
