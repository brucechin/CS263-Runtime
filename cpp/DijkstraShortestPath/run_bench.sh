#!/bin/bash
test_name=graph
test_file=dijkstra_v3.cpp

echo "[Cpp][Benchmark] ${test_name}"
g++ -O3 -std=c++11 -o test $test_file
./test
echo
