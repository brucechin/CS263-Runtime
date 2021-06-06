#!/bin/bash

# Go tests
cd go/

source run_bench.sh sorting
source run_bench.sh matrixmul
source run_bench.sh DijkstraShortestPath

# ./run_bench.sh roaring
cd roaring && source run_bench.sh real && ..

cd ..
# End of Go tests

# Cpp tests
cd cpp/
source run_bench.sh DijkstraShortestPath
source run_bench.sh matmul
source run_bench.sh cpp_sort
source run_bench.sh CRoaring # bitmap
source run_bench.sh ehttp
cd ..
#end of Cpp tests

# Rs tests
cd rust/
source run_bench.sh graph
source run_bench.sh matrixmul  
source run_bench.sh sorting 
cd ..


#http server should be tested manually with two consoles.
cd go/ 
source run_bench.sh net

cd ../cpp/

cd ../rust/
source run_bench.sh simple-web-server
cd ..
