package trie

import (
	"container/list"
)

// Trie provides accessors for trie-tree
type Trie interface {
	Root() Node
	Get(string) Node
	Put(string, interface{}) Node
	Size() int
}

// NewTrie creates a new Trie instance.
func NewTrie() Trie {
	return NewTernaryTrie()
}

// Get gets a node for k as key.
func Get(t Trie, k string) Node {
	if t == nil {
		return nil
	}
	n := t.Root()
	for _, c := range k {
		n = n.Get(c)
		if n == nil {
			return nil
		}
	}
	return n
}

// Put puts a pair of key and value then returns the node for it.
func Put(t Trie, k string, v interface{}) Node {
	if t == nil {
		return nil
	}
	n := t.Root()
	for _, c := range k {
		n, _ = n.Dig(c)
	}
	n.SetValue(v)
	return n
}

// EachDepth enumerates nodes in trie for depth.
func EachDepth(t Trie, proc func(Node) bool) {
	if t == nil {
		return
	}
	r := t.Root()
	var f func(Node) bool
	f = func(n Node) bool {
		n.Each(f)
		return proc(n)
	}
	r.Each(f)
}

// EachWidth enumerates nodes in trie for width.
func EachWidth(t Trie, proc func(Node) bool) {
	if t == nil {
		return
	}
	q := list.New()
	q.PushBack(t.Root())
	for q.Len() != 0 {
		f := q.Front()
		q.Remove(f)
		t := f.Value.(Node)
		if !proc(t) {
			break
		}
		t.Each(func(n Node) bool {
			q.PushBack(n)
			return true
		})
	}
}

// Node provides accessors for nodes of trie-tree
type Node interface {

	// Get returns a child node for k.
	Get(k rune) Node

	// Dig digs a node for k. it returns node and a flag for whether dig or
	// not.
	Dig(k rune) (Node, bool)

	// HasChildren returns the node hash any children or not.
	HasChildren() bool

	// Size counts descended nodes.
	Size() int

	// Each enumerates descended nodes.
	Each(func(Node) bool)

	// RemoveAll removes all descended nodes.
	RemoveAll()

	// Label returns a label rune.
	Label() rune

	// Value returns a value for the node.
	Value() interface{}

	// SetValue set a value for the node.
	SetValue(v interface{})
}

// Children returns all children of the node.
func Children(n Node) []Node {
	children := make([]Node, n.Size())
	idx := 0
	n.Each(func(n Node) bool {
		children[idx] = n
		idx++
		return true
	})
	return children
}
