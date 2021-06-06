#!/bin/bash

echo "[Go][Benchmark] bitmap"

# go test -bench Benchmark -run -
echo

real=${1:-"noreal"}

if [ $real == "real" ] 
then

echo "[Go][Benchmark] bitmap real datasets"

go get github.com/RoaringBitmap/real-roaring-datasets
BENCH_REAL_DATA=1 go test -bench BenchmarkRealData -run -
echo
fi
