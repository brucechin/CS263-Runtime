// file: sort_test.go
// 		benchmark sorting algorithms in go libraries
// 		run `go test -bench . -benchmem` in the same folder
//
// 		reference code golang/go/src/compress/flate/*
//

package compressing

import (
	"bytes"
	"compress/flate"
	"io"
	"os"
	"runtime"
	"testing"
)

// RFC: golang/go/src/compress/flate/reader_test
func benchDecompress(b *testing.B, inputFile string, level int, reader func(io.Reader) io.ReadCloser) {
	buf, err := os.ReadFile(inputFile)
	if err != nil {
		b.Fatal(err)
	}
	if len(buf) == 0 {
		b.Fatalf("test file %q has no data", inputFile)
	}

	b.ReportAllocs()
	b.StopTimer()
	b.SetBytes(int64(len(buf)))

	compressed := new(bytes.Buffer)
	w, err := flate.NewWriter(compressed, level)
	if err != nil {
		b.Fatal(err)
	}

	io.Copy(w, bytes.NewReader(buf))
	w.Close()

	buf1 := compressed.Bytes()
	buf, compressed, w = nil, nil, nil
	runtime.GC()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		io.Copy(io.Discard, reader(bytes.NewReader(buf1)))
	}
}

func benchCompress(b *testing.B, inputFile string, level int, writer func(io.Writer, int) (*flate.Writer, error)) {
	buf, err := os.ReadFile(inputFile)
	if err != nil {
		b.Fatal(err)
	}
	if len(buf) == 0 {
		b.Fatalf("test file %q has no data", inputFile)
	}

	b.StopTimer()
	b.SetBytes(int64(len(buf)))

	w, err := writer(io.Discard, level)
	if err != nil {
		b.Fatal(err)
	}
	runtime.GC()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w.Reset(io.Discard)
		w.Write(buf)
		w.Close()
	}
}

// generated by python code;
// source
//
// inputs={'Input1':'compress.txt'}
// algorithms={'Compress':'flate.NewWriter', 'Decompress':'flate.NewReader'}
// for level in range(4): # level (0-3)
//   for file in inputs.keys():
//     for algo in algorithms.keys():
//       ans = "func Benchmark%s%sLevel%s(b *testing.B) { bench%s(b, \"%s\", %s, %s) }"%(algo, file, level, algo, inputs[file], level, algorithms[algo])
// 	  print(ans)

func BenchmarkCompressInput1Level0(b *testing.B) {
	benchCompress(b, "compress.txt", 0, flate.NewWriter)
}
func BenchmarkDecompressInput1Level0(b *testing.B) {
	benchDecompress(b, "compress.txt", 0, flate.NewReader)
}
func BenchmarkCompressInput1Level1(b *testing.B) {
	benchCompress(b, "compress.txt", 1, flate.NewWriter)
}
func BenchmarkDecompressInput1Level1(b *testing.B) {
	benchDecompress(b, "compress.txt", 1, flate.NewReader)
}
func BenchmarkCompressInput1Level2(b *testing.B) {
	benchCompress(b, "compress.txt", 2, flate.NewWriter)
}
func BenchmarkDecompressInput1Level2(b *testing.B) {
	benchDecompress(b, "compress.txt", 2, flate.NewReader)
}
func BenchmarkCompressInput1Level3(b *testing.B) {
	benchCompress(b, "compress.txt", 3, flate.NewWriter)
}
func BenchmarkDecompressInput1Level3(b *testing.B) {
	benchDecompress(b, "compress.txt", 3, flate.NewReader)
}
