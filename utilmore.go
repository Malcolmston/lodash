package lodash

// MatchesProperty returns a predicate that reports whether the value found at
// path (see ToPath) within a map[string]any is deeply equal to srcValue. It is
// the property-path analogue of Matches and mirrors lodash's matchesProperty.
func MatchesProperty(path string, srcValue any) func(map[string]any) bool {
	return func(m map[string]any) bool {
		v, ok := Get(m, path)
		return ok && IsEqual(v, srcValue)
	}
}
