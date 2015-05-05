package utils

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestByteMask(t *testing.T) {
	var allowed ascMask
	assert.StringEqual(t, "allowed", allowed, `""`)

	allowed['a'] = true
	assert.StringEqual(t, "allowed", allowed, `"a"`)

	allowed = ascMaskFromString("aBc")
	assert.StringEqual(t, "allowed", allowed, `"Bac"`)

	allowed.UnionWith(ascMaskFromString("DeF"))
	assert.StringEqual(t, "allowed", allowed, `"BDFace"`)

	allowed.SetRange('0', '9')
	assert.StringEqual(t, "allowed", allowed, `"0123456789BDFace"`)
}
