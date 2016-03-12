package base36

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {

	assert.Equal(t, "15N9Z8L3AU4EB", Encode(5481594952936519619))

	assert.Equal(t, "1Y2P0IJ32E8E7", Encode(math.MaxInt64))
}

func TestDecode(t *testing.T) {

	assert.Equal(t, int64(5481594952936519619), Decode("15N9Z8L3AU4EB"))

	assert.Equal(t, int64(math.MaxInt64), Decode("1Y2P0IJ32E8E7"))
}
