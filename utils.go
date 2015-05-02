package utils

import (
	"strconv"

	"github.com/golangplus/bytes"

	. "github.com/gohtml/elements"
)

// A HTMLNode set implemented by a slice.
// A slice is more time and memofy efficient than a map when
// the number of elements is small.
type HTMLNodeSet []HTMLNode

// Puts a new element into the set.
func (set *HTMLNodeSet) Put(s HTMLNode) {
	for _, el := range *set {
		if el == s {
			return
		}
	}

	*set = append(*set, s)
}

// Deletes a new element from the set.
func (set *HTMLNodeSet) Del(s HTMLNode) {
	for i, el := range *set {
		if el == s {
			*set = append((*set)[:i], (*set)[i+1:]...)
			return
		}
	}
}

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
