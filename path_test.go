package lodash_test

import (
	"testing"

	lodash "github.com/malcolmston/lodash"
)

func TestToPath(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"a.b.c", []string{"a", "b", "c"}},
		{"a[0].b", []string{"a", "0", "b"}},
		{"a['b.c'].d", []string{"a", "b.c", "d"}},
		{"a[\"x\"]", []string{"a", "x"}},
		{"", nil},
	}
	for _, c := range cases {
		got := lodash.ToPath(c.in)
		if !lodash.IsEqual(got, c.want) {
			t.Errorf("ToPath(%q) = %v want %v", c.in, got, c.want)
		}
	}
}

func sampleObj() map[string]any {
	return map[string]any{
		"a": map[string]any{
			"b": []any{map[string]any{"c": 3}},
		},
		"x": 1,
	}
}

func TestGetHasBracket(t *testing.T) {
	obj := sampleObj()
	if v, ok := lodash.Get(obj, "a.b[0].c"); !ok || v != 3 {
		t.Errorf("Get bracket = %v %v", v, ok)
	}
	if _, ok := lodash.Get(obj, "a.b[5].c"); ok {
		t.Errorf("Get out of range should fail")
	}
	if _, ok := lodash.Get(obj, "nope.here"); ok {
		t.Errorf("Get missing should fail")
	}
	if !lodash.Has(obj, "a.b[0].c") || lodash.Has(obj, "a.z") {
		t.Errorf("Has")
	}
	if lodash.GetOr(obj, "a.z", "def") != "def" {
		t.Errorf("GetOr fallback")
	}
	if lodash.GetOr(obj, "x", "def") != 1 {
		t.Errorf("GetOr present")
	}
}

func TestAt(t *testing.T) {
	obj := sampleObj()
	got := lodash.At(obj, "x", "a.b[0].c", "missing")
	if got[0] != 1 || got[1] != 3 || got[2] != nil {
		t.Errorf("At = %v", got)
	}
}

func TestResult(t *testing.T) {
	obj := map[string]any{
		"fn":  func() any { return "computed" },
		"val": 5,
	}
	if v, ok := lodash.Result(obj, "fn"); !ok || v != "computed" {
		t.Errorf("Result func")
	}
	if v, ok := lodash.Result(obj, "val"); !ok || v != 5 {
		t.Errorf("Result value")
	}
	if _, ok := lodash.Result(obj, "missing"); ok {
		t.Errorf("Result missing")
	}
}

func TestSet(t *testing.T) {
	obj := sampleObj()
	out := lodash.Set(obj, "a.new.deep", 9)
	if v, _ := lodash.Get(out, "a.new.deep"); v != 9 {
		t.Errorf("Set nested = %v", v)
	}
	// Original not mutated.
	if lodash.Has(obj, "a.new.deep") {
		t.Errorf("Set mutated original")
	}
	// Existing sibling preserved.
	if v, _ := lodash.Get(out, "x"); v != 1 {
		t.Errorf("Set dropped sibling")
	}
	if v, _ := lodash.Get(out, "a.b[0].c"); v != 3 {
		t.Errorf("Set dropped nested sibling")
	}
}

func TestUpdate(t *testing.T) {
	obj := map[string]any{"count": 5}
	out := lodash.Update(obj, "count", func(old any) any {
		return lodash.ToInteger(old) + 1
	})
	if v, _ := lodash.Get(out, "count"); v != 6 {
		t.Errorf("Update = %v", v)
	}
	out2 := lodash.Update(obj, "fresh", func(old any) any {
		if old == nil {
			return "init"
		}
		return old
	})
	if v, _ := lodash.Get(out2, "fresh"); v != "init" {
		t.Errorf("Update absent")
	}
}

func TestUnset(t *testing.T) {
	obj := map[string]any{"a": map[string]any{"b": 1, "c": 2}}
	out, ok := lodash.Unset(obj, "a.b")
	if !ok {
		t.Errorf("Unset should report removed")
	}
	if lodash.Has(out, "a.b") {
		t.Errorf("Unset did not remove")
	}
	if !lodash.Has(out, "a.c") {
		t.Errorf("Unset removed sibling")
	}
	// Original not mutated.
	if !lodash.Has(obj, "a.b") {
		t.Errorf("Unset mutated original")
	}
	if _, ok := lodash.Unset(obj, "a.z"); ok {
		t.Errorf("Unset missing should report false")
	}
}
