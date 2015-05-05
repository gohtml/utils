package utils

import (
	"unicode"
	"unicode/utf8"

	"github.com/golangplus/bytes"
)

type byteEscapeTable []string

func escapeString(s string, table byteEscapeTable) string {
	var bs bytesp.ByteSlice
	scanned := 0
	for i, r := range s {
		if int(r) < len(table) {
			if esc := table[r]; len(esc) != 0 {
				bs.WriteString(s[scanned:i])
				bs.WriteString(esc)

				scanned = i + utf8.RuneLen(r)
			}
		}
	}

	if scanned == 0 {
		return s
	}
	bs.WriteString(s[scanned:])
	return string(bs)
}

// http://www.w3.org/TR/html5/syntax.html#syntax-attributes
// Excluding 0x7f..0x9f
var invalidAttrNameBytes ascMask

func init() {
	invalidAttrNameBytes = IsSpaceCharacter
	invalidAttrNameBytes.UnionWith(ascMaskFromString("\"'>/="))
	invalidAttrNameBytes.SetRange(0x00, 0x1F)
}

// Normalize an attribute name string.
// http://www.w3.org/TR/html5/syntax.html#syntax-attributes
func NormAttrName(s string) string {
	var bs bytesp.ByteSlice
	scanned := 0
	for i, r := range s {
		lower := unicode.ToLower(r)
		if lower != r {
			bs.WriteString(s[scanned:i])
			bs.WriteRune(lower)
			scanned = i + utf8.RuneLen(r)
		} else if (r < 128 && invalidAttrNameBytes[r]) || unicode.IsControl(r) {
			bs.WriteString(s[scanned:i])
			scanned = i + utf8.RuneLen(r)
		}
	}

	if scanned == 0 {
		return s
	}

	bs.WriteString(s[scanned:])

	return string(bs)
}

var htmlEscapeTable = byteEscapeTable{
	0xA0: "&nbsp",
	'"':  "&quot;",
	'&':  "&amp;",
	'<':  "&lt;",
	'>':  "&gt;",
}

func EscapeHTML(s string) string {
	return s
}

// http://www.w3.org/TR/html5/syntax.html#escapingString
var attrEscapeTable = byteEscapeTable{
	0xA0: "&nbsp;",
	'"':  "&quot;",
	'&':  "&amp;",
}

// Escapes a string so that it is a valid attribute value.
func EscapeAttr(s string) string {
	return escapeString(s, attrEscapeTable)
}

func appendByteMaskFilteredString(bs bytesp.ByteSlice, s string, allowed *ascMask) bytesp.ByteSlice {
	scanned := 0
	for i, n := 0, len(s); i < n; i++ {
		b := s[i]
		if b >= 128 || !allowed[b] {
			bs.WriteString(s[scanned:i])
			scanned = i + 1
		}
	}

	bs.WriteString(s[scanned:])

	return bs
}

var hexDigits = [16]byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F',
}

var dec2hex [256]string

func init() {
	for i := range dec2hex {
		dec2hex[i] = string([]byte{hexDigits[i/16], hexDigits[i%16]})
	}
}

func appendAscMaskPctEncodedString(bs bytesp.ByteSlice, s string, unchanged *ascMask) bytesp.ByteSlice {
	scanned := 0
	for i, n := 0, len(s); i < n; i++ {
		b := s[i]
		if b >= 128 || !unchanged[b] {
			if i > scanned {
				bs.WriteString(s[scanned:i])
			}
			bs.WriteByte('%')
			bs.WriteString(dec2hex[b])

			scanned = i + 1
		}
	}

	bs.WriteString(s[scanned:len(s)])

	return bs
}

// RFC 3986: reserved
var isUrlUnreserved ascMask

func init() {
	isUrlUnreserved.SetRange('a', 'z')
	isUrlUnreserved.SetRange('A', 'Z')
	isUrlUnreserved.SetRange('0', '9')
	isUrlUnreserved.SetByString("*-._~")
}

// RFC 3986: gen-delims
var isUrlGenDelimis = ascMaskFromString(":/?#[]@")

// RFC 3986: sub-delims
var isUrlSubDelims = ascMaskFromString("!$&'()*+,;=")

var isUrlIpLiteralChars ascMask

func init() {
	isUrlIpLiteralChars = isUrlUnreserved
	isUrlIpLiteralChars.UnionWith(isUrlSubDelims)
	isUrlIpLiteralChars[':'] = true
}

var isUrlRegNameChars ascMask

func init() {
	isUrlRegNameChars = isUrlUnreserved
	isUrlRegNameChars.UnionWith(isUrlSubDelims)
}

// Escapes a string so that it is a valid query part(name or value).
func EscapeQuery(s string) string {
	var bs bytesp.ByteSlice
	scanned := 0
	for i, n := 0, len(s); i < n; i++ {
		b := s[i]
		if b >= 128 || !isUrlUnreserved[b] {
			bs.WriteString(s[scanned:i])
			if b == ' ' {
				bs.WriteByte('+')
			} else {
				bs.WriteByte('%')
				bs.WriteString(dec2hex[b])
			}
			scanned = i + 1
		}
	}

	if scanned == 0 {
		return s
	}
	bs.WriteString(s[scanned:])
	return string(bs)
}

func escapeIPliteral(s string) string {
	var bs bytesp.ByteSlice
	bs.WriteByte('[')
	bs = appendByteMaskFilteredString(bs, s[1:len(s)-1], &isUrlIpLiteralChars)
	bs.WriteByte(']')
	return string(bs)
}

// Escapes/filters a string so that it is a valid hostname.
func EscapeHost(s string) string {
	if len(s) > 4 && s[0] == '[' && s[len(s)-1] == ']' {
		// RFC 3986: IP-literal
		return escapeIPliteral(s)
	}

	return string(appendByteMaskFilteredString(nil, s, &isUrlRegNameChars))
}
