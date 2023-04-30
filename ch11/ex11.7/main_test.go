package set

import (
	"math/rand"
	"testing"
)

func populateSet(set IntSet, n int) {
	for i := 0; i < n; i++ {
		set.Add(rand.Intn(1000))
	}
}

func BenchAdd(b *testing.B, set IntSet) {
	for i := 0; i < b.N; i++ {
		set.Add(rand.Intn(1000))
	}
}

func benchUnionWith(b *testing.B, setA, setB IntSet) {
	populateSet(setA, b.N)
	populateSet(setB, b.N)

	for i := 0; i < b.N; i++ {
		setA.UnionWith(&setB)
	}
}

func benchHas(b *testing.B, set IntSet) {
	populateSet(set, b.N)
	for i := 0; i < b.N; i++ {
		set.Has(i)
	}
}

func BenchmarkAdd(b *testing.B) {
	BenchAdd(b, IntSet{})
}

func BenchmarkUnionWith(b *testing.B) {
	benchUnionWith(b, IntSet{}, IntSet{})
}

func BenchmarkHas(b *testing.B) {
	benchHas(b, IntSet{})
}
