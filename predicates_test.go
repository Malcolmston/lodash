package lodash

import (
	"regexp"
	"testing"
	"time"
)

func TestIsFunction(t *testing.T) {
	cases := []struct {
		name string
		val  any
		want bool
	}{
		{"func", func() {}, true},
		{"method-like", TestIsFunction, true},
		{"int", 3, false},
		{"nil", nil, false},
		{"string", "x", false},
	}
	for _, c := range cases {
		if got := IsFunction(c.val); got != c.want {
			t.Errorf("IsFunction(%s) = %v, want %v", c.name, got, c.want)
		}
	}
}

func TestIsRegExpDateBuffer(t *testing.T) {
	if !IsRegExp(regexp.MustCompile(`a`)) {
		t.Error("IsRegExp(*regexp.Regexp) = false")
	}
	if IsRegExp("a") {
		t.Error("IsRegExp(string) = true")
	}
	now := time.Now()
	if !IsDate(now) || !IsDate(&now) {
		t.Error("IsDate(time.Time) failed")
	}
	if IsDate("2020") {
		t.Error("IsDate(string) = true")
	}
	if !IsBuffer([]byte("x")) {
		t.Error("IsBuffer([]byte) = false")
	}
	if IsBuffer([]int{1}) {
		t.Error("IsBuffer([]int) = true")
	}
}

func TestIsSafeIntegerAndLength(t *testing.T) {
	cases := []struct {
		val        any
		safe, leng bool
	}{
		{3, true, true},
		{-3, true, false},
		{0, true, true},
		{3.0, true, true},
		{3.5, false, false},
		{maxSafeInt, true, true},
		{maxSafeInt + 1, false, false},
		{"3", false, false},
	}
	for _, c := range cases {
		if got := IsSafeInteger(c.val); got != c.safe {
			t.Errorf("IsSafeInteger(%v) = %v, want %v", c.val, got, c.safe)
		}
		if got := IsLength(c.val); got != c.leng {
			t.Errorf("IsLength(%v) = %v, want %v", c.val, got, c.leng)
		}
	}
}

func TestIsNullUndefinedObject(t *testing.T) {
	var p *int
	if !IsNull(nil) || !IsNull(p) {
		t.Error("IsNull failed for nil values")
	}
	if IsNull(3) {
		t.Error("IsNull(3) = true")
	}
	if !IsUndefined(nil) {
		t.Error("IsUndefined(nil) = false")
	}
	if IsUndefined(p) {
		t.Error("IsUndefined(typed nil) = true")
	}
	objCases := []struct {
		val  any
		want bool
	}{
		{map[string]any{}, true},
		{[]int{}, true},
		{func() {}, true},
		{struct{}{}, true},
		{3, false},
		{"x", false},
		{true, false},
		{nil, false},
	}
	for _, c := range objCases {
		if got := IsObject(c.val); got != c.want {
			t.Errorf("IsObject(%v) = %v, want %v", c.val, got, c.want)
		}
	}
}

func TestIsArrayLike(t *testing.T) {
	cases := []struct {
		val              any
		like, likeObject bool
	}{
		{[]int{1}, true, true},
		{[3]int{}, true, true},
		{"abc", true, false},
		{map[string]int{}, false, false},
		{3, false, false},
		{func() {}, false, false},
	}
	for _, c := range cases {
		if got := IsArrayLike(c.val); got != c.like {
			t.Errorf("IsArrayLike(%v) = %v, want %v", c.val, got, c.like)
		}
		if got := IsArrayLikeObject(c.val); got != c.likeObject {
			t.Errorf("IsArrayLikeObject(%v) = %v, want %v", c.val, got, c.likeObject)
		}
	}
}

func TestToLength(t *testing.T) {
	cases := []struct {
		val  any
		want int
	}{
		{3, 3},
		{3.9, 3},
		{-5, 0},
		{"7", 7},
		{maxSafeInt + 10, maxSafeInt},
	}
	for _, c := range cases {
		if got := ToLength(c.val); got != c.want {
			t.Errorf("ToLength(%v) = %d, want %d", c.val, got, c.want)
		}
	}
}

func TestIsEqualWith(t *testing.T) {
	// Case-insensitive string comparison via customizer.
	ci := func(a, b any) (bool, bool) {
		sa, oka := a.(string)
		sb, okb := b.(string)
		if oka && okb {
			return len(sa) == len(sb), true
		}
		return false, false
	}
	if !IsEqualWith("abc", "xyz", ci) {
		t.Error("IsEqualWith customizer not applied")
	}
	if IsEqualWith("ab", "xyz", ci) {
		t.Error("IsEqualWith customizer wrong result")
	}
	// Falls back to IsEqual when not handled.
	if !IsEqualWith([]int{1, 2}, []int{1, 2}, ci) {
		t.Error("IsEqualWith fallback failed")
	}
	if !IsEqualWith(1, 1, nil) {
		t.Error("IsEqualWith(nil customizer) failed")
	}
}

func TestIsMatchWith(t *testing.T) {
	obj := map[string]any{"a": 1, "b": 2, "c": 3}
	always := func(_, _ any) bool { return true }
	if !IsMatchWith(obj, map[string]any{"a": 99}, always) {
		t.Error("IsMatchWith(always) = false")
	}
	eq := func(o, s any) bool { return o == s }
	if !IsMatchWith(obj, map[string]any{"a": 1, "b": 2}, eq) {
		t.Error("IsMatchWith(eq) failed on matching subset")
	}
	if IsMatchWith(obj, map[string]any{"a": 5}, eq) {
		t.Error("IsMatchWith(eq) matched wrong value")
	}
	if IsMatchWith(obj, map[string]any{"z": 1}, eq) {
		t.Error("IsMatchWith matched missing key")
	}
	if !IsMatchWith(obj, map[string]any{}, eq) {
		t.Error("IsMatchWith(empty source) = false")
	}
}
