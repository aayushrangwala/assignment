package internal

import (
	"strings"
	"unicode"
)

// CeaserCipherEncode is the ceaser cipher encoder helper function based on the input string and key for shift
func CeaserCipherEncode(in string, key int) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			return r
		}
		// since the alphabets are all lower case in the requirement
		if unicode.IsUpper(r) {
			r += 32 // make lower
		}
		r += rune(key % AlphabetsLength)
		if r > 'z' {
			r = 'a' + (r - ('z' + 1))
		} else if r < 'a' {
			r = 'z' - (('a' - 1) - r)
		}
		return r
	}, in)
}
