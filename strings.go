package lodash

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Words splits a string into its component words. It understands spaces,
// punctuation, snake_case, kebab-case and camelCase / PascalCase boundaries as
// well as runs of digits.
func Words(s string) []string {
	var words []string
	runes := []rune(s)
	n := len(runes)
	i := 0
	for i < n {
		// Skip any separator that is not a letter or digit.
		if !isWordRune(runes[i]) {
			i++
			continue
		}
		start := i
		for i < n && isWordRune(runes[i]) {
			// Break on a lower-to-upper transition (camelCase).
			if i > start && unicode.IsUpper(runes[i]) && unicode.IsLower(runes[i-1]) {
				break
			}
			// Break before the last uppercase of an acronym run that is
			// followed by a lowercase letter (e.g. "HTMLParser" -> HTML, Parser).
			if i > start && unicode.IsUpper(runes[i-1]) && unicode.IsUpper(runes[i]) &&
				i+1 < n && unicode.IsLower(runes[i+1]) {
				break
			}
			// Break on letter-to-digit and digit-to-letter transitions.
			if i > start {
				prevDigit := unicode.IsDigit(runes[i-1])
				curDigit := unicode.IsDigit(runes[i])
				if prevDigit != curDigit {
					break
				}
			}
			i++
		}
		words = append(words, string(runes[start:i]))
	}
	return words
}

func isWordRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// CamelCase converts a string to camelCase.
func CamelCase(s string) string {
	words := Words(s)
	var b strings.Builder
	for i, w := range words {
		if i == 0 {
			b.WriteString(strings.ToLower(w))
		} else {
			b.WriteString(capitalizeWord(w))
		}
	}
	return b.String()
}

// PascalCase converts a string to PascalCase (a.k.a. UpperCamelCase).
func PascalCase(s string) string {
	words := Words(s)
	var b strings.Builder
	for _, w := range words {
		b.WriteString(capitalizeWord(w))
	}
	return b.String()
}

// SnakeCase converts a string to snake_case.
func SnakeCase(s string) string {
	return joinLower(Words(s), "_")
}

// KebabCase converts a string to kebab-case.
func KebabCase(s string) string {
	return joinLower(Words(s), "-")
}

// StartCase converts a string to Start Case, capitalizing the first letter of
// each word and separating words with single spaces.
func StartCase(s string) string {
	words := Words(s)
	for i, w := range words {
		words[i] = capitalizeWord(w)
	}
	return strings.Join(words, " ")
}

func joinLower(words []string, sep string) string {
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return strings.Join(words, sep)
}

func capitalizeWord(w string) string {
	if w == "" {
		return w
	}
	r := []rune(w)
	r[0] = unicode.ToUpper(r[0])
	for i := 1; i < len(r); i++ {
		r[i] = unicode.ToLower(r[i])
	}
	return string(r)
}

// Capitalize upper-cases the first character of the string and lower-cases the
// rest.
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	for i := 1; i < len(r); i++ {
		r[i] = unicode.ToLower(r[i])
	}
	return string(r)
}

// UpperFirst upper-cases the first character of the string, leaving the rest
// unchanged.
func UpperFirst(s string) string {
	if s == "" {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}

// LowerFirst lower-cases the first character of the string, leaving the rest
// unchanged.
func LowerFirst(s string) string {
	if s == "" {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[size:]
}

// Pad pads a string on both sides with chars until it reaches length. If the
// string is already at least length long it is returned unchanged.
func Pad(s string, length int, chars string) string {
	strLen := utf8.RuneCountInString(s)
	if strLen >= length || chars == "" {
		return s
	}
	total := length - strLen
	left := total / 2
	right := total - left
	return makePadding(chars, left) + s + makePadding(chars, right)
}

// PadStart pads a string on the left with chars until it reaches length.
func PadStart(s string, length int, chars string) string {
	strLen := utf8.RuneCountInString(s)
	if strLen >= length || chars == "" {
		return s
	}
	return makePadding(chars, length-strLen) + s
}

// PadEnd pads a string on the right with chars until it reaches length.
func PadEnd(s string, length int, chars string) string {
	strLen := utf8.RuneCountInString(s)
	if strLen >= length || chars == "" {
		return s
	}
	return s + makePadding(chars, length-strLen)
}

// makePadding builds a padding string of exactly n runes by repeating chars.
func makePadding(chars string, n int) string {
	if n <= 0 || chars == "" {
		return ""
	}
	padRunes := []rune(chars)
	r := make([]rune, n)
	for i := 0; i < n; i++ {
		r[i] = padRunes[i%len(padRunes)]
	}
	return string(r)
}

// Truncate shortens a string to at most length runes, appending omission (which
// counts toward the length) when truncation occurs.
func Truncate(s string, length int, omission string) string {
	if utf8.RuneCountInString(s) <= length {
		return s
	}
	omitLen := utf8.RuneCountInString(omission)
	keep := length - omitLen
	if keep < 0 {
		keep = 0
	}
	r := []rune(s)
	return string(r[:keep]) + omission
}

// Repeat returns a string consisting of n copies of s.
func Repeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(s, n)
}

// deburrMap maps common accented Latin characters to their ASCII equivalents.
var deburrMap = map[rune]rune{
	'à': 'a', 'á': 'a', 'â': 'a', 'ã': 'a', 'ä': 'a', 'å': 'a',
	'ç': 'c', 'è': 'e', 'é': 'e', 'ê': 'e', 'ë': 'e',
	'ì': 'i', 'í': 'i', 'î': 'i', 'ï': 'i',
	'ñ': 'n', 'ò': 'o', 'ó': 'o', 'ô': 'o', 'õ': 'o', 'ö': 'o',
	'ù': 'u', 'ú': 'u', 'û': 'u', 'ü': 'u', 'ý': 'y', 'ÿ': 'y',
	'À': 'A', 'Á': 'A', 'Â': 'A', 'Ã': 'A', 'Ä': 'A', 'Å': 'A',
	'Ç': 'C', 'È': 'E', 'É': 'E', 'Ê': 'E', 'Ë': 'E',
	'Ì': 'I', 'Í': 'I', 'Î': 'I', 'Ï': 'I',
	'Ñ': 'N', 'Ò': 'O', 'Ó': 'O', 'Ô': 'O', 'Õ': 'O', 'Ö': 'O',
	'Ù': 'U', 'Ú': 'U', 'Û': 'U', 'Ü': 'U', 'Ý': 'Y',
}

// Deburr replaces Latin-1 accented characters with their ASCII counterparts.
func Deburr(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if repl, ok := deburrMap[r]; ok {
			b.WriteRune(repl)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// Trim removes leading and trailing characters found in cutset from s. When
// cutset is empty it trims whitespace.
func Trim(s, cutset string) string {
	if cutset == "" {
		return strings.TrimSpace(s)
	}
	return strings.Trim(s, cutset)
}

// TrimStart removes leading characters found in cutset from s. When cutset is
// empty it trims leading whitespace.
func TrimStart(s, cutset string) string {
	if cutset == "" {
		return strings.TrimLeftFunc(s, unicode.IsSpace)
	}
	return strings.TrimLeft(s, cutset)
}

// TrimEnd removes trailing characters found in cutset from s. When cutset is
// empty it trims trailing whitespace.
func TrimEnd(s, cutset string) string {
	if cutset == "" {
		return strings.TrimRightFunc(s, unicode.IsSpace)
	}
	return strings.TrimRight(s, cutset)
}
