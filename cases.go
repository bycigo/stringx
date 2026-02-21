package stringx

import (
	"strings"
	"unicode"
)

// CamelCase converts a string to camel case.
func CamelCase(s string) string {
	return studlyCase(s, false)
}

// PascalCase converts a string to pascal case.
func PascalCase(s string) string {
	return studlyCase(s, true)
}

func studlyCase(s string, upperFirst bool) string {
	if len(s) == 0 {
		return s
	}

	var (
		isFirst  = true
		afterSep = false
	)

	b := strings.Builder{}
	for _, r := range s {
		if r == '-' || r == '_' || unicode.IsSpace(r) {
			afterSep = true
			continue
		}
		if isFirst {
			if upperFirst {
				r = unicode.ToUpper(r)
			} else {
				r = unicode.ToLower(r)
			}
			isFirst = false
			afterSep = false
		} else if afterSep {
			r = unicode.ToUpper(r)
			afterSep = false
		}
		b.WriteRune(r)
	}
	return b.String()
}

// SnakeCase converts a string to snake case.
func SnakeCase(s string) string {
	return snakeCase(s, '_')
}

// KebabCase converts a string to kebab case.
func KebabCase(s string) string {
	return snakeCase(s, '-')
}

func snakeCase(s string, sep byte) string {
	if len(s) == 0 {
		return s
	}

	type runeKind int
	const (
		runeKindUnknown runeKind = iota
		runeKindLower
		runeKindSep
		runeKindUpper
	)
	confirmRuneKine := func(r rune) runeKind {
		if r == '-' || r == '_' || unicode.IsSpace(r) {
			return runeKindSep
		} else if unicode.IsUpper(r) {
			return runeKindUpper
		}
		return runeKindLower
	}

	var (
		prevRune      rune
		prevRuneKind  runeKind
		prevWriteKind runeKind
	)

	writeSep := func(b *strings.Builder, sep byte) {
		if prevWriteKind != runeKindUnknown && prevWriteKind != runeKindSep {
			b.WriteByte(sep)
			prevWriteKind = runeKindSep
		}
	}
	writeRune := func(b *strings.Builder, r rune, kind runeKind) {
		b.WriteRune(r)
		prevWriteKind = kind
	}

	b := strings.Builder{}
	for _, r := range s {
		kind := confirmRuneKine(r)

		switch kind {
		case runeKindLower:
			if prevRuneKind == runeKindUnknown {
				// do nothing
			} else if prevRuneKind == runeKindLower {
				// do nothing
			} else if prevRuneKind == runeKindSep {
				// do nothing
			} else if prevRuneKind == runeKindUpper {
				writeSep(&b, sep)
				writeRune(&b, prevRune, prevRuneKind)
			}
			writeRune(&b, r, kind)
		case runeKindSep:
			if prevRuneKind == runeKindUnknown {
				// do nothing
			} else if prevRuneKind == runeKindLower {
				writeSep(&b, sep)
			} else if prevRuneKind == runeKindSep {
				// do nothing
			} else if prevRuneKind == runeKindUpper {
				writeRune(&b, prevRune, prevRuneKind)
				writeSep(&b, sep)
			}
		case runeKindUpper:
			r = unicode.ToLower(r) // force lowercase
			if prevRuneKind == runeKindUnknown {
				// do nothing
			} else if prevRuneKind == runeKindLower {
				writeSep(&b, sep)
			} else if prevRuneKind == runeKindSep {
				// do nothing
			} else if prevRuneKind == runeKindUpper {
				writeRune(&b, prevRune, prevRuneKind)
			}
		}

		prevRuneKind = kind
		prevRune = r
	}

	// do not ignore the last uppercase character
	if prevRuneKind == runeKindUpper {
		writeRune(&b, prevRune, prevRuneKind)
	}

	// remove trailing separator
	s = b.String()
	if prevRuneKind == runeKindSep {
		s = s[:len(s)-1]
	}

	return s
}
