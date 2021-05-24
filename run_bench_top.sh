#!/bin/bash

# Go tests
cd go/

./run_bench.sh sorting
./run_bench.sh compress
./run_bench.sh net
./run_bench.sh matrixmul
./run_bench.sh DijkstraShortestPath
./run_bench.sh roaring
cd roaring && ./run_bench.sh real && ..

cd ..
# End of Go tests

# Cpp tests
cd cpp/
./run_bench.sh compression-algorithms
./run_bench.sh DijkstraShortestPath
./run_bench.sh ehttp
./run_bench.sh matmul
cd ..
#end of Cpp tests

# Rs tests
