CS263-Runtime
===
## Team member 

Guyue Huang and Lianke Qin

## Targets

Compare the performance difference among static language including C++, Go and Rust by profiling 5-10 programs under different compiler optimization levels and input sizes.

1. Graph algorith: Dijkstra single-source shortest path.
2. matrix multiplication
3. sorting algorithm
4. simple HTTP request handle throughput
5. Bitmap set operations.


## How to reproduce results
A top-level script  runs all tests.
```bash
source run_bench_top.sh
```
The test programs for three languages are located under the folder with the language name. Under all test folders, there is a script ```run_bench.sh```,  so tests can be run individually, as well. For example,
```bash
# run individual test, e.g. matrix multiplication of C++
cd cpp/matmul && source run_bench.sh
```

## Acknowledgements
We refer to several code resources during this project, including the links below.
* Go by example. https://gobyexample.com 
* Your basic graph: Golang library of basic graph algorithms. https://github.com/yourbasic/graph
* Hyper: a fast and correct HTTP implementation for Rust. https://github.com/hyperium/hyper
* The Roaring Bitmap project. http://roaringbitmap.org 
