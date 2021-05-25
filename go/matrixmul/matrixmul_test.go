package matrixmul

import (
	"math/rand"
	"sync"
	"testing"
)

func benchInt(b *testing.B, size int, f func(int, [][]int, [][]int, [][]int)) {
	matrixA := make([][]int, size)
	matrixB := make([][]int, size)
	result := make([][]int, size)
	for i := 0; i < size; i++ {
		matrixA[i] = make([]int, size)
		matrixB[i] = make([]int, size)
		result[i] = make([]int, size)
	}

	randomMatrix(size, matrixA)
	randomMatrix(size, matrixB)
	zeroMatrix(size, result)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		f(size, matrixA, matrixB, result)
	}
}

// test problem sizes 500, 1000, 2000, 4000, 8000

func BenchmarkMatrixMulInt500Serial(b *testing.B) {
	benchInt(b, 500, matrixMulSerial)
}
func BenchmarkMatrixMulInt1000Serial(b *testing.B) {
	benchInt(b, 1000, matrixMulSerial)
}

func BenchmarkMatrixMulInt2000Serial(b *testing.B) {
	benchInt(b, 2000, matrixMulSerial)
}
func BenchmarkMatrixMulInt4000Serial(b *testing.B) {
	benchInt(b, 4000, matrixMulSerial)
}
func BenchmarkMatrixMulInt8000Serial(b *testing.B) {
	benchInt(b, 8000, matrixMulSerial)
}

// func BenchmarkMatrixMulInt500RowWise(b *testing.B) {
// 	benchInt(b, 500, matrixMulRowWise)
// }
// func BenchmarkMatrixMulInt1000RowWise(b *testing.B) {
// 	benchInt(b, 1000, matrixMulRowWise)
// }
// func BenchmarkMatrixMulInt2000RowWise(b *testing.B) {
// 	benchInt(b, 2000, matrixMulRowWise)
// }
// func BenchmarkMatrixMulInt4000RowWise(b *testing.B) {
// 	benchInt(b, 4000, matrixMulRowWise)
// }
// func BenchmarkMatrixMulInt8000RowWise(b *testing.B) {
// 	benchInt(b, 8000, matrixMulRowWise)
// }

func BenchmarkMatrixMulInt500Blocked(b *testing.B) {
	benchInt(b, 500, matrixMulBlocked)
}
func BenchmarkMatrixMulInt1000Blocked(b *testing.B) {
	benchInt(b, 1000, matrixMulBlocked)
}
func BenchmarkMatrixMulInt2000Blocked(b *testing.B) {
	benchInt(b, 2000, matrixMulBlocked)
}
func BenchmarkMatrixMulInt4000Blocked(b *testing.B) {
	benchInt(b, 4000, matrixMulBlocked)
}
func BenchmarkMatrixMulInt8000Blocked(b *testing.B) {
	benchInt(b, 8000, matrixMulBlocked)
}

//
// algorithms
//

// 0. sequential algorithm
func matrixMulSerial(size int, matrixA [][]int, matrixB [][]int, result [][]int) {
	Arow := size
	Brow := size
	Bcol := size
	for ii := 0; ii < Arow; ii++ {
		for jj := 0; jj < Bcol; jj++ {
			// fmt.Printf("i %d j %d\n", ii, jj)
			var res = 0
			for kk := 0; kk < Brow; kk++ {
				res += matrixA[ii][kk] * matrixB[kk][jj]
			}
			result[ii][jj] = res
		}
	}
}

// 1. row-wise parallel algorithms using go routines
func matrixMulRowWise(size int, matrixA [][]int, matrixB [][]int, result [][]int) {
	Arow := size
	Brow := size
	Bcol := size

	var waitComplete sync.WaitGroup

	for ii := 0; ii < Arow; ii++ {
		waitComplete.Add(1)

		go func(row int) {
			defer waitComplete.Done()

			for jj := 0; jj < Bcol; jj++ {
				// fmt.Printf("i %d j %d\n", ii, jj)
				var res = 0
				for kk := 0; kk < Brow; kk++ {
					res += matrixA[row][kk] * matrixB[kk][jj]
				}
				result[row][jj] = res
			}
		}(ii)
	}
	waitComplete.Wait()
}

// 2. blocked parallel algorithms using go routines
func matrixMulBlocked(size int, matrixA [][]int, matrixB [][]int, result [][]int) {
	tile := 32

	Arow := size
	Brow := size
	Bcol := size

	// nWorkers := ((Arow + tile - 1) / tile) * ((Bcol + tile - 1) / tile)

	// var waitComplete sync.WaitGroup

	for rowStart := 0; rowStart < Arow; rowStart += tile {
		for colStart := 0; colStart < Bcol; colStart += tile {

			// waitComplete.Add(1)

			// go func(rowStartLocal int, colStartLocal int) {

			// defer waitComplete.Done()
			rowStartLocal := rowStart
			colStartLocal := colStart

			var rowStop = rowStartLocal + tile
			if rowStop > Arow {
				rowStop = Arow
			}
			var colStop = colStartLocal + tile
			if colStop > Bcol {
				colStop = Bcol
			}
			for i := rowStartLocal; i < rowStop; i++ {
				for j := colStartLocal; j < colStop; j++ {
					var res = 0
					for k := 0; k < Brow; k++ {
						res += matrixA[i][k] * matrixB[k][j]
					}
					result[i][j] = res
				}
			}
			// }(rowStart, colStart)
		}
	}
	// waitComplete.Wait()
}

//
// helper functions
//

func randomMatrix(size int, matrix [][]int) {
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			matrix[row][col] = rand.Intn(10) - 5
		}
	}
}

func zeroMatrix(size int, matrix [][]int) {
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			matrix[row][col] = 0
		}
	}
}
