package lodash

import "strings"

// htmlEscapes maps HTML-sensitive characters to their entity encodings, matching
// lodash's _.escape.
var htmlEscapes = []struct{ from, to string }{
	{"&", "&amp;"},
	{"<", "&lt;"},
	{">", "&gt;"},
	{`"`, "&quot;"},
	{"'", "&#39;"},
}

// Escape converts the characters "&", "<", ">", '"' and "'" in s to their
// corresponding HTML entities.
func Escape(s string) string {
	// Ampersand must be replaced first so entity ampersands are not re-escaped.
	for _, e := range htmlEscapes {
		s = strings.ReplaceAll(s, e.from, e.to)
	}
	return s
}

// Unescape reverses Escape, converting the HTML entities for "&", "<", ">", '"'
// and "'" back into their characters.
func Unescape(s string) string {
	// Reverse order so "&amp;" is restored last.
	for i := len(htmlEscapes) - 1; i >= 0; i-- {
		s = strings.ReplaceAll(s, htmlEscapes[i].to, htmlEscapes[i].from)
	}
	return s
}

// EscapeRegExp escapes the characters that have special meaning in regular
// expressions so that s can be matched literally.
func EscapeRegExp(s string) string {
	const special = `\^$.*+?()[]{}|`
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if strings.ContainsRune(special, r) {
			b.WriteByte('\\')
		}
		b.WriteRune(r)
	}
	return b.String()
}

// LowerCase converts s to space-separated lower case words, splitting on case,
// digit and separator boundaries. Like lodash's _.lowerCase it deburrs Latin
// letters and removes contraction apostrophes before splitting (see
// compoundWords), so e.g. LowerCase("À") yields "a".
func LowerCase(s string) string {
	return joinCase(compoundWords(s), " ", strings.ToLower)
}

// UpperCase converts s to space-separated UPPER CASE words. Like lodash's
// _.upperCase it deburrs Latin letters and removes contraction apostrophes
// before splitting (see compoundWords), so e.g. UpperCase("À") yields "A".
func UpperCase(s string) string {
	return joinCase(compoundWords(s), " ", strings.ToUpper)
}

func joinCase(words []string, sep string, transform func(string) string) string {
	out := make([]string, len(words))
	for i, w := range words {
		out[i] = transform(w)
	}
	return strings.Join(out, sep)
}

// ToLower returns s with all Unicode letters mapped to their lower case,
// without any word splitting.
func ToLower(s string) string { return strings.ToLower(s) }

// ToUpper returns s with all Unicode letters mapped to their upper case,
// without any word splitting.
func ToUpper(s string) string { return strings.ToUpper(s) }

// StartsWith reports whether s begins with prefix.
func StartsWith(s, prefix string) bool { return strings.HasPrefix(s, prefix) }

// EndsWith reports whether s ends with suffix.
func EndsWith(s, suffix string) bool { return strings.HasSuffix(s, suffix) }

// Replace returns s with the first non-overlapping occurrence of old replaced by
// replacement.
func Replace(s, old, replacement string) string {
	return strings.Replace(s, old, replacement, 1)
}

// Split divides s around each instance of sep. When limit is greater than zero
// at most limit substrings are returned (the final one holding the unsplit
// remainder is dropped beyond the limit); a non-positive limit returns all
// substrings.
func Split(s, sep string, limit int) []string {
	parts := strings.Split(s, sep)
	if limit > 0 && limit < len(parts) {
		parts = parts[:limit]
	}
	return parts
}

// TemplateOptions configures the delimiters recognized by Template. Zero-value
// fields fall back to the lodash defaults: "<%=" ... "%>" for interpolation,
// "<%-" ... "%>" for HTML-escaped interpolation and "<%" ... "%>" for evaluate
// blocks.
type TemplateOptions struct {
	// Open is the opening delimiter (default "<%").
	Open string
	// Close is the closing delimiter (default "%>").
	Close string
}

// Template compiles a text template into a renderer function. Within the
// template, "<%= path %>" interpolates the value at path from the data map,
// "<%- path %>" interpolates it HTML-escaped, and "<% ... %>" evaluate blocks
// are recognized and removed (they are not executed, since no scripting engine
// is embedded). Paths are resolved with Get, so dotted and bracket notation are
// supported.
func Template(tmpl string, opts ...TemplateOptions) func(data map[string]any) string {
	open, closeD := "<%", "%>"
	if len(opts) > 0 {
		if opts[0].Open != "" {
			open = opts[0].Open
		}
		if opts[0].Close != "" {
			closeD = opts[0].Close
		}
	}
	type segment struct {
		literal string
		kind    byte // 0 literal, '=' interpolate, '-' escape, 'e' evaluate
		path    string
	}
	var segs []segment
	rest := tmpl
	for {
		i := strings.Index(rest, open)
		if i < 0 {
			segs = append(segs, segment{literal: rest})
			break
		}
		if i > 0 {
			segs = append(segs, segment{literal: rest[:i]})
		}
		rest = rest[i+len(open):]
		j := strings.Index(rest, closeD)
		if j < 0 {
			// Unterminated block: treat the remainder as literal.
			segs = append(segs, segment{literal: open + rest})
			break
		}
		body := rest[:j]
		rest = rest[j+len(closeD):]
		kind := byte('e')
		if len(body) > 0 && (body[0] == '=' || body[0] == '-') {
			kind = body[0]
			body = body[1:]
		}
		segs = append(segs, segment{kind: kind, path: strings.TrimSpace(body)})
	}
	return func(data map[string]any) string {
		var b strings.Builder
		for _, s := range segs {
			switch s.kind {
			case '=':
				if v, ok := Get(data, s.path); ok {
					b.WriteString(ToString(v))
				}
			case '-':
				if v, ok := Get(data, s.path); ok {
					b.WriteString(Escape(ToString(v)))
				}
			case 'e':
				// Evaluate blocks are not executed.
			default:
				b.WriteString(s.literal)
			}
		}
		return b.String()
	}
}
