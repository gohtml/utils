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

func TestByteMask(t *testing.T) {
	var allowed byteMask
	assert.StringEqual(t, "allowed", allowed, `""`)

	allowed['a'] = true
	assert.StringEqual(t, "allowed", allowed, `"a"`)

	allowed = byteMaskFromString("aBc")
	assert.StringEqual(t, "allowed", allowed, `"Bac"`)

	allowed.UnionWith(byteMaskFromString("DeF"))
	assert.StringEqual(t, "allowed", allowed, `"BDFace"`)

	allowed.SetRange('0', '9')
	assert.StringEqual(t, "allowed", allowed, `"0123456789BDFace"`)
}
