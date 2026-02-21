package stringx

import (
	"strings"
	"unicode/utf8"
)

// Len returns the actual character length of the string, not the byte length.
func Len(s string) int {
	return utf8.RuneCountInString(s)
}

// Reverse reverses the string.
func Reverse(s string) string {
	if len(s) <= 1 {
		return s
	}

	var (
		r    rune
		size int
	)
	b := strings.Builder{}
	for i := len(s); i > 0; i -= size {
		r, size = utf8.DecodeLastRuneInString(s[:i])
		b.WriteRune(r)
	}
	return b.String()
}

// PadLeft pads the string on the left with the specified padding string until it reaches the desired length.
func PadLeft(s string, length int, padding string) string {
	return pad(s, length, padding, -1)
}

// PadRight pads the string on the right with the specified padding string until it reaches the desired length.
func PadRight(s string, length int, padding string) string {
	return pad(s, length, padding, 1)
}

// PadBoth pads the string on both sides with the specified padding string until it reaches the desired length.
func PadBoth(s string, length int, padding string) string {
	return pad(s, length, padding, 0)
}

func pad(s string, length int, padding string, kind int) string {
	if len(padding) == 0 {
		return s
	}

	slen := Len(s)
	if slen >= length {
		return s
	}

	var leftLen, rightLen int
	switch {
	case kind < 0:
		// pad left
		leftLen = length - slen
	case kind > 0:
		// pad right
		rightLen = length - slen
	default:
		// pad both
		leftLen = (length - slen) / 2
		rightLen = length - slen - leftLen
	}

	plen := Len(padding)
	leftRepeat := leftLen / plen
	rightRepeat := rightLen / plen
	leftRemain := leftLen % plen
	rightRemain := rightLen % plen

	b := strings.Builder{}

	if leftRepeat > 0 {
		b.WriteString(strings.Repeat(padding, leftRepeat))
	}
	if leftRemain > 0 {
		count := 0
		for _, r := range padding {
			if count < leftRemain {
				b.WriteRune(r)
				count++
			} else {
				break
			}
		}
	}

	b.WriteString(s)

	if rightRepeat > 0 {
		b.WriteString(strings.Repeat(padding, rightRepeat))
	}
	if rightRemain > 0 {
		count := 0
		for _, r := range padding {
			if count >= rightRemain {
				break
			}
			b.WriteRune(r)
			count++
		}
	}

	return b.String()
}
