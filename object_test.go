package lodash

import (
	"reflect"
	"sort"
	"testing"
)

func TestKeysValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	ks := Keys(m)
	sort.Strings(ks)
	if !reflect.DeepEqual(ks, []string{"a", "b", "c"}) {
		t.Fatalf("keys = %v", ks)
	}
	if !reflect.DeepEqual(SortedKeys(m), []string{"a", "b", "c"}) {
		t.Fatalf("sortedkeys = %v", SortedKeys(m))
	}
	vs := Values(m)
	sort.Ints(vs)
	if !reflect.DeepEqual(vs, []int{1, 2, 3}) {
		t.Fatalf("values = %v", vs)
	}
}

func TestEntries(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	es := Entries(m)
	sort.Slice(es, func(i, j int) bool { return es[i].Key < es[j].Key })
	if len(es) != 2 || es[0].Key != "a" || es[0].Value != 1 {
		t.Fatalf("entries = %v", es)
	}
	back := FromEntries(es)
	if !reflect.DeepEqual(back, m) {
		t.Fatalf("fromentries = %v", back)
	}
}

func TestPickOmit(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	if !reflect.DeepEqual(Pick(m, "a", "c", "z"), map[string]int{"a": 1, "c": 3}) {
		t.Fatalf("pick = %v", Pick(m, "a", "c", "z"))
	}
	if !reflect.DeepEqual(Omit(m, "b"), map[string]int{"a": 1, "c": 3}) {
		t.Fatalf("omit = %v", Omit(m, "b"))
	}
	if !reflect.DeepEqual(PickBy(m, func(_ string, v int) bool { return v > 1 }), map[string]int{"b": 2, "c": 3}) {
		t.Fatal("pickby")
	}
	if !reflect.DeepEqual(OmitBy(m, func(_ string, v int) bool { return v > 1 }), map[string]int{"a": 1}) {
		t.Fatal("omitby")
	}
}

func TestMapKeysValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	mv := MapValues(m, func(_ string, v int) int { return v * 10 })
	if mv["a"] != 10 || mv["b"] != 20 {
		t.Fatalf("mapvalues = %v", mv)
	}
	mk := MapKeys(m, func(k string, _ int) string { return k + "!" })
	if mk["a!"] != 1 || mk["b!"] != 2 {
		t.Fatalf("mapkeys = %v", mk)
	}
}

func TestInvertMergeAssign(t *testing.T) {
	inv := Invert(map[string]int{"a": 1, "b": 2})
	if inv[1] != "a" || inv[2] != "b" {
		t.Fatalf("invert = %v", inv)
	}
	merged := Merge(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 20, "c": 3})
	if !reflect.DeepEqual(merged, map[string]int{"a": 1, "b": 20, "c": 3}) {
		t.Fatalf("merge = %v", merged)
	}
	dst := map[string]int{"a": 1}
	Assign(dst, map[string]int{"b": 2}, map[string]int{"a": 9})
	if !reflect.DeepEqual(dst, map[string]int{"a": 9, "b": 2}) {
		t.Fatalf("assign = %v", dst)
	}
}

func TestGetHas(t *testing.T) {
	m := map[string]any{
		"user": map[string]any{
			"name": "Ada",
			"address": map[string]any{
				"city": "London",
			},
		},
	}
	v, ok := Get(m, "user.address.city")
	if !ok || v != "London" {
		t.Fatalf("get nested = %v,%v", v, ok)
	}
	if _, ok := Get(m, "user.missing.city"); ok {
		t.Fatal("get missing")
	}
	if _, ok := Get(m, ""); ok {
		t.Fatal("get empty path")
	}
	if _, ok := Get(m, "user.name.deeper"); ok {
		t.Fatal("get through non-map")
	}
	if !Has(m, "user.name") {
		t.Fatal("has")
	}
	if Has(m, "user.zzz") {
		t.Fatal("has miss")
	}
}
