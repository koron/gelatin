package trie

// TernaryTrie provides ternary trie-tree.
type TernaryTrie struct {
	root TernaryNode
}

// NewTernaryTrie creates a ternary trie-tree.
func NewTernaryTrie() *TernaryTrie {
	return &TernaryTrie{}
}

// Root returns the root node of the trie-tree.
func (t *TernaryTrie) Root() Node {
	return &t.root
}

// Get returns a node for key k.
func (t *TernaryTrie) Get(k string) Node {
	return Get(t, k)
}

// Put puts a pair of key and value to trie-tree.
func (t *TernaryTrie) Put(k string, v interface{}) Node {
	return Put(t, k, v)
}

// Size counts nodes in the trie-tree.
func (t *TernaryTrie) Size() int {
	count := 0
	EachDepth(t, func(Node) bool {
		count++
		return true
	})
	return count
}

// Balance balances all nodes of trie-tree in each layers.
func (t *TernaryTrie) Balance() {
	EachDepth(t, func(n Node) bool {
		n.(*TernaryNode).Balance()
		return true
	})
	t.root.Balance()
}

// TernaryNode provides node of ternary trie-tree.
type TernaryNode struct {
	label      rune
	firstChild *TernaryNode
	low, high  *TernaryNode
	value      interface{}
}

// NewTernaryNode creates a node instance.
func NewTernaryNode(l rune) *TernaryNode {
	return &TernaryNode{label: l}
}

// Get returns a child node for k.
func (n *TernaryNode) Get(k rune) Node {
	curr := n.firstChild
	for curr != nil {
		if k == curr.label {
			return curr
		} else if k < curr.label {
			curr = curr.low
		} else {
			curr = curr.high
		}
	}
	return nil
}

// Dig digs a node for k. it returns node and a flag for whether dig or not.
func (n *TernaryNode) Dig(k rune) (node Node, isnew bool) {
	curr := n.firstChild
	if curr == nil {
		n.firstChild = NewTernaryNode(k)
		return n.firstChild, true
	}
	for {
		if k == curr.label {
			return curr, false
		} else if k < curr.label {
			if curr.low == nil {
				curr.low = NewTernaryNode(k)
				return curr.low, true
			}
			curr = curr.low
		} else {
			if curr.high == nil {
				curr.high = NewTernaryNode(k)
				return curr.high, true
			}
			curr = curr.high
		}
	}
}

// FirstChild returns first child node.
func (n *TernaryNode) FirstChild() *TernaryNode {
	return n.firstChild
}

// HasChildren returns the node hash any children or not.
func (n *TernaryNode) HasChildren() bool {
	return n.firstChild != nil
}

// Size counts descended nodes.
func (n *TernaryNode) Size() int {
	if n.firstChild == nil {
		return 0
	}
	count := 0
	n.Each(func(Node) bool {
		count++
		return true
	})
	return count
}

// Each enumerates descended nodes.
func (n *TernaryNode) Each(proc func(Node) bool) {
	var f func(*TernaryNode) bool
	f = func(n *TernaryNode) bool {
		if n != nil {
			if !f(n.low) || !proc(n) || !f(n.high) {
				return false
			}
		}
		return true
	}
	f(n.firstChild)
}

// RemoveAll removes all descended nodes.
func (n *TernaryNode) RemoveAll() {
	n.firstChild = nil
}

// Label returns a label rune.
func (n *TernaryNode) Label() rune {
	return n.label
}

// Value returns a value for the node.
func (n *TernaryNode) Value() interface{} {
	return n.value
}

// SetValue set a value for the node.
func (n *TernaryNode) SetValue(v interface{}) {
	n.value = v
}

func (n *TernaryNode) children() []*TernaryNode {
	children := make([]*TernaryNode, n.Size())
	if n.firstChild == nil {
		return children
	}
	idx := 0
	n.Each(func(child Node) bool {
		children[idx] = child.(*TernaryNode)
		idx++
		return true
	})
	return children
}

// Balance balances all descended nodes.
func (n *TernaryNode) Balance() {
	if n.firstChild == nil {
		return
	}
	children := n.children()
	for _, child := range children {
		child.low = nil
		child.high = nil
	}
	n.firstChild = balance(children, 0, len(children))
}

func balance(nodes []*TernaryNode, s, e int) *TernaryNode {
	count := e - s
	if count <= 0 {
		return nil
	} else if count == 1 {
		return nodes[s]
	} else if count == 2 {
		nodes[s].high = nodes[s+1]
		return nodes[s]
	}
	mid := (s + e) / 2
	n := nodes[mid]
	n.low = balance(nodes, s, mid)
	n.high = balance(nodes, mid+1, e)
	return n
}
