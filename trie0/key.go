package trie

type Order int

const (
	MISMATCH Order = iota
	EQUAL
	BEFORE
	AFTER
)

type KeySeq interface {
	Keys() []Key
}

type Key interface {
	Compare(Key) Order
}

func Equal(a, b Key) bool {
	return a.Compare(b) == EQUAL
}
