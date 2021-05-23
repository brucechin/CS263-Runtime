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

func BenchmarkMatrixMulInt1000RowWise(b *testing.B) {
	benchInt(b, 1000, matrixMulRowWise)
}

func BenchmarkMatrixMulInt1000Blocked(b *testing.B) {
	benchInt(b, 1000, matrixMulBlocked)
}

//
// algorithms
//

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
	tile := 16

	Arow := size
	Brow := size
	Bcol := size

	// nWorkers := ((Arow + tile - 1) / tile) * ((Bcol + tile - 1) / tile)

	var waitComplete sync.WaitGroup

	for rowStart := 0; rowStart < Arow; rowStart += tile {
		for colStart := 0; colStart < Bcol; colStart += tile {

			waitComplete.Add(1)

			go func(rowStartLocal int, colStartLocal int) {

				defer waitComplete.Done()

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
			}(rowStart, colStart)
		}
	}
	waitComplete.Wait()
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
