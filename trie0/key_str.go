package trie

import (
	"unicode/utf8"
)

// KeySeqString represents KeySeq for string.
type KeySeqString string

// Keys implements KeySeq.Keys for string.
func (s KeySeqString) Keys() []Key {
	keys := make([]Key, 0, utf8.RuneCountInString(string(s)))
	for _, r := range s {
		keys = append(keys, KeyRune(r))
	}
	return keys
}

// KeyRune represents Key for string (rune).
type KeyRune rune

// Compare implements KeyRune.Compare for string.
func (r KeyRune) Compare(v Key) Order {
	r2, ok := v.(KeyRune)
	if !ok {
		return MISMATCH
	} else if r == r2 {
		return EQUAL
	} else if r < r2 {
		return BEFORE
	}
	return AFTER
}
