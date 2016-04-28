package base36

import (
	"math"
	"strings"
)

var (
	base36 = []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z'}
)

// Encode encodes a value to base36
func Encode(value uint64) string {

	res := ""
	for value != 0 {
		res = string(base36[value%36]) + res
		value /= 36
	}
	return res
}

// Decode decodes a base36-encoded string
func Decode(s string) int64 {

	s = reverse(strings.ToUpper(s))
	res := int64(0)

	for idx, c := range s {
		byteOffset := indexOf(byte(c))
		res += int64(int64(byteOffset) * int64(math.Pow(36, float64(idx))))
	}
	return res
}

func indexOf(c byte) int {

	for idx, b := range base36 {
		if b == c {
			return idx
		}
	}
	return 0
}

// reverse returns its argument string reversed rune-wise.
func reverse(s string) string {

	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
