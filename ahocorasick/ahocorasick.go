package ahocorasick

import (
	"github.com/koron/gelatin/trie"
)

type Matcher struct {
	trie *trie.TernaryTrie
}

type Match struct {
	Index   int
	Pattern string
	Value   interface{}
}

type NodeData struct {
	pattern *string
	value   interface{}
	failure *trie.TernaryNode
}

func New() *Matcher {
	return &Matcher{
		trie: trie.NewTernaryTrie(),
	}
}

func (m *Matcher) Add(pattern string, v interface{}) {
	m.trie.Put(pattern, &NodeData{
		pattern: &pattern,
		value:   v,
	})
}

func (m *Matcher) Compile() error {
	m.trie.Balance()
	root := m.trie.Root().(*trie.TernaryNode)
	root.SetValue(&NodeData{failure: root})
	// fill data.failure of each node.
	trie.EachWidth(m.trie, func(n trie.Node) bool {
		parent := n.(*trie.TernaryNode)
		parent.Each(func(m trie.Node) bool {
			fillFailure(m.(*trie.TernaryNode), root, parent)
			return true
		})
		return true
	})
	return nil
}

func fillFailure(curr, root, parent *trie.TernaryNode) {
	data := nodeData(curr)
	if data == nil {
		data = &NodeData{}
		curr.SetValue(data)
	}
	if parent == root {
		data.failure = root
		return
	}
	// Determine failure node.
	r := curr.Label()
	node := nodeDataFailure(parent, root)
	for {
		next, _ := node.Get(r).(*trie.TernaryNode)
		if next != nil {
			node = next
			break
		} else if node == root {
			break
		}
		node = nodeDataFailure(node, root)
	}
	data.failure = node
}

func (m *Matcher) Match(text string) <-chan Match {
	ch := make(chan Match, 1)
	go m.startMatch(text, ch)
	return ch
}

func (m *Matcher) startMatch(text string, ch chan<- Match) {
	defer close(ch)
	root := m.trie.Root().(*trie.TernaryNode)
	curr := root
	for i, r := range text {
		curr = nextState(curr, root, r)
		if curr == root {
			continue
		}
		fireAll(curr, root, ch, i)
	}
}

func nextState(curr, root *trie.TernaryNode, r rune) *trie.TernaryNode {
	for {
		next, _ := curr.Get(r).(*trie.TernaryNode)
		if next != nil {
			return next
		} else if curr == root {
			return root
		}
		curr = nodeDataFailure(curr, root)
	}
}

func fireAll(curr, root *trie.TernaryNode, ch chan<- Match, idx int) {
	for curr != root {
		data := nodeData(curr)
		if data.pattern != nil {
			ch <- Match{
				Index:   idx - len(*data.pattern) + 1,
				Pattern: *data.pattern,
				Value:   data.value,
			}
		}
		curr = data.failure
	}
}

func nodeData(node *trie.TernaryNode) *NodeData {
	d, _ := node.Value().(*NodeData)
	return d
}

func nodeDataFailure(node, root *trie.TernaryNode) *trie.TernaryNode {
	next := nodeData(node).failure
	if next == nil {
		return root
	}
	return next
}
