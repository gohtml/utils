package utils

import (
	"strconv"

	"github.com/golangplus/bytes"
)

// IsSpaceCharacters is a byteMask defining whether a byte is a space characters.
// http://www.w3.org/TR/html5/infrastructure.html#space-character
var IsSpaceCharacter ascMask = ascMaskFromString(" \t\n\f\r")

// byteMask represents a set of bytes by setting a boolean
// value for each possible byte.
type ascMask [128]bool

func (mask ascMask) String() string {
	var bs bytesp.ByteSlice
	for c, in := range mask {
		if in {
			bs.WriteByte(byte(c))
		}
	}
	return strconv.Quote(string(bs))
}

func ascMaskFromString(s string) (mask ascMask) {
	mask.SetByString(s)
	return
}

func (this ascMask) Union(that ascMask) ascMask {
	return *this.UnionWith(that)
}

// Returns self for chaining grammar.
// that is not passed as a pointer for easier chain operation. This is not efficient but
// more convinient. Use it only in initialization part.
func (this *ascMask) UnionWith(that ascMask) *ascMask {
	for i, el := range that {
		if el {
			this[i] = true
		}
	}

	return this
}

func (arr *ascMask) SetByString(s string) {
	for _, c := range s {
		if c < rune(len(arr)) {
			arr[c] = true
		}
	}
}

// SetRange sets bytes between mn and mx inclusively.
func (arr *ascMask) SetRange(mn, mx byte) {
	for i := int(mn); i <= int(mx); i++ {
		arr[i] = true
	}
}
