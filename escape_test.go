package utils

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestAppendByteMaskFilteredString(t *testing.T) {
	allowed := ascMaskFromString("ABC")
	assert.StringEqual(t, "filtered",
		string(appendByteMaskFilteredString(nil, "ABC123", &allowed)), `ABC`)
	assert.StringEqual(t, "filtered",
		string(appendByteMaskFilteredString(nil, "ABC中文123", &allowed)), `ABC`)
}

func TestAppendByteMaskPctEncodedString(t *testing.T) {
	unchanged := ascMaskFromString("ABC")
	assert.StringEqual(t, "encoded",
		string(appendAscMaskPctEncodedString(nil, "ABC123", &unchanged)), `ABC%31%32%33`)
	assert.StringEqual(t, "encoded",
		string(appendAscMaskPctEncodedString(nil, "ABC中文123", &unchanged)), `ABC%E4%B8%AD%E6%96%87%31%32%33`)
}

func TestEscapeQuery(t *testing.T) {
	assert.Equal(t, "EscapeQuery", EscapeQuery("English中文"), "English%E4%B8%AD%E6%96%87")
	assert.Equal(t, "EscapeQuery", EscapeQuery("abcABC1.2-3"), "abcABC1.2-3")
	assert.Equal(t, "EscapeQuery", EscapeQuery("abc DEF"), "abc+DEF")
	assert.Equal(t, "EscapeQuery", EscapeQuery("abc&123"), "abc%26123")
}

func TestEscapeHost(t *testing.T) {
	assert.Equal(t, "EscapeHost", EscapeHost("English中文"), "English")
	assert.Equal(t, "EscapeHost", EscapeHost("www.example.com"), "www.example.com")
	assert.Equal(t, "EscapeHost", EscapeHost("www.?git-hub.com://"), "www.git-hub.com")
	assert.Equal(t, "EscapeHost", EscapeHost("www.e中文e.com"), "www.ee.com")
}

func TestEscapeAttr(t *testing.T) {
	assert.Equal(t, "EscapeAttr", EscapeAttr("English中文"), "English中文")
	assert.Equal(t, "EscapeAttr", EscapeAttr("\"&amp;value\""), "&quot;&amp;amp;value&quot;")
	assert.Equal(t, "EscapeAttr", EscapeAttr("\u00a0"), "&nbsp;")
}

func TestNormAttrName(t *testing.T) {
	assert.Equal(t, "NormAttrName", NormAttrName("English中文"), "english中文")
	assert.Equal(t, "NormAttrName", NormAttrName("S RC"), "src")
	assert.Equal(t, "NormAttrName", NormAttrName("a\t<=\rb"), "a<b")
	assert.Equal(t, "NormAttrName", NormAttrName("a\u0080b"), "ab")
}
