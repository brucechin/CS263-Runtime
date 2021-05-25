package graph

import (
	"math/rand"
	"testing"
)

func benchSSSP(b *testing.B, size int) {
	b.StopTimer()
	g := New(size)
	for i := 0; i < size; i++ {
		// for j := 0; j < size/100; j++ {
		for j := 0; j < 20; j++ {
			g.Add(i, rand.Intn(size))
			// g.Add(0, rand.Intn(n))
			// g.Add(rand.Intn(n), rand.Intn(n))
		}
	}
	b.StartTimer()

	var dist = make([]int64, size)
	var parent = make([]int, size)

	for i := 0; i < b.N; i++ {
		ShortestPaths2(g, parent, dist, 0)
	}
}

func BenchmarkSSSP1e4(b *testing.B)  { benchSSSP(b, 10000) }
func BenchmarkSSSP4e4(b *testing.B)  { benchSSSP(b, 40000) }
func BenchmarkSSSP16e4(b *testing.B) { benchSSSP(b, 160000) }
