package trie

import (
	"container/list"
)

// Trie provide interface of trie data structure.
type Trie interface {
	Root() Node
	Get(KeySeq) Node
	Put(KeySeq, interface{}) Node
	Size() int
}

// NewTrie returns a new Trie instance.
func NewTrie() Trie {
	return NewTernaryTrie()
}

// Get finds and returns Node for KeySeq in Trie.
func Get(t Trie, k KeySeq) Node {
	if t == nil {
		return nil
	}
	n := t.Root()
	for _, c := range k.Keys() {
		n = n.Get(c)
		if n == nil {
			return nil
		}
	}
	return n
}

// Put adds a pair of key and value to Trie.
func Put(t Trie, k KeySeq, v interface{}) Node {
	if t == nil {
		return nil
	}
	n := t.Root()
	for _, c := range k.Keys() {
		n, _ = n.Dig(c)
	}
	n.SetValue(v)
	return n
}

// EachDepth traverses all nodes in Trie by depth first order.
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

// EachWidth traverses all nodes in Trie by breadth first order.
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

// Node provides interface for a node of Trie.
type Node interface {
	Get(k Key) Node
	Dig(k Key) (Node, bool)
	HasChildren() bool
	Size() int
	Each(func(Node) bool)
	RemoveAll()

	Label() Key
	Value() interface{}
	SetValue(v interface{})
}

// Children returns all children of a node.
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
