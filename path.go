package lodash

import "strconv"

// ToPath splits a property path string into its component keys. It understands
// dotted notation ("a.b.c") and bracket notation with optional quotes
// ("a[0].b", "a['b.c']"), yielding ["a", "0", "b"] and ["a", "b.c"]
// respectively.
func ToPath(path string) []string {
	var (
		parts []string
		buf   []rune
	)
	flush := func() {
		if len(buf) > 0 {
			parts = append(parts, string(buf))
			buf = buf[:0]
		}
	}
	runes := []rune(path)
	for i := 0; i < len(runes); i++ {
		switch c := runes[i]; c {
		case '.':
			flush()
		case '[':
			flush()
			// Consume up to the matching ']'.
			i++
			var quote rune
			if i < len(runes) && (runes[i] == '\'' || runes[i] == '"') {
				quote = runes[i]
				i++
			}
			for i < len(runes) && runes[i] != ']' {
				if quote != 0 && runes[i] == quote {
					i++
					break
				}
				buf = append(buf, runes[i])
				i++
			}
			// Advance past a closing ']' if present.
			for i < len(runes) && runes[i] != ']' {
				i++
			}
			flush()
		default:
			buf = append(buf, c)
		}
	}
	flush()
	return parts
}

// Get retrieves a nested value from a map[string]any using a property path.
// The path may use dotted and bracket notation (see ToPath) and may traverse
// []any slices via integer indices. The second return value reports whether the
// full path resolved to a value.
func Get(m map[string]any, path string) (any, bool) {
	keys := ToPath(path)
	if len(keys) == 0 {
		return nil, false
	}
	var cur any = m
	for _, k := range keys {
		switch node := cur.(type) {
		case map[string]any:
			v, ok := node[k]
			if !ok {
				return nil, false
			}
			cur = v
		case []any:
			idx, err := strconv.Atoi(k)
			if err != nil || idx < 0 || idx >= len(node) {
				return nil, false
			}
			cur = node[idx]
		default:
			return nil, false
		}
	}
	return cur, true
}

// GetOr is like Get but returns fallback when the path does not resolve.
func GetOr(m map[string]any, path string, fallback any) any {
	if v, ok := Get(m, path); ok {
		return v
	}
	return fallback
}

// Has reports whether path resolves to a value in the nested map.
func Has(m map[string]any, path string) bool {
	_, ok := Get(m, path)
	return ok
}

// At returns the values found at each of the given paths, in order. Paths that
// do not resolve yield a nil entry.
func At(m map[string]any, paths ...string) []any {
	r := make([]any, len(paths))
	for i, p := range paths {
		if v, ok := Get(m, p); ok {
			r[i] = v
		}
	}
	return r
}

// Result is like Get but, when the resolved value is a func() any, it invokes
// that function and returns its result. This mirrors lodash's _.result.
func Result(m map[string]any, path string) (any, bool) {
	v, ok := Get(m, path)
	if !ok {
		return nil, false
	}
	if fn, isFn := v.(func() any); isFn {
		return fn(), true
	}
	return v, true
}

// Set returns a copy of m with the value at path set to value, creating
// intermediate maps as needed. The input map is not mutated; only the
// containers along the path are cloned (copy-on-write).
func Set(m map[string]any, path string, value any) map[string]any {
	keys := ToPath(path)
	if len(keys) == 0 {
		return m
	}
	result := setPath(m, keys, value)
	if rm, ok := result.(map[string]any); ok {
		return rm
	}
	return m
}

func setPath(node any, keys []string, value any) any {
	key := keys[0]
	last := len(keys) == 1

	m, _ := node.(map[string]any)
	out := make(map[string]any, len(m)+1)
	for k, v := range m {
		out[k] = v
	}
	if last {
		out[key] = value
		return out
	}
	out[key] = setPath(out[key], keys[1:], value)
	return out
}

// Update returns a copy of m with the value at path replaced by the result of
// updater applied to the existing value (or nil when the path is absent).
// Intermediate maps are created as needed and the input is not mutated.
func Update(m map[string]any, path string, updater func(old any) any) map[string]any {
	old, _ := Get(m, path)
	return Set(m, path, updater(old))
}

// Unset returns a copy of m with the value at path removed, along with a boolean
// reporting whether anything was removed. The input map is not mutated.
func Unset(m map[string]any, path string) (map[string]any, bool) {
	keys := ToPath(path)
	if len(keys) == 0 {
		return m, false
	}
	out, removed := unsetPath(m, keys)
	rm, _ := out.(map[string]any)
	return rm, removed
}

func unsetPath(node any, keys []string) (any, bool) {
	m, ok := node.(map[string]any)
	if !ok {
		return node, false
	}
	key := keys[0]
	out := make(map[string]any, len(m))
	for k, v := range m {
		out[k] = v
	}
	if len(keys) == 1 {
		if _, exists := out[key]; !exists {
			return out, false
		}
		delete(out, key)
		return out, true
	}
	child, exists := out[key]
	if !exists {
		return out, false
	}
	newChild, removed := unsetPath(child, keys[1:])
	out[key] = newChild
	return out, removed
}
