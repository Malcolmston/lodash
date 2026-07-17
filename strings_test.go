package lodash

import (
	"reflect"
	"testing"
)

func TestWords(t *testing.T) {
	cases := map[string][]string{
		"fooBar":         {"foo", "Bar"},
		"foo_bar-baz":    {"foo", "bar", "baz"},
		"HTMLParser":     {"HTML", "Parser"},
		"  hello world ": {"hello", "world"},
		"version2Point0": {"version", "2", "Point", "0"},
		"":               nil,
	}
	for in, want := range cases {
		if got := Words(in); !reflect.DeepEqual(got, want) {
			t.Fatalf("Words(%q) = %v, want %v", in, got, want)
		}
	}
}

func TestCaseConversions(t *testing.T) {
	if CamelCase("Foo Bar-baz") != "fooBarBaz" {
		t.Fatalf("camel = %q", CamelCase("Foo Bar-baz"))
	}
	if PascalCase("foo_bar") != "FooBar" {
		t.Fatalf("pascal = %q", PascalCase("foo_bar"))
	}
	if SnakeCase("fooBarBaz") != "foo_bar_baz" {
		t.Fatalf("snake = %q", SnakeCase("fooBarBaz"))
	}
	if KebabCase("Foo Bar") != "foo-bar" {
		t.Fatalf("kebab = %q", KebabCase("Foo Bar"))
	}
	if StartCase("foo_bar-baz") != "Foo Bar Baz" {
		t.Fatalf("start = %q", StartCase("foo_bar-baz"))
	}
}

func TestCapitalizeFirst(t *testing.T) {
	if Capitalize("hELLO") != "Hello" {
		t.Fatalf("capitalize = %q", Capitalize("hELLO"))
	}
	if UpperFirst("foo") != "Foo" {
		t.Fatal("upperfirst")
	}
	if LowerFirst("FOO") != "fOO" {
		t.Fatal("lowerfirst")
	}
	if Capitalize("") != "" || UpperFirst("") != "" || LowerFirst("") != "" {
		t.Fatal("empty handling")
	}
}

func TestPad(t *testing.T) {
	if PadStart("5", 3, "0") != "005" {
		t.Fatalf("padstart = %q", PadStart("5", 3, "0"))
	}
	if PadEnd("5", 3, "0") != "500" {
		t.Fatalf("padend = %q", PadEnd("5", 3, "0"))
	}
	if Pad("ab", 6, "-") != "--ab--" {
		t.Fatalf("pad = %q", Pad("ab", 6, "-"))
	}
	if Pad("abc", 5, "*") != "*abc*" {
		t.Fatalf("pad odd = %q", Pad("abc", 5, "*"))
	}
	if Pad("toolong", 3, "-") != "toolong" {
		t.Fatal("pad noop")
	}
	if PadStart("x", 5, "ab") != "ababx" {
		t.Fatalf("padstart multi = %q", PadStart("x", 5, "ab"))
	}
	if PadStart("x", 5, "") != "x" {
		t.Fatal("pad empty chars")
	}
}

func TestTruncateRepeat(t *testing.T) {
	if Truncate("hello world", 8, "...") != "hello..." {
		t.Fatalf("truncate = %q", Truncate("hello world", 8, "..."))
	}
	if Truncate("hi", 8, "...") != "hi" {
		t.Fatal("truncate noop")
	}
	if Truncate("hello", 2, "...") != "..." {
		t.Fatalf("truncate short = %q", Truncate("hello", 2, "..."))
	}
	if Repeat("ab", 3) != "ababab" {
		t.Fatal("repeat")
	}
	if Repeat("ab", 0) != "" {
		t.Fatal("repeat zero")
	}
}

func TestDeburrTrim(t *testing.T) {
	if Deburr("déjà vu Ñ") != "deja vu N" {
		t.Fatalf("deburr = %q", Deburr("déjà vu Ñ"))
	}
	if Trim("  hi  ", "") != "hi" {
		t.Fatal("trim ws")
	}
	if Trim("xxhixx", "x") != "hi" {
		t.Fatal("trim cutset")
	}
	if TrimStart("__hi__", "_") != "hi__" {
		t.Fatal("trimstart")
	}
	if TrimEnd("__hi__", "_") != "__hi" {
		t.Fatal("trimend")
	}
	if TrimStart("  hi", "") != "hi" {
		t.Fatal("trimstart ws")
	}
	if TrimEnd("hi  ", "") != "hi" {
		t.Fatal("trimend ws")
	}
}
