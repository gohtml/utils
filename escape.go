package utils

import (
	"github.com/golangplus/bytes"
)

type byteEscapeTable [256]string

func escapeString(s string, table *byteEscapeTable) string {
	var bs bytesp.ByteSlice
	scanned := 0
	for i, n := 0, len(s); i < n; i++ {
		b := s[i]
		if esc := table[b]; len(esc) != 0 {
			bs.WriteString(s[scanned:i])
			bs.WriteString(esc)

			scanned = i + 1
		}
	}

	if scanned == 0 {
		return s
	}
	bs.WriteString(s[scanned:])
	return string(bs)
}

func appendByteMaskFilteredString(bs bytesp.ByteSlice, s string, allowed *byteMask) bytesp.ByteSlice {
	scanned := 0
	for i, n := 0, len(s); i < n; i++ {
		b := s[i]
		if !allowed[b] {
			if i > scanned {
				bs.WriteString(s[scanned:i])
			}
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

func appendByteMaskPctEncodedString(bs bytesp.ByteSlice, s string, unchanged *byteMask) bytesp.ByteSlice {
	scanned := 0
	for i, n := 0, len(s); i < n; i++ {
		b := s[i]
		if !unchanged[b] {
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
var isUrlUnreserved byteMask

func init() {
	isUrlUnreserved.SetRange('a', 'z')
	isUrlUnreserved.SetRange('A', 'Z')
	isUrlUnreserved.SetRange('0', '9')
	isUrlUnreserved.SetByString("*-._~")
}

// RFC 3986: gen-delims
var isUrlGenDelimis = byteMaskFromString(":/?#[]@")

// RFC 3986: sub-delims
var isUrlSubDelims = byteMaskFromString("!$&'()*+,;=")

var isUrlIpLiteralChars byteMask

func init() {
	isUrlIpLiteralChars = isUrlUnreserved
	isUrlIpLiteralChars.UnionWith(isUrlSubDelims)
	isUrlIpLiteralChars[':'] = true
}

var isUrlRegNameChars byteMask

func init() {
	isUrlRegNameChars = isUrlUnreserved
	isUrlRegNameChars.UnionWith(isUrlSubDelims)
}

func EscapeQuery(s string) string {
	var bs bytesp.ByteSlice
	scanned := 0
	for i, n := 0, len(s); i < n; i++ {
		b := s[i]
		if !isUrlUnreserved[b] {
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
