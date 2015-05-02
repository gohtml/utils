package utils

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestAppendByteMaskFilteredString(t *testing.T) {
	allowed := byteMaskFromString("ABC")
	assert.StringEqual(t, "filtered",
		string(appendByteMaskFilteredString(nil, "ABC123", &allowed)), `ABC`)
}

func TestAppendByteMaskPctEncodedString(t *testing.T) {
	unchanged := byteMaskFromString("ABC")
	assert.StringEqual(t, "encoded",
		string(appendByteMaskPctEncodedString(nil, "ABC123", &unchanged)), `ABC%31%32%33`)
}

func TestEscapeQuery(t *testing.T) {
	assert.Equal(t, "EscapeQuery", EscapeQuery("abcABC1.2-3"), "abcABC1.2-3")
	assert.Equal(t, "EscapeQuery", EscapeQuery("abc DEF"), "abc+DEF")
	assert.Equal(t, "EscapeQuery", EscapeQuery("abc&123"), "abc%26123")
}
