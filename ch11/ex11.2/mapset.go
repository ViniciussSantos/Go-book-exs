package set

import (
	"bytes"
	"fmt"
)

type MapSet struct {
	words map[int]bool
}

func (s *MapSet) Has(x int) bool {
	_, ok := s.words[x]
	return ok
}

func (s *MapSet) Add(x int) {
	s.words[x] = true
}

func (s *MapSet) UnionWith(t *MapSet) {
	for k := range t.words {
		s.words[k] = true
	}
}

func (s *MapSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k := range s.words {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", k)
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *MapSet) Len() int {
	return len(s.words)
}

func (s *MapSet) Remove(x int) {
	delete(s.words, x)
}

func (s *MapSet) Clear() {
	s.words = map[int]bool{}
}
