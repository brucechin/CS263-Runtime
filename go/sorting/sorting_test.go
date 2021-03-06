// file: sort_test.go
// 		benchmark sorting algorithms in go libraries
// 		run `go test -bench . -benchmem` in the same folder
//
// 		code adapted from golang/go/src/sort/sort.go.
//

package sorting

import (
	"math/rand"
	"sort"
	"testing"
)

// testbed for int sorting, adapted from go/src/sort/sort_test.go
func benchInt(b *testing.B, size int, algo func([]int), initializer func(int, []int)) {

	b.StopTimer()
	unsorted := make([]int, size)
	initializer(size, unsorted)

	data := make([]int, size)
	for i := 0; i < b.N; i++ {
		// copy the original unsorted slice
		copy(data, unsorted)
		// run algorithm
		b.StartTimer()
		algo(data)
		b.StopTimer()
	}
}

// testbed for float64 sorting, adapted from go/src/sort/sort_test.go
func benchDouble(b *testing.B, size int, algo func([]float64), initializer func(int, []float64)) {

	b.StopTimer()
	unsorted := make([]float64, size)
	initializer(size, unsorted)

	data := make([]float64, size)
	for i := 0; i < b.N; i++ {
		// copy the original unsorted slice
		copy(data, unsorted)
		// run algorithm
		b.StartTimer()
		algo(data)
		b.StopTimer()
	}
}

//
// initializers
//

// a random initializer for int
func initRandomInt(size int, s []int) {
	for i := 0; i < size; i++ {
		s[i] = i ^ 0xabc
	}
}

// a random initializer for float64
func initRandomDouble(size int, s []float64) {
	r := rand.New(rand.NewSource(2021))
	for i := 0; i < size; i++ {
		s[i] = r.Float64() * 10
	}
}

// possible other file-loaders or std-in scanners

//
// algorithms
//

// 1. sort.Sort . This function internally uses quicksort + heapsort (leaf)
func SortInt(data []int)        { sort.Sort(sort.IntSlice(data)) }
func SortDouble(data []float64) { sort.Sort(sort.Float64Slice(data)) }

// 2. sort.Stable . This function internally uses mergesort + insertion (leaf)
func StableInt(data []int)        { sort.Stable(sort.IntSlice(data)) }
func StableDouble(data []float64) { sort.Stable(sort.Float64Slice(data)) }

// generated by  this python snippet
//
// for algo in ['Sort','Stable']:
//   for t in ['Int','Double']:
//     for s in [10000, 40000, 160000, 640000, 2560000, 10240000]:
//       ans = "func BenchmarkSort%s%s%s(b *testing.B) { bench%s(b, %s, %s%s, initRandom%s)}"%(t, s, algo, t, s, algo, t, t)
// 	  print(ans)

func BenchmarkSortInt10000Sort(b *testing.B)    { benchInt(b, 10000, SortInt, initRandomInt) }
func BenchmarkSortInt40000Sort(b *testing.B)    { benchInt(b, 40000, SortInt, initRandomInt) }
func BenchmarkSortInt160000Sort(b *testing.B)   { benchInt(b, 160000, SortInt, initRandomInt) }
func BenchmarkSortInt640000Sort(b *testing.B)   { benchInt(b, 640000, SortInt, initRandomInt) }
func BenchmarkSortInt2560000Sort(b *testing.B)  { benchInt(b, 2560000, SortInt, initRandomInt) }
func BenchmarkSortInt10240000Sort(b *testing.B) { benchInt(b, 10240000, SortInt, initRandomInt) }

// func BenchmarkSortDouble10000Sort(b *testing.B) { benchDouble(b, 10000, SortDouble, initRandomDouble) }
// func BenchmarkSortDouble40000Sort(b *testing.B) { benchDouble(b, 40000, SortDouble, initRandomDouble) }
// func BenchmarkSortDouble160000Sort(b *testing.B) {
// 	benchDouble(b, 160000, SortDouble, initRandomDouble)
// }
// func BenchmarkSortDouble640000Sort(b *testing.B) {
// 	benchDouble(b, 640000, SortDouble, initRandomDouble)
// }
// func BenchmarkSortDouble2560000Sort(b *testing.B) {
// 	benchDouble(b, 2560000, SortDouble, initRandomDouble)
// }
// func BenchmarkSortDouble10240000Sort(b *testing.B) {
// 	benchDouble(b, 10240000, SortDouble, initRandomDouble)
// }
// func BenchmarkSortInt10000Stable(b *testing.B)    { benchInt(b, 10000, StableInt, initRandomInt) }
// func BenchmarkSortInt40000Stable(b *testing.B)    { benchInt(b, 40000, StableInt, initRandomInt) }
// func BenchmarkSortInt160000Stable(b *testing.B)   { benchInt(b, 160000, StableInt, initRandomInt) }
// func BenchmarkSortInt640000Stable(b *testing.B)   { benchInt(b, 640000, StableInt, initRandomInt) }
// func BenchmarkSortInt2560000Stable(b *testing.B)  { benchInt(b, 2560000, StableInt, initRandomInt) }
// func BenchmarkSortInt10240000Stable(b *testing.B) { benchInt(b, 10240000, StableInt, initRandomInt) }
// func BenchmarkSortDouble10000Stable(b *testing.B) {
// 	benchDouble(b, 10000, StableDouble, initRandomDouble)
// }
// func BenchmarkSortDouble40000Stable(b *testing.B) {
// 	benchDouble(b, 40000, StableDouble, initRandomDouble)
// }
// func BenchmarkSortDouble160000Stable(b *testing.B) {
// 	benchDouble(b, 160000, StableDouble, initRandomDouble)
// }
// func BenchmarkSortDouble640000Stable(b *testing.B) {
// 	benchDouble(b, 640000, StableDouble, initRandomDouble)
// }
// func BenchmarkSortDouble2560000Stable(b *testing.B) {
// 	benchDouble(b, 2560000, StableDouble, initRandomDouble)
// }
// func BenchmarkSortDouble10240000Stable(b *testing.B) {
// 	benchDouble(b, 10240000, StableDouble, initRandomDouble)
// }
