package sqlparser

// DontEscape tells you if a character should not be escaped.
var dontEscape = byte(255)

// sqlEncodeMap specifies how to escape binary data with '\'.
// Complies to http://dev.mysql.com/doc/refman/5.7/en/string-literals.html
var sqlEncodeMap [256]byte

// sqlDecodeMap is the reverse of sqlEncodeMap
var sqlDecodeMap [256]byte

var encodeRef = map[byte]byte{
	'\x00': '0',
	'\'':   '\'',
	'"':    '"',
	'\b':   'b',
	'\n':   'n',
	'\r':   'r',
	'\t':   't',
	26:     'Z', // ctl-Z
	'\\':   '\\',
}

func init() {
	for i := range sqlEncodeMap {
		sqlEncodeMap[i] = dontEscape
		sqlDecodeMap[i] = dontEscape
	}
	for i := range sqlEncodeMap {
		if to, ok := encodeRef[byte(i)]; ok {
			sqlEncodeMap[byte(i)] = to
			sqlDecodeMap[to] = byte(i)
		}
	}
}

func encodeBytesSQL(val []byte, b *TrackedBuffer) {
	writebyte('\'', b)
	for _, ch := range val {
		if encodedChar := sqlEncodeMap[ch]; encodedChar == dontEscape {
			writebyte(ch, b)
		} else {
			writebyte('\\', b)
			writebyte(encodedChar, b)
		}
	}
	writebyte('\'', b)
}

func writebyte(c byte, b *TrackedBuffer) {
	if err := b.WriteByte(c); err != nil {
		panic(err)
	}
}
