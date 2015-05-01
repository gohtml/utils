package utils

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestHTMLNodeSet(t *testing.T) {
	var set HTMLNodeSet
	set.Put("a")
	assert.StringEqual(t, "set", "[a]", set)

	set.Put("a")
	assert.StringEqual(t, "set", "[a]", set)

	set.Put("b")
	assert.StringEqual(t, "set", "[a b]", set)

	set.Put("d")
	assert.StringEqual(t, "set", "[a b d]", set)

	set.Put("a")
	assert.StringEqual(t, "set", "[a b d]", set)

	set.Del("c")
	assert.StringEqual(t, "set", "[a b d]", set)

	set.Del("a")
	assert.StringEqual(t, "set", "[b d]", set)
}
