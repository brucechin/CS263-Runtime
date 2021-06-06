#!/bin/bash
test_name=Bitmap

echo "[Cpp][Benchmark] ${test_name}"

echo "Compiling..."

./amalgamation.sh
mkdir -p build 
cd build
cmake ..
cmake --build .

echo "Building finished. Running benchmark ..."
echo "Dataset 1:"
./real_bitmaps_benchmark ../benchmarks/realdata/census-income
echo "Dataset 2:"
./real_bitmaps_benchmark ../benchmarks/realdata/census1881
echo "Dataset 3:"
./real_bitmaps_benchmark ../benchmarks/realdata/weather_sept_85

cd ..

echo
