package ahocorasick

import (
	"github.com/koron/gelatin/trie"
	"testing"
)

func checkNode(t *testing.T, node trie.Node, size int, data NodeData) {
	if node == nil {
		t.Error("Nil node:", data)
	}
	if node.Size() != size {
		t.Errorf("Unexpected childrens: %d != %d", node.Size(), size)
	}
	d := node.Value().(*NodeData)
	if d == nil {
		t.Error("Nil data:", data, node)
	}
	if data.pattern != nil && *d.pattern != *data.pattern {
		t.Error("Pattern unmatched:", data, node, *d.pattern)
	}
	if data.value != nil && d.value != data.value {
		t.Error("Value unmatched:", data, node, d.value)
	}
	if d.failure == nil {
		t.Error("Nil failure:", data, node)
	} else if d.failure != data.failure {
		t.Errorf("Failure unmatched: data=%+v node=%+v d.failure=%+v",
			data, node, d.failure)
	}
}

func invalidData(failure trie.Node) NodeData {
	return NodeData{
		failure: failure.(*trie.TernaryNode),
	}
}

func validData(pattern string, value interface{}, failure trie.Node) NodeData {
	return NodeData{
		pattern: &pattern,
		value:   value,
		failure: failure.(*trie.TernaryNode),
	}
}

func newTestMatcher() *Matcher {
	m := New()
	m.Add("ab", 2)
	m.Add("bc", 4)
	m.Add("bab", 6)
	m.Add("d", 7)
	m.Add("abcde", 10)
	m.Compile()
	return m
}

func TestTree(t *testing.T) {
	m := newTestMatcher()
	// Check tree structure.
	r := m.trie.Root()
	checkNode(t, r, 3, invalidData(r))
	n1 := r.Get('a')
	checkNode(t, n1, 1, invalidData(r))
	n3 := r.Get('b')
	checkNode(t, n3, 2, invalidData(r))
	n7 := r.Get('d')
	checkNode(t, n7, 0, invalidData(r))
	n2 := n1.Get('b')
	checkNode(t, n2, 1, validData("ab", 2, n3))
	n4 := n3.Get('c')
	checkNode(t, n4, 0, validData("bc", 4, r))
	n5 := n3.Get('a')
	checkNode(t, n5, 1, invalidData(n1))
	n8 := n2.Get('c')
	checkNode(t, n8, 1, invalidData(n4))
	n6 := n5.Get('b')
	checkNode(t, n6, 0, validData("bab", 6, n2))
	n9 := n8.Get('d')
	checkNode(t, n9, 1, invalidData(n7))
	n10 := n9.Get('e')
	checkNode(t, n10, 0, validData("abcde", 10, r))
}

func assertMatches(t *testing.T, exp, act []Match) {
	if len(act) != len(exp) {
		t.Errorf("[]Match length is not %d (%d)", len(exp), len(act))
		t.Logf("  expected: %v", exp)
		t.Logf("  actually: %v", act)
	}
	for i, e := range exp {
		dump := false
		a := act[i]
		if a.Index != e.Index {
			t.Errorf("Index not matched at #%d\n", i)
			dump = true
		}
		if a.Pattern != e.Pattern {
			t.Errorf("Pattern not matched at #%d\n", i)
			dump = true
		}
		if a.Value != e.Value {
			t.Errorf("Value not matched at #%d\n", i)
			dump = true
		}
		if (dump) {
			t.Logf("  expected: %+v", e)
			t.Logf("  actually: %+v", a)
		}
	}
}

func TestResults(t *testing.T) {
	m := newTestMatcher()

	r1 := MatchAll(m, "abcde")
	assertMatches(t, []Match{
		Match{0, "ab", 2},
		Match{1, "bc", 4},
		Match{3, "d", 7},
		Match{0, "abcde", 10},
	}, r1)
}
