Go benchmark
===
Benchmarks
* compress
* dijkstra algorithm for graph SSSP
* matrixmul (serial, row-wise parallel, blocked parallel)
* sorting (quicksort, mergesort)
* socket 
* bitmap

# How to run
Enter each folder and run ``go test -bench . -benchmem``. The report from go-test is printed to the shell.

Only for socket, no need to use go-test. ``cd net`` and use 

``go run socket.go \[port\]
``
The 4-digit port is optional. By default it is 8082.
