package internal

import "testing"

func TestCeaserCipherEncode(t *testing.T) {
	var testCases = []struct {
		in, out string
		shift   int
	}{
		{"if-he-had-anything-confidential", "ol-nk-ngj-gteznotm-iutlojktzogr", 32},
		{"to-say", "zu-yge", 32},
		{"by-so-changing-the", "if-zv-johunpun-aol", 7},
		{"order-of-the", "xamna-xo-cqn", 9},
		{"letters of the", "exmmxkl hy max", 19},
		{"alphabet, that", "ximexybq, qexq", 23},
		{"not-a-word-could", "yze-l-hzco-nzfwo", 37},
		{"be made out.", "mp xlop zfe.", 11},
	}

	for _, tc := range testCases {
		out := CeaserCipher(tc.in, tc.shift, CeaserCipherEncode)
		t.Logf("Ceaser cipher, Input: [%s], Shift: [%d], Expected: [%s], out: [%s]", tc.in, tc.shift, tc.out, out)

		if out != tc.out {
			t.Errorf("[FAILED]. expected: %v, returned: %v ", tc.out, out)
		}
	}
}

func TestCeaserCipherDecode(t *testing.T) {
	var testCases = []struct {
		in, out string
		shift   int
	}{
		{"ol-nk-ngj-gteznotm-iutlojktzogr", "if-he-had-anything-confidential", 32},
		{"zu-yge", "to-say", 32},
		{"if-zv-johunpun-aol", "by-so-changing-the", 7},
		{"xamna-xo-cqn", "order-of-the", 9},
		{"exmmxkl hy max", "letters of the", 19},
		{"ximexybq, qexq", "alphabet, that", 23},
		{"yze-l-hzco-nzfwo", "not-a-word-could", 37},
		{"mp xlop zfe.", "be made out.", 11},
	}

	for _, tc := range testCases {
		out := CeaserCipher(tc.in, tc.shift, CeaserCipherDecode)
		t.Logf("Ceaser cipher, Input: [%s], Shift: [%d], Expected: [%s], out: [%s]", tc.in, tc.shift, tc.out, out)

		if out != tc.out {
			t.Errorf("[FAILED]. expected: %v, returned: %v ", tc.out, out)
		}
	}
}
