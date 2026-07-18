package lodash

import (
	"reflect"
	"testing"
)

// The vectors in this file are transcribed verbatim from lodash's own test
// suite (lodash/lodash test/test.js at v4.17.21). Each assertion pins this port
// to the exact output the upstream QUnit tests assert.

// TestParityCamelCase mirrors lodash.camelCase assertions.
func TestParityCamelCase(t *testing.T) {
	cases := map[string]string{
		"12 feet":           "12Feet",
		"enable 6h format":  "enable6HFormat",
		"enable 24H format": "enable24HFormat",
		"too legit 2 quit":  "tooLegit2Quit",
		"walk 500 miles":    "walk500Miles",
		"xhr2 request":      "xhr2Request",
	}
	for in, want := range cases {
		if got := CamelCase(in); got != want {
			t.Errorf("CamelCase(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestParityCapitalize mirrors lodash.capitalize assertions.
func TestParityCapitalize(t *testing.T) {
	cases := map[string]string{"fred": "Fred", "Fred": "Fred", " fred": " fred"}
	for in, want := range cases {
		if got := Capitalize(in); got != want {
			t.Errorf("Capitalize(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestParityUpperLowerFirst mirrors lodash.upperFirst / lodash.lowerFirst.
func TestParityUpperLowerFirst(t *testing.T) {
	upper := map[string]string{"fred": "Fred", "Fred": "Fred", "FRED": "FRED"}
	for in, want := range upper {
		if got := UpperFirst(in); got != want {
			t.Errorf("UpperFirst(%q) = %q, want %q", in, got, want)
		}
	}
	lower := map[string]string{"fred": "fred", "Fred": "fred", "FRED": "fRED"}
	for in, want := range lower {
		if got := LowerFirst(in); got != want {
			t.Errorf("LowerFirst(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestParityStartCase mirrors lodash.startCase, which preserves interior letter
// case (it upper-cases only the first letter of each word).
func TestParityStartCase(t *testing.T) {
	cases := map[string]string{
		"--foo-bar--": "Foo Bar",
		"fooBar":      "Foo Bar",
		"__FOO_BAR__": "FOO BAR",
	}
	for in, want := range cases {
		if got := StartCase(in); got != want {
			t.Errorf("StartCase(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestParityWords mirrors lodash.words single-argument assertions.
func TestParityWords(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"12ft", []string{"12", "ft"}},
		{"aeiouAreVowels", []string{"aeiou", "Are", "Vowels"}},
		{"enable 6h format", []string{"enable", "6", "h", "format"}},
		{"enable 24H format", []string{"enable", "24", "H", "format"}},
		{"isISO8601", []string{"is", "ISO", "8601"}},
		{"LETTERSAeiouAreVowels", []string{"LETTERS", "Aeiou", "Are", "Vowels"}},
		{"tooLegit2Quit", []string{"too", "Legit", "2", "Quit"}},
		{"walk500Miles", []string{"walk", "500", "Miles"}},
		{"xhr2Request", []string{"xhr", "2", "Request"}},
		{"XMLHttp", []string{"XML", "Http"}},
		{"XmlHTTP", []string{"Xml", "HTTP"}},
		{"XmlHttp", []string{"Xml", "Http"}},
	}
	for _, c := range cases {
		if got := Words(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Words(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

// TestParityEscapeUnescape mirrors lodash.escape / lodash.unescape.
func TestParityEscapeUnescape(t *testing.T) {
	unescaped := `&<>"'`
	escaped := "&amp;&lt;&gt;&quot;&#39;"
	if got := Escape(unescaped); got != escaped {
		t.Errorf("Escape(%q) = %q, want %q", unescaped, got, escaped)
	}
	if got := Escape("abc"); got != "abc" {
		t.Errorf("Escape(abc) = %q", got)
	}
	if got := Unescape("&amp;lt;"); got != "&lt;" {
		t.Errorf("Unescape(&amp;lt;) = %q, want &lt;", got)
	}
	if got := Unescape(escaped); got != unescaped {
		t.Errorf("Unescape(%q) = %q, want %q", escaped, got, unescaped)
	}
	if got := Unescape("abc"); got != "abc" {
		t.Errorf("Unescape(abc) = %q", got)
	}
}

// TestParityEscapeRegExp mirrors lodash.escapeRegExp.
func TestParityEscapeRegExp(t *testing.T) {
	unescaped := `^$.*+?()[]{}|\`
	escaped := `\^\$\.\*\+\?\(\)\[\]\{\}\|\\`
	if got := EscapeRegExp(unescaped); got != escaped {
		t.Errorf("EscapeRegExp(%q) = %q, want %q", unescaped, got, escaped)
	}
	if got := EscapeRegExp("abc"); got != "abc" {
		t.Errorf("EscapeRegExp(abc) = %q", got)
	}
}

// TestParityRepeat mirrors lodash.repeat.
func TestParityRepeat(t *testing.T) {
	if got := Repeat("*", 3); got != "***" {
		t.Errorf("Repeat(*,3) = %q", got)
	}
	if got := Repeat("abc", 2); got != "abcabc" {
		t.Errorf("Repeat(abc,2) = %q", got)
	}
	if got := Repeat("abc", 0); got != "" {
		t.Errorf("Repeat(abc,0) = %q", got)
	}
	if got := Repeat("abc", -2); got != "" {
		t.Errorf("Repeat(abc,-2) = %q", got)
	}
}

// TestParityPad mirrors lodash.pad / padStart / padEnd.
func TestParityPad(t *testing.T) {
	if got := Pad("abc", 8, " "); got != "  abc   " {
		t.Errorf("Pad(abc,8) = %q, want %q", got, "  abc   ")
	}
	if got := Pad("abc", 8, "_-"); got != "_-abc_-_" {
		t.Errorf("Pad(abc,8,_-) = %q, want %q", got, "_-abc_-_")
	}
	if got := PadStart("abc", 6, "_-"); got != "_-_abc" {
		t.Errorf("PadStart(abc,6,_-) = %q, want %q", got, "_-_abc")
	}
	if got := PadEnd("abc", 6, "_-"); got != "abc_-_" {
		t.Errorf("PadEnd(abc,6,_-) = %q, want %q", got, "abc_-_")
	}
}

// TestParityTruncate mirrors lodash.truncate (length + omission form).
func TestParityTruncate(t *testing.T) {
	s := "hi-diddly-ho there, neighborino"
	if got := Truncate(s, 30, "..."); got != "hi-diddly-ho there, neighbo..." {
		t.Errorf("Truncate default = %q", got)
	}
	if got := Truncate(s, len(s), "..."); got != s {
		t.Errorf("Truncate full-length = %q", got)
	}
	if got := Truncate(s, 24, "..."); got != "hi-diddly-ho there, n..." {
		t.Errorf("Truncate(24) = %q", got)
	}
	if got := Truncate(s, 4, "..."); got != "h..." {
		t.Errorf("Truncate(4) = %q", got)
	}
}

// TestParityDeburr mirrors lodash.deburr: it converts the Latin-1 Supplement and
// Latin Extended-A blocks to basic Latin, expands ligatures, and drops combining
// diacritical marks while leaving mathematical operators untouched.
func TestParityDeburr(t *testing.T) {
	burred := "ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõöøùúûüýþÿĀāĂăĄąĆćĈĉĊċČčĎďĐđĒēĔĕĖėĘęĚěĜĝĞğĠġĢģĤĥĦħĨĩĪīĬĭĮįİıĲĳĴĵĶķĸĹĺĻļĽľĿŀŁłŃńŅņŇňŉŊŋŌōŎŏŐőŒœŔŕŖŗŘřŚśŜŝŞşŠšŢţŤťŦŧŨũŪūŬŭŮůŰűŲųŴŵŶŷŸŹźŻżŽžſ"
	deburred := []string{"A", "A", "A", "A", "A", "A", "Ae", "C", "E", "E", "E", "E", "I", "I", "I", "I", "D", "N", "O", "O", "O", "O", "O", "O", "U", "U", "U", "U", "Y", "Th", "ss", "a", "a", "a", "a", "a", "a", "ae", "c", "e", "e", "e", "e", "i", "i", "i", "i", "d", "n", "o", "o", "o", "o", "o", "o", "u", "u", "u", "u", "y", "th", "y", "A", "a", "A", "a", "A", "a", "C", "c", "C", "c", "C", "c", "C", "c", "D", "d", "D", "d", "E", "e", "E", "e", "E", "e", "E", "e", "E", "e", "G", "g", "G", "g", "G", "g", "G", "g", "H", "h", "H", "h", "I", "i", "I", "i", "I", "i", "I", "i", "I", "i", "IJ", "ij", "J", "j", "K", "k", "k", "L", "l", "L", "l", "L", "l", "L", "l", "L", "l", "N", "n", "N", "n", "N", "n", "'n", "N", "n", "O", "o", "O", "o", "O", "o", "Oe", "oe", "R", "r", "R", "r", "R", "r", "S", "s", "S", "s", "S", "s", "S", "s", "T", "t", "T", "t", "T", "t", "U", "u", "U", "u", "U", "u", "U", "u", "U", "u", "U", "u", "W", "w", "Y", "y", "Y", "Z", "z", "Z", "z", "Z", "z", "s"}
	br := []rune(burred)
	if len(br) != len(deburred) {
		t.Fatalf("vector length mismatch: %d vs %d", len(br), len(deburred))
	}
	for i, r := range br {
		if got := Deburr(string(r)); got != deburred[i] {
			t.Errorf("Deburr(%q) = %q, want %q", string(r), got, deburred[i])
		}
	}
	for _, op := range []string{"×", "÷"} {
		if got := Deburr(op); got != op {
			t.Errorf("Deburr(%q) = %q, want unchanged", op, got)
		}
	}
	for _, m := range []rune{0x0300, 0x0301, 0x0308, 0x030a, 0x0327, 0xfe20, 0xfe23} {
		in := "e" + string(m) + "i"
		if got := Deburr(in); got != "ei" {
			t.Errorf("Deburr(%q) = %q, want ei", in, got)
		}
	}
}
