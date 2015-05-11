package trie

// Order represents result of key elements (Key) comparation.
type Order int

const (
	// MISMATCH means type mismatch for key elements.
	MISMATCH Order = iota
	// EQUAL means matched key elements.
	EQUAL
	// BEFORE means target Key is placed before argument one.
	BEFORE
	// AFTER means target Key is placed after argument one.
	AFTER
)

// KeySeq provides interface for key of trie.
type KeySeq interface {
	Keys() []Key
}

// Key provides interface for key element.
type Key interface {
	Compare(Key) Order
}

// Equal checks two keys are same or not.
func Equal(a, b Key) bool {
	return a.Compare(b) == EQUAL
}
