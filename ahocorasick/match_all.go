package ahocorasick

import (
	"container/list"
)

func MatchAll(m *Matcher, text string) []Match {
	list := list.New()
	ch := m.Match(text)
	for n := range ch {
		list.PushBack(n)
	}
	all := make([]Match, list.Len())
	idx := 0
	for e := list.Front(); e != nil; e = e.Next() {
		all[idx] = e.Value.(Match)
		idx++
	}
	return all
}
