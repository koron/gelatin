package ahocorasick

import (
	"unicode/utf8"

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

type nodeData struct {
	pattern *string
	offset  int
	value   interface{}
	failure *trie.TernaryNode
}

func New() *Matcher {
	return &Matcher{
		trie: trie.NewTernaryTrie(),
	}
}

func (m *Matcher) Add(pattern string, v interface{}) {
	_, n := utf8.DecodeLastRuneInString(pattern)
	m.trie.Put(pattern, &nodeData{
		pattern: &pattern,
		offset:  len(pattern) - n,
		value:   v,
	})
}

func (m *Matcher) Compile() error {
	m.trie.Balance()
	root := m.trie.Root().(*trie.TernaryNode)
	root.SetValue(&nodeData{failure: root})
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
	data := getNodeData(curr)
	if data == nil {
		data = &nodeData{}
		curr.SetValue(data)
	}
	if parent == root {
		data.failure = root
		return
	}
	// Determine failure node.
	fnode := getNextNode(getNodeFailure(parent, root), root, curr.Label())
	data.failure = fnode
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
		curr = getNextNode(curr, root, r)
		if curr == root {
			continue
		}
		fireAll(curr, root, ch, i)
	}
}

func getNextNode(node, root *trie.TernaryNode, r rune) *trie.TernaryNode {
	for {
		next, _ := node.Get(r).(*trie.TernaryNode)
		if next != nil {
			return next
		} else if node == root {
			return root
		}
		node = getNodeFailure(node, root)
	}
}

func fireAll(curr, root *trie.TernaryNode, ch chan<- Match, idx int) {
	for curr != root {
		data := getNodeData(curr)
		if data.pattern != nil {
			ch <- Match{
				Index:   idx - data.offset,
				Pattern: *data.pattern,
				Value:   data.value,
			}
		}
		curr = data.failure
	}
}

func getNodeData(node *trie.TernaryNode) *nodeData {
	d, _ := node.Value().(*nodeData)
	return d
}

func getNodeFailure(node, root *trie.TernaryNode) *trie.TernaryNode {
	next := getNodeData(node).failure
	if next == nil {
		return root
	}
	return next
}
