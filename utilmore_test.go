package lodash

import "testing"

func TestMatchesProperty(t *testing.T) {
	users := []map[string]any{
		{"name": "fred", "age": 40},
		{"name": "barney", "age": 36},
	}
	isFred := MatchesProperty("name", "fred")
	if !isFred(users[0]) {
		t.Error("MatchesProperty did not match fred")
	}
	if isFred(users[1]) {
		t.Error("MatchesProperty wrongly matched barney")
	}
	// Nested path.
	nested := map[string]any{"a": map[string]any{"b": 5}}
	if !MatchesProperty("a.b", 5)(nested) {
		t.Error("MatchesProperty nested path failed")
	}
	// Missing path returns false.
	if MatchesProperty("a.z", 5)(nested) {
		t.Error("MatchesProperty matched missing path")
	}
}
