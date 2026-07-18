package lodash

// SetWith returns a copy of m with value assigned at path, using customizer to
// create any intermediate containers that do not yet exist. For each missing
// segment, customizer is called with the key of the segment about to be
// created; if it returns a non-nil map[string]any that container is used,
// otherwise a fresh map[string]any is created. The input map is not mutated. It
// mirrors lodash's setWith.
func SetWith(m map[string]any, path string, value any, customizer func(key string) any) map[string]any {
	keys := ToPath(path)
	if len(keys) == 0 {
		return m
	}
	result := setWithPath(m, keys, value, customizer)
	if rm, ok := result.(map[string]any); ok {
		return rm
	}
	return m
}

// setWithPath is the recursive worker for SetWith. It rebuilds each visited
// map so the original structure is left unmodified.
func setWithPath(node any, keys []string, value any, customizer func(key string) any) any {
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
	child, exists := out[key]
	if !exists {
		if customizer != nil {
			if created, ok := customizer(keys[1]).(map[string]any); ok && created != nil {
				child = created
			}
		}
	}
	out[key] = setWithPath(child, keys[1:], value, customizer)
	return out
}

// UpdateWith returns a copy of m with the value at path replaced by the result
// of updater applied to the existing value (or nil when the path is absent),
// using customizer to create any intermediate containers exactly as SetWith
// does. The input map is not mutated. It mirrors lodash's updateWith.
func UpdateWith(m map[string]any, path string, updater func(old any) any, customizer func(key string) any) map[string]any {
	old, _ := Get(m, path)
	return SetWith(m, path, updater(old), customizer)
}
