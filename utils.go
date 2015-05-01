package utils

import (
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
