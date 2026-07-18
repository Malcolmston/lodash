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

// apostropheReplacer strips the contraction apostrophes lodash removes before
// splitting a string into words for the case compounders (so "a b'd c" becomes
// "a bd c"). It covers both the ASCII apostrophe and the right single quote.
var apostropheReplacer = strings.NewReplacer("'", "", "’", "")

// compoundWords is the shared front-end for the case compounders (CamelCase,
// PascalCase, SnakeCase, KebabCase, StartCase, LowerCase and UpperCase). Like
// lodash's createCompounder, it deburrs Latin letters and removes contraction
// apostrophes before splitting the string into words, so e.g. CamelCase("déjà")
// yields "deja". The plain Words function does neither, matching lodash's
// _.words.
func compoundWords(s string) []string {
	return Words(apostropheReplacer.Replace(Deburr(s)))
}

// CamelCase converts a string to camelCase.
func CamelCase(s string) string {
	words := compoundWords(s)
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
	words := compoundWords(s)
	var b strings.Builder
	for _, w := range words {
		b.WriteString(capitalizeWord(w))
	}
	return b.String()
}

// SnakeCase converts a string to snake_case.
func SnakeCase(s string) string {
	return joinLower(compoundWords(s), "_")
}

// KebabCase converts a string to kebab-case.
func KebabCase(s string) string {
	return joinLower(compoundWords(s), "-")
}

// StartCase converts a string to Start Case, upper-casing the first letter of
// each word and separating words with single spaces. Unlike Capitalize it does
// not lower-case the remaining letters, so acronyms such as "FOO" are
// preserved. This mirrors lodash's _.startCase, which composes upperFirst over
// each detected word (e.g. "__FOO_BAR__" -> "FOO BAR").
func StartCase(s string) string {
	words := compoundWords(s)
	for i, w := range words {
		words[i] = UpperFirst(w)
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

// deburrMap maps Latin-1 Supplement and Latin Extended-A letters to their basic
// Latin equivalents. It reproduces lodash's _.deburr conversion table exactly,
// including multi-rune expansions such as 'Æ' -> "Ae", 'Þ' -> "Th" and
// 'ß' -> "ss".
var deburrMap = map[rune]string{
	// Latin-1 Supplement block.
	'À': "A", 'Á': "A", 'Â': "A", 'Ã': "A", 'Ä': "A", 'Å': "A",
	'à': "a", 'á': "a", 'â': "a", 'ã': "a", 'ä': "a", 'å': "a",
	'Ç': "C", 'ç': "c",
	'Ð': "D", 'ð': "d",
	'È': "E", 'É': "E", 'Ê': "E", 'Ë': "E",
	'è': "e", 'é': "e", 'ê': "e", 'ë': "e",
	'Ì': "I", 'Í': "I", 'Î': "I", 'Ï': "I",
	'ì': "i", 'í': "i", 'î': "i", 'ï': "i",
	'Ñ': "N", 'ñ': "n",
	'Ò': "O", 'Ó': "O", 'Ô': "O", 'Õ': "O", 'Ö': "O", 'Ø': "O",
	'ò': "o", 'ó': "o", 'ô': "o", 'õ': "o", 'ö': "o", 'ø': "o",
	'Ù': "U", 'Ú': "U", 'Û': "U", 'Ü': "U",
	'ù': "u", 'ú': "u", 'û': "u", 'ü': "u",
	'Ý': "Y", 'ý': "y", 'ÿ': "y",
	'Æ': "Ae", 'æ': "ae",
	'Þ': "Th", 'þ': "th",
	'ß': "ss",
	// Latin Extended-A block.
	'Ā': "A", 'Ă': "A", 'Ą': "A",
	'ā': "a", 'ă': "a", 'ą': "a",
	'Ć': "C", 'Ĉ': "C", 'Ċ': "C", 'Č': "C",
	'ć': "c", 'ĉ': "c", 'ċ': "c", 'č': "c",
	'Ď': "D", 'Đ': "D", 'ď': "d", 'đ': "d",
	'Ē': "E", 'Ĕ': "E", 'Ė': "E", 'Ę': "E", 'Ě': "E",
	'ē': "e", 'ĕ': "e", 'ė': "e", 'ę': "e", 'ě': "e",
	'Ĝ': "G", 'Ğ': "G", 'Ġ': "G", 'Ģ': "G",
	'ĝ': "g", 'ğ': "g", 'ġ': "g", 'ģ': "g",
	'Ĥ': "H", 'Ħ': "H", 'ĥ': "h", 'ħ': "h",
	'Ĩ': "I", 'Ī': "I", 'Ĭ': "I", 'Į': "I", 'İ': "I",
	'ĩ': "i", 'ī': "i", 'ĭ': "i", 'į': "i", 'ı': "i",
	'Ĵ': "J", 'ĵ': "j",
	'Ķ': "K", 'ķ': "k", 'ĸ': "k",
	'Ĺ': "L", 'Ļ': "L", 'Ľ': "L", 'Ŀ': "L", 'Ł': "L",
	'ĺ': "l", 'ļ': "l", 'ľ': "l", 'ŀ': "l", 'ł': "l",
	'Ń': "N", 'Ņ': "N", 'Ň': "N", 'Ŋ': "N",
	'ń': "n", 'ņ': "n", 'ň': "n", 'ŋ': "n",
	'Ō': "O", 'Ŏ': "O", 'Ő': "O",
	'ō': "o", 'ŏ': "o", 'ő': "o",
	'Ŕ': "R", 'Ŗ': "R", 'Ř': "R",
	'ŕ': "r", 'ŗ': "r", 'ř': "r",
	'Ś': "S", 'Ŝ': "S", 'Ş': "S", 'Š': "S",
	'ś': "s", 'ŝ': "s", 'ş': "s", 'š': "s",
	'Ţ': "T", 'Ť': "T", 'Ŧ': "T",
	'ţ': "t", 'ť': "t", 'ŧ': "t",
	'Ũ': "U", 'Ū': "U", 'Ŭ': "U", 'Ů': "U", 'Ű': "U", 'Ų': "U",
	'ũ': "u", 'ū': "u", 'ŭ': "u", 'ů': "u", 'ű': "u", 'ų': "u",
	'Ŵ': "W", 'ŵ': "w",
	'Ŷ': "Y", 'ŷ': "y", 'Ÿ': "Y",
	'Ź': "Z", 'Ż': "Z", 'Ž': "Z",
	'ź': "z", 'ż': "z", 'ž': "z",
	'Ĳ': "IJ", 'ĳ': "ij",
	'Œ': "Oe", 'œ': "oe",
	'ŉ': "'n", 'ſ': "s",
}

// isComboMark reports whether r is a combining diacritical mark that lodash's
// _.deburr strips (the U+0300–U+036F block plus U+FE20–U+FE23).
func isComboMark(r rune) bool {
	return (r >= 0x0300 && r <= 0x036f) || (r >= 0xfe20 && r <= 0xfe23)
}

// Deburr converts Latin-1 Supplement and Latin Extended-A letters in s to their
// basic Latin equivalents and removes combining diacritical marks, matching
// lodash's _.deburr. Ligatures expand to multiple characters (e.g. 'Æ' -> "Ae")
// and Latin mathematical operators such as '×' and '÷' are left unchanged.
func Deburr(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if repl, ok := deburrMap[r]; ok {
			b.WriteString(repl)
		} else if !isComboMark(r) {
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
