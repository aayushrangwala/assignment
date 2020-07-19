package internal

import (
	"strings"
	"unicode"
)

// CeaserCipher is the ceaser cipher encoder/decoder helper function based on the input string and key for shift
// It takes an op as a parameter where if op = +1 => encode, op = -1 => decode
func CeaserCipher(in string, key, op int) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			return r
		}
		// since the alphabets are all lower case in the requirement
		if unicode.IsUpper(r) {
			r += 32 // make lower
		}
		r += rune(op * (key % AlphabetsLength))
		if r > 'z' {
			r = 'a' + (r - ('z' + 1))
		} else if r < 'a' {
			r = 'z' - (('a' - 1) - r)
		}
		return r
	}, in)
}
