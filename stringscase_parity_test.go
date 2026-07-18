package lodash

import "testing"

// Vectors transcribed from lodash test/test.js (v4.17.21) "case methods" module,
// which runs every burred letter and every contraction through the case
// compounders. These lock in the lodash behavior that the compounders deburr
// Latin letters and strip contraction apostrophes before splitting words, which
// the plain Words function deliberately does not do.

// TestParityCaseMethodsShared mirrors the "should convert `string` to <case>"
// assertions over the shared input set. For camel/kebab/lower/snake all inputs
// converge on the same value; startCase preserves the all-caps "FOO BAR".
func TestParityCaseMethodsShared(t *testing.T) {
	inputs := []string{
		"foo bar", "Foo bar", "foo Bar", "Foo Bar",
		"FOO BAR", "fooBar", "--foo-bar--", "__foo_bar__",
	}
	for _, in := range inputs {
		if got := CamelCase(in); got != "fooBar" {
			t.Errorf("CamelCase(%q) = %q, want fooBar", in, got)
		}
		if got := KebabCase(in); got != "foo-bar" {
			t.Errorf("KebabCase(%q) = %q, want foo-bar", in, got)
		}
		if got := LowerCase(in); got != "foo bar" {
			t.Errorf("LowerCase(%q) = %q, want 'foo bar'", in, got)
		}
		if got := SnakeCase(in); got != "foo_bar" {
			t.Errorf("SnakeCase(%q) = %q, want foo_bar", in, got)
		}
		if got := UpperCase(in); got != "FOO BAR" {
			t.Errorf("UpperCase(%q) = %q, want 'FOO BAR'", in, got)
		}
		wantStart := "Foo Bar"
		if in == "FOO BAR" {
			wantStart = "FOO BAR" // startCase upper-cases only the first letter.
		}
		if got := StartCase(in); got != wantStart {
			t.Errorf("StartCase(%q) = %q, want %q", in, got, wantStart)
		}
	}
}

// TestParityCaseMethodsDeburr mirrors the "should deburr letters" assertions:
// each compounder converts an accented letter to its basic-Latin form. Spot
// checks with 'À' (-> A) and the ligatures 'Æ' (-> Ae) and 'ß' (-> ss).
func TestParityCaseMethodsDeburr(t *testing.T) {
	if got := CamelCase("À"); got != "a" {
		t.Errorf("CamelCase(À) = %q, want a", got)
	}
	if got := SnakeCase("À"); got != "a" {
		t.Errorf("SnakeCase(À) = %q, want a", got)
	}
	if got := UpperCase("À"); got != "A" {
		t.Errorf("UpperCase(À) = %q, want A", got)
	}
	if got := LowerCase("Æ"); got != "ae" {
		t.Errorf("LowerCase(Æ) = %q, want ae", got)
	}
	if got := UpperCase("ß"); got != "SS" {
		t.Errorf("UpperCase(ß) = %q, want SS", got)
	}
	// startCase of the Latin-1 ligature preserves interior letters (upperFirst).
	if got := StartCase("Ĳ"); got != "IJ" {
		t.Errorf("StartCase(Ĳ) = %q, want IJ", got)
	}
}

// TestParityCaseMethodsApostrophe mirrors the "should remove contraction
// apostrophes" assertions for the ASCII apostrophe and the right single quote,
// over lodash's postfix set.
func TestParityCaseMethodsApostrophe(t *testing.T) {
	postfixes := []string{"d", "ll", "m", "re", "s", "t", "ve"}
	for _, apos := range []string{"'", "’"} {
		for _, p := range postfixes {
			in := "a b" + apos + p + " c"
			if got := CamelCase(in); got != "aB"+p+"C" {
				t.Errorf("CamelCase(%q) = %q, want %q", in, got, "aB"+p+"C")
			}
			if got := KebabCase(in); got != "a-b"+p+"-c" {
				t.Errorf("KebabCase(%q) = %q, want %q", in, got, "a-b"+p+"-c")
			}
			if got := SnakeCase(in); got != "a_b"+p+"_c" {
				t.Errorf("SnakeCase(%q) = %q, want %q", in, got, "a_b"+p+"_c")
			}
		}
	}
}

// TestParityCaseMethodsOperators mirrors the "should remove Latin mathematical
// operators" assertions: '×' and '÷' produce empty output.
func TestParityCaseMethodsOperators(t *testing.T) {
	for _, op := range []string{"×", "÷"} {
		if got := CamelCase(op); got != "" {
			t.Errorf("CamelCase(%q) = %q, want empty", op, got)
		}
		if got := SnakeCase(op); got != "" {
			t.Errorf("SnakeCase(%q) = %q, want empty", op, got)
		}
	}
}
