package lodash_test

import (
	"testing"

	lodash "github.com/malcolmston/lodash"
)

func TestEscapeUnescape(t *testing.T) {
	raw := `Fred, Barney & Pebbles <"'>`
	escaped := lodash.Escape(raw)
	want := `Fred, Barney &amp; Pebbles &lt;&quot;&#39;&gt;`
	if escaped != want {
		t.Errorf("Escape = %q", escaped)
	}
	if lodash.Unescape(escaped) != raw {
		t.Errorf("Unescape roundtrip = %q", lodash.Unescape(escaped))
	}
}

func TestEscapeRegExp(t *testing.T) {
	got := lodash.EscapeRegExp("[lodash](https://lodash.com/)")
	want := `\[lodash\]\(https://lodash\.com/\)`
	if got != want {
		t.Errorf("EscapeRegExp = %q", got)
	}
}

func TestLowerUpperCase(t *testing.T) {
	if lodash.LowerCase("fooBar") != "foo bar" {
		t.Errorf("LowerCase = %q", lodash.LowerCase("fooBar"))
	}
	if lodash.UpperCase("--foo-bar--") != "FOO BAR" {
		t.Errorf("UpperCase = %q", lodash.UpperCase("--foo-bar--"))
	}
	if lodash.ToLower("FOO") != "foo" || lodash.ToUpper("foo") != "FOO" {
		t.Errorf("ToLower/ToUpper")
	}
}

func TestStartsEndsWithReplaceSplit(t *testing.T) {
	if !lodash.StartsWith("abcdef", "abc") || lodash.StartsWith("abc", "xyz") {
		t.Errorf("StartsWith")
	}
	if !lodash.EndsWith("abcdef", "def") || lodash.EndsWith("abc", "xyz") {
		t.Errorf("EndsWith")
	}
	if lodash.Replace("aXbXc", "X", "-") != "a-bXc" {
		t.Errorf("Replace")
	}
	if !lodash.IsEqual(lodash.Split("a-b-c-d", "-", 2), []string{"a", "b"}) {
		t.Errorf("Split limit")
	}
	if !lodash.IsEqual(lodash.Split("a-b-c", "-", 0), []string{"a", "b", "c"}) {
		t.Errorf("Split all")
	}
}

func TestTemplate(t *testing.T) {
	render := lodash.Template("Hello <%= name %>, score <%= stats.score %>!")
	got := render(map[string]any{"name": "Bob", "stats": map[string]any{"score": 42}})
	if got != "Hello Bob, score 42!" {
		t.Errorf("Template = %q", got)
	}
	escRender := lodash.Template("<%- html %>")
	if escRender(map[string]any{"html": "<b>"}) != "&lt;b&gt;" {
		t.Errorf("Template escape")
	}
	evalRender := lodash.Template("a<% ignored %>b")
	if evalRender(map[string]any{}) != "ab" {
		t.Errorf("Template evaluate should be stripped")
	}
	custom := lodash.Template("{{= x }}", lodash.TemplateOptions{Open: "{{", Close: "}}"})
	if custom(map[string]any{"x": 7}) != "7" {
		t.Errorf("Template custom delimiters")
	}
	// Missing key renders nothing.
	if lodash.Template("<%= missing %>")(map[string]any{}) != "" {
		t.Errorf("Template missing key")
	}
}
