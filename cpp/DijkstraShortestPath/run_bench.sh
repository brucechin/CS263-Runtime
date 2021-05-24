#!/bin/bash
test_name=compress
test_file=dijkstra2.cpp

echo "[Cpp][Benchmark] ${test_name}"
g++ -O3 -std=c++11 -o test $test_file
./test
echo
