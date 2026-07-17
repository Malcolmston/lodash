package lodash_test

import (
	"errors"
	"math"
	"testing"

	lodash "github.com/malcolmston/lodash"
)

func TestCloneAndCloneDeep(t *testing.T) {
	orig := map[string]any{"a": 1, "nested": map[string]any{"b": 2}}

	shallow := lodash.Clone(orig)
	shallow["a"] = 99
	if orig["a"] != 1 {
		t.Fatalf("Clone mutated top-level of original")
	}
	// Shallow clone shares nested containers.
	shallow["nested"].(map[string]any)["b"] = 42
	if orig["nested"].(map[string]any)["b"] != 42 {
		t.Fatalf("Clone should share nested maps")
	}

	orig2 := map[string]any{"nested": map[string]any{"b": 2}, "list": []int{1, 2}}
	deep := lodash.CloneDeep(orig2)
	deep["nested"].(map[string]any)["b"] = 42
	deep["list"].([]int)[0] = 9
	if orig2["nested"].(map[string]any)["b"] != 2 {
		t.Fatalf("CloneDeep should not share nested maps")
	}
	if orig2["list"].([]int)[0] != 1 {
		t.Fatalf("CloneDeep should not share slices")
	}
	if !lodash.IsEqual(orig2, map[string]any{"nested": map[string]any{"b": 2}, "list": []int{1, 2}}) {
		t.Fatalf("CloneDeep changed original")
	}
}

func TestCloneScalar(t *testing.T) {
	if lodash.Clone(5) != 5 || lodash.CloneDeep("x") != "x" {
		t.Fatalf("clone of scalar failed")
	}
}

func TestIsEqualAndEq(t *testing.T) {
	if !lodash.IsEqual([]int{1, 2}, []int{1, 2}) {
		t.Fatalf("IsEqual slices")
	}
	if lodash.IsEqual(1, 2) {
		t.Fatalf("IsEqual mismatch")
	}
	if !lodash.Eq(3, 3) || lodash.Eq("a", "b") {
		t.Fatalf("Eq failed")
	}
}

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		v    any
		want bool
	}{
		{nil, true},
		{"", true},
		{"x", false},
		{[]int{}, true},
		{[]int{1}, false},
		{map[string]int{}, true},
		{map[string]int{"a": 1}, false},
		{0, true},
		{5, true},
		{true, true},
	}
	for _, c := range cases {
		if got := lodash.IsEmpty(c.v); got != c.want {
			t.Errorf("IsEmpty(%v) = %v, want %v", c.v, got, c.want)
		}
	}
	var p *int
	if !lodash.IsEmpty(p) {
		t.Errorf("IsEmpty(nil ptr) should be true")
	}
}

func TestIsNil(t *testing.T) {
	var p *int
	var m map[string]int
	if !lodash.IsNil(nil) || !lodash.IsNil(p) || !lodash.IsNil(m) {
		t.Fatalf("IsNil should be true")
	}
	if lodash.IsNil(5) || lodash.IsNil("") {
		t.Fatalf("IsNil should be false")
	}
}

func TestIsPlainObject(t *testing.T) {
	if !lodash.IsPlainObject(map[string]any{"a": 1}) {
		t.Fatalf("map[string]any should be plain object")
	}
	if lodash.IsPlainObject(map[int]int{1: 1}) || lodash.IsPlainObject([]int{1}) || lodash.IsPlainObject(nil) {
		t.Fatalf("non string-keyed map should not be plain object")
	}
}

func TestIsMatch(t *testing.T) {
	obj := map[string]any{"a": 1, "b": 2, "c": map[string]any{"d": 3, "e": 4}}
	if !lodash.IsMatch(obj, map[string]any{"a": 1}) {
		t.Fatalf("IsMatch simple")
	}
	if !lodash.IsMatch(obj, map[string]any{"c": map[string]any{"d": 3}}) {
		t.Fatalf("IsMatch nested")
	}
	if lodash.IsMatch(obj, map[string]any{"a": 99}) {
		t.Fatalf("IsMatch should fail on value mismatch")
	}
	if lodash.IsMatch(obj, map[string]any{"z": 1}) {
		t.Fatalf("IsMatch should fail on missing key")
	}
}

func TestTypePredicates(t *testing.T) {
	if !lodash.IsString("x") || lodash.IsString(1) {
		t.Errorf("IsString")
	}
	if !lodash.IsBool(true) || lodash.IsBool(1) {
		t.Errorf("IsBool")
	}
	if !lodash.IsNumber(1) || !lodash.IsNumber(1.5) || lodash.IsNumber("x") {
		t.Errorf("IsNumber")
	}
	if !lodash.IsInteger(5) || !lodash.IsInteger(5.0) || lodash.IsInteger(5.5) {
		t.Errorf("IsInteger")
	}
	if !lodash.IsSlice([]int{}) || lodash.IsSlice(3) {
		t.Errorf("IsSlice")
	}
	if !lodash.IsMap(map[string]int{}) || lodash.IsMap(3) {
		t.Errorf("IsMap")
	}
	if !lodash.IsError(errors.New("x")) || lodash.IsError("x") || lodash.IsError(nil) {
		t.Errorf("IsError")
	}
	if !lodash.IsObjectLike([]int{1}) || lodash.IsObjectLike(3) || lodash.IsObjectLike(nil) {
		t.Errorf("IsObjectLike")
	}
	if !lodash.IsNaN(math.NaN()) || lodash.IsNaN(1.0) {
		t.Errorf("IsNaN")
	}
	if !lodash.IsFinite(1) || lodash.IsFinite(math.Inf(1)) || lodash.IsFinite("x") {
		t.Errorf("IsFinite")
	}
}

func TestCastArrayAndToArray(t *testing.T) {
	if !lodash.IsEqual(lodash.CastArray(1), []int{1}) {
		t.Errorf("CastArray single")
	}
	if len(lodash.CastArray[int]()) != 0 {
		t.Errorf("CastArray empty")
	}
	src := []int{1, 2}
	cp := lodash.ToArray(src)
	cp[0] = 9
	if src[0] != 1 {
		t.Errorf("ToArray should copy")
	}
}

func TestDefaultToAndRelational(t *testing.T) {
	if lodash.DefaultTo(0, 5) != 5 || lodash.DefaultTo(3, 5) != 3 {
		t.Errorf("DefaultTo")
	}
	if !lodash.Gt(2, 1) || !lodash.Gte(2, 2) || !lodash.Lt(1, 2) || !lodash.Lte(2, 2) {
		t.Errorf("relational")
	}
}

func TestToNumber(t *testing.T) {
	cases := []struct {
		in   any
		want float64
	}{
		{nil, 0}, {true, 1}, {false, 0}, {"42", 42}, {" 3.5 ", 3.5}, {int8(7), 7}, {uint(9), 9}, {2.5, 2.5},
	}
	for _, c := range cases {
		if got := lodash.ToNumber(c.in); got != c.want {
			t.Errorf("ToNumber(%v) = %v want %v", c.in, got, c.want)
		}
	}
	if !math.IsNaN(lodash.ToNumber("abc")) {
		t.Errorf("ToNumber non-numeric string should be NaN")
	}
}

func TestToIntegerFiniteSafe(t *testing.T) {
	if lodash.ToInteger("3.9") != 3 || lodash.ToInteger(-2.9) != -2 || lodash.ToInteger("x") != 0 {
		t.Errorf("ToInteger")
	}
	if lodash.ToFinite("x") != 0 || lodash.ToFinite(2.5) != 2.5 {
		t.Errorf("ToFinite")
	}
	if lodash.ToFinite(math.Inf(1)) != math.MaxFloat64 {
		t.Errorf("ToFinite inf")
	}
	if lodash.ToSafeInteger(1<<60) != 1<<53-1 {
		t.Errorf("ToSafeInteger clamp high")
	}
	if lodash.ToSafeInteger(5) != 5 {
		t.Errorf("ToSafeInteger passthrough")
	}
}

func TestToString(t *testing.T) {
	if lodash.ToString(nil) != "" || lodash.ToString("x") != "x" || lodash.ToString(42) != "42" {
		t.Errorf("ToString")
	}
}

func TestConforms(t *testing.T) {
	spec := map[string]func(any) bool{
		"age": func(v any) bool { return lodash.ToNumber(v) >= 18 },
		"ok":  func(v any) bool { b, _ := v.(bool); return b },
	}
	pred := lodash.Conforms(spec)
	if !pred(map[string]any{"age": 20, "ok": true}) {
		t.Errorf("Conforms should match")
	}
	if pred(map[string]any{"age": 5, "ok": true}) {
		t.Errorf("Conforms should fail on predicate")
	}
	if lodash.ConformsTo(map[string]any{"age": 20}, spec) {
		t.Errorf("ConformsTo should fail on missing key")
	}
}
