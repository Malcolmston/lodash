package lodash_test

import (
	"errors"
	"strings"
	"testing"

	lodash "github.com/malcolmston/lodash"
)

func TestIdentityConstantNoop(t *testing.T) {
	if lodash.Identity(7) != 7 {
		t.Errorf("Identity")
	}
	c := lodash.Constant("x")
	if c() != "x" {
		t.Errorf("Constant")
	}
	lodash.Noop() // must not panic
}

func TestTimes(t *testing.T) {
	got := lodash.Times(3, func(i int) int { return i * i })
	if !lodash.IsEqual(got, []int{0, 1, 4}) {
		t.Errorf("Times = %v", got)
	}
	if len(lodash.Times(-1, func(i int) int { return i })) != 0 {
		t.Errorf("Times negative")
	}
}

func TestStubs(t *testing.T) {
	if len(lodash.StubArray[int]()) != 0 {
		t.Errorf("StubArray")
	}
	if len(lodash.StubObject()) != 0 {
		t.Errorf("StubObject")
	}
	if lodash.StubString() != "" {
		t.Errorf("StubString")
	}
	if !lodash.StubTrue() || lodash.StubFalse() {
		t.Errorf("StubTrue/StubFalse")
	}
}

func TestUniqueID(t *testing.T) {
	a := lodash.UniqueID("id_")
	b := lodash.UniqueID("id_")
	if a == b {
		t.Errorf("UniqueID should be unique: %s == %s", a, b)
	}
	if !strings.HasPrefix(a, "id_") {
		t.Errorf("UniqueID prefix")
	}
}

func TestPropertyMatches(t *testing.T) {
	obj := map[string]any{"a": map[string]any{"b": 5}}
	get := lodash.Property("a.b")
	if v, ok := get(obj); !ok || v != 5 {
		t.Errorf("Property")
	}
	of := lodash.PropertyOf(obj)
	if v, ok := of("a.b"); !ok || v != 5 {
		t.Errorf("PropertyOf")
	}
	m := lodash.Matches(map[string]any{"a": map[string]any{"b": 5}})
	if !m(obj) {
		t.Errorf("Matches")
	}
}

func TestAttempt(t *testing.T) {
	v, err := lodash.Attempt(func() int { return 42 })
	if err != nil || v != 42 {
		t.Errorf("Attempt success")
	}
	_, err = lodash.Attempt(func() int { panic("boom") })
	if err == nil || err.Error() != "boom" {
		t.Errorf("Attempt panic string, got %v", err)
	}
	_, err = lodash.Attempt(func() int { panic(errors.New("kaboom")) })
	if err == nil || err.Error() != "kaboom" {
		t.Errorf("Attempt panic error, got %v", err)
	}
}
