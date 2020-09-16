package perlcrc

import (
	"testing"
)

func TestCRC32String(t *testing.T) {

	// Generated with perl's one-liner
	// using its module String::CRC32.
	var cases = map[string]uint32{
		"12ioj@!#":         769226953,
		"daskdjos":         2776331782,
		"mail@mail.ru":     568599782,
		"das sad912 &*@!0": 1837883029,
		"!!!!@(k0829382":   3997295746,
		"THisIsText":       147032099,
		"12817":            1423523988,
	}

	for str, expect := range cases {
		if got := CRC32String(str); got != expect {
			t.Errorf("Invalid result: got %d, expected %d", got, expect)
		}
	}
}
