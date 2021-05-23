package graph

import (
	"math/rand"
	"testing"
)

func BenchmarkShortestPaths(b *testing.B) {
	n := 1000
	b.StopTimer()
	g := New(n)
	for i := 0; i < n; i++ {
		g.Add(0, rand.Intn(n))
		g.Add(rand.Intn(n), rand.Intn(n))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ShortestPaths(g, 0)
	}
}
