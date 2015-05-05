package utils

import (
	"strconv"

	"github.com/golangplus/bytes"
)

// byteMask represents a set of bytes by setting a boolean
// value for each possible byte.
type byteMask [256]bool

func (mask byteMask) String() string {
	var bs bytesp.ByteSlice
	for c, in := range mask {
		if in {
			bs.WriteByte(byte(c))
		}
	}
	return strconv.Quote(string(bs))
}

func byteMaskFromString(s string) (mask byteMask) {
	mask.SetByString(s)
	return
}

// Returns self for chaining grammar.
// that is not passed as a pointer for easier chain operation. This is not efficient but
// more convinient. Use it only in initialization part.
func (this *byteMask) UnionWith(that byteMask) *byteMask {
	for i, el := range that {
		if el {
			this[i] = true
		}
	}

	return this
}

func (arr *byteMask) SetByString(s string) {
	for _, c := range s {
		if c < rune(len(arr)) {
			arr[c] = true
		}
	}
}

func (arr *byteMask) SetRange(mn, mx byte) {
	for i := int(mn); i <= int(mx); i++ {
		arr[i] = true
	}
}
