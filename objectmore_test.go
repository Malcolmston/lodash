package lodash

import "testing"

func TestSetWith(t *testing.T) {
	// Default container creation (map) when customizer returns nil.
	got := SetWith(map[string]any{}, "a.b.c", 1, func(key string) any { return nil })
	if v, ok := Get(got, "a.b.c"); !ok || v != 1 {
		t.Errorf("SetWith default = %v,%v", v, ok)
	}
	// Customizer provides a pre-seeded container.
	seed := func(key string) any {
		return map[string]any{"seeded": true}
	}
	got2 := SetWith(map[string]any{}, "x.y", 9, seed)
	if v, ok := Get(got2, "x.y"); !ok || v != 9 {
		t.Errorf("SetWith seeded value = %v,%v", v, ok)
	}
	if v, ok := Get(got2, "x.seeded"); !ok || v != true {
		t.Errorf("SetWith seeded container = %v,%v", v, ok)
	}
	// Input is not mutated.
	orig := map[string]any{"a": 1}
	_ = SetWith(orig, "a", 2, nil)
	if orig["a"] != 1 {
		t.Errorf("SetWith mutated input: %v", orig)
	}
}

func TestUpdateWith(t *testing.T) {
	m := map[string]any{"a": map[string]any{"count": 1}}
	got := UpdateWith(m, "a.count", func(old any) any {
		if n, ok := old.(int); ok {
			return n + 10
		}
		return 0
	}, nil)
	if v, ok := Get(got, "a.count"); !ok || v != 11 {
		t.Errorf("UpdateWith = %v,%v", v, ok)
	}
	// Absent path: updater receives nil.
	got2 := UpdateWith(map[string]any{}, "new.path", func(old any) any {
		if old == nil {
			return "created"
		}
		return "wrong"
	}, nil)
	if v, ok := Get(got2, "new.path"); !ok || v != "created" {
		t.Errorf("UpdateWith absent = %v,%v", v, ok)
	}
}
