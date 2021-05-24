Benchmarking Results
===

Tests
- Go
[x] compress
[x] sorting
[x] http
[x] graph
[-] matrix-mul

- Cpp
[x] compress
[x] matrix-mul
[x] http
[x] graph
[?] sorting

- Rust
[x] compress
[x] matmul
[x] sorting 
[x] graph
[x] http


# Compress

alice29.txt

## Go

| compress level | 0 | 1 | 2 | 3 |
|-|-|-|-|-|
| compress | 0.0082ms | 2.6ms | 4.4ms | 5.1ms |
| decompress | 0.039ms | 2.3ms | 2.1ms | 2.1ms|

## Cpp

- Encode 5.64ms
- Decode 190ms

## Rust
Compress 13.41ms
Decompress 3.99ms


# Sorting

## Go

| array size |1e4 | 4e4 | 16e4 | 64e4 | 256e4 | 1024e4|
|-|-|-|-|-|-|-|
| Int32 quicksort | 1.21ms | 5.22ms | 22.7ms | 96.2ms | 412ms| 1738ms |
| Int32 mergesort | 1.43ms | 5.83ms | 22.8ms | 90.8ms | 255ms | 1479ms |
| FP64 quicksort | 1.97ms | 9.01ms | 41.1ms | 183ms | 799ms | 3419ms | 
| FP64 mergesort | 4.60ms | 22.8ms | 112ms | 532ms | 2478ms | 11422ms|

## Rust
FP64?
```sh
algorithm n comparisons time
quicksort 10000 209480 0.001685205
heapsort -- length 10000 -- 0.002430233 seconds
quicksort 10000 213674 0.001356819
heapsort -- length 10000 -- 0.001540946 seconds
quicksort 10000 199461 0.001043446
heapsort -- length 10000 -- 0.001549172 seconds
quicksort 10000 215069 0.001053229
heapsort -- length 10000 -- 0.001550263 seconds
quicksort 10000 210452 0.001055747
heapsort -- length 10000 -- 0.001546407 seconds
quicksort 40000 1010491 0.004822627
heapsort -- length 40000 -- 0.007951113 seconds
quicksort 40000 1009626 0.00496106
heapsort -- length 40000 -- 0.008323976 seconds
quicksort 40000 988543 0.004854179
heapsort -- length 40000 -- 0.008336881 seconds
quicksort 40000 984158 0.004797956
heapsort -- length 40000 -- 0.00834971 seconds
quicksort 40000 997081 0.004859035
heapsort -- length 40000 -- 0.007992629 seconds
quicksort 160000 4440040 0.02179013
heapsort -- length 160000 -- 0.0385834 seconds
quicksort 160000 4387720 0.017668769
heapsort -- length 160000 -- 0.036957814 seconds
quicksort 160000 4844290 0.018414431
heapsort -- length 160000 -- 0.037086143 seconds
quicksort 160000 4369427 0.017608558
heapsort -- length 160000 -- 0.038577139 seconds
quicksort 160000 4576812 0.018090142
heapsort -- length 160000 -- 0.03694934 seconds
quicksort 640000 20311749 0.080397075
heapsort -- length 640000 -- 0.196022673 seconds
quicksort 640000 20208539 0.079760247
heapsort -- length 640000 -- 0.196053362 seconds
quicksort 640000 20271272 0.079766031
heapsort -- length 640000 -- 0.197531722 seconds
quicksort 640000 23296606 0.084546757
heapsort -- length 640000 -- 0.196011329 seconds
quicksort 640000 19810405 0.078407861
heapsort -- length 640000 -- 0.20516382 seconds
quicksort 2560000 92949480 0.356570101
heapsort -- length 2560000 -- 1.067864577 seconds
quicksort 2560000 92669857 0.353061229
heapsort -- length 2560000 -- 1.072109616 seconds
quicksort 2560000 92974133 0.357309812
heapsort -- length 2560000 -- 1.06730618 seconds
quicksort 2560000 93451940 0.352126864
heapsort -- length 2560000 -- 1.0682497 seconds
quicksort 2560000 91422712 0.353306754
heapsort -- length 2560000 -- 1.072257001 seconds
quicksort 10240000 410214601 1.555183392
heapsort -- length 10240000 -- 6.421657276 seconds
quicksort 10240000 421918335 1.574710361
heapsort -- length 10240000 -- 6.420579885 seconds
quicksort 10240000 414527545 1.567240243
heapsort -- length 10240000 -- 6.43026879 seconds
quicksort 10240000 413395241 1.567898861
heapsort -- length 10240000 -- 6.43167655 seconds
quicksort 10240000 395809507 1.53187655
heapsort -- length 10240000 -- 6.440489583 seconds
```

| array size |1e4 | 4e4 | 16e4 | 64e4 | 256e4 | 1024e4|
|-|-|-|-|-|-|-|


# Dijkstra SSSP

Average 1/100 edges


| graph size | 1e4 | 4e4 | 16e4|
|-|-|-|-|
| **Go** time | 25ms | 425ms | 7142ms |
| **Cpp** time |0.198sec | 3.42sec | 69.2sec |
| **Rust** time | 126ms | 1.70sec | 33.6sec |
| **Go** mem | 1.3MB | 5.6MB | 24MB |



# Matmul

| matrix shape | 500 | 1000 | 2000 | 4000 | 8000
|-|-|-|-|-|-|
| **cpp** | 0.22s | 1.56s | 12.23s | 98.34s | |
| **go** (serial) | 576ms | 4.33s | 40.2s | |
| **go** (concurrent) | 26ms | 257ms | 3054ms | 47s ||

## Rust

```sh
Torch -- length 500 -- 0.015731704 seconds
Torch -- length 500 -- 0.003839715 seconds
Torch -- length 500 -- 0.003705638 seconds
Torch -- length 500 -- 0.003779632 seconds
Torch -- length 500 -- 0.003673097 seconds
Torch -- length 1000 -- 0.008562981 seconds
Torch -- length 1000 -- 0.005970103 seconds
Torch -- length 1000 -- 0.006609141 seconds
Torch -- length 1000 -- 0.004770905 seconds
Torch -- length 1000 -- 0.007574518 seconds
Torch -- length 2000 -- 0.0359662 seconds
Torch -- length 2000 -- 0.031215053 seconds
Torch -- length 2000 -- 0.029853522 seconds
Torch -- length 2000 -- 0.027443492 seconds
Torch -- length 2000 -- 0.025899107 seconds
Torch -- length 4000 -- 0.159750684 seconds
Torch -- length 4000 -- 0.152678472 seconds
Torch -- length 4000 -- 0.147233991 seconds
Torch -- length 4000 -- 0.143422171 seconds
Torch -- length 4000 -- 0.14109469 seconds
Torch -- length 8000 -- 0.801356911 seconds
Torch -- length 8000 -- 0.777312168 seconds
Torch -- length 8000 -- 0.77974416 seconds
Torch -- length 8000 -- 0.779772821 seconds
Torch -- length 8000 -- 0.777966411 seconds
```

# Http

## Go 
```sh
Running 8s test @ http://127.0.0.1:2345
  4 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   819.99us    3.89ms 119.39ms   98.78%
    Req/Sec    70.58k     7.70k   75.10k    96.25%
  Latency Distribution
     50%  324.00us
     75%  619.00us
     90%    1.33ms
     99%    5.99ms
  2247291 requests in 8.00s, 341.81MB read
Requests/sec: 280753.10
Transfer/sec:     42.70MB
```

## Cpp
```sh
Running 8s test @ http://127.0.0.1:2346
  4 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     7.72ms    4.71ms  47.41ms   86.04%
    Req/Sec     5.11k     2.82k   12.07k    63.90%
  Latency Distribution
     50%    6.28ms
     75%    7.37ms
     90%   12.65ms
     99%   28.23ms
  161247 requests in 8.10s, 21.68MB read
  Non-2xx or 3xx responses: 161247
Requests/sec:  19908.65
Transfer/sec:      2.68MB
```

## Rust
Hyper
```sh
Running 8s test @ http://127.0.0.1:8082
  4 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   312.66us  116.62us   1.82ms   62.43%
    Req/Sec   106.42k     8.78k  113.95k    82.50%
  Latency Distribution
     50%  333.00us
     75%  403.00us
     90%  447.00us
     99%  589.00us
  3387982 requests in 8.00s, 284.33MB read
Requests/sec: 423311.49
Transfer/sec:     35.53MB
```

Rocket
(build failed..)

# Bitmap

