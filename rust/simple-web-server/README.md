## Benchmarking Tool

I used `wrk` to test those http servers:

```sh
$ wrk --latency -t4 -c200 -d8s http://127.0.0.1:8000
```

## Results

I ran all those tests on:

`Intel Core i7 CPU @ 3.20GHz × 12, 16GB of RAM`

Before run, I created a release build:

```sh
$ cargo build --release
```

### Framework 1: hyper (https://github.com/hyperium/hyper)

_Requests/sec:_ `520718.02` and _Latency:_ `368.31us`

```txt
Running 8s test @ http://127.0.0.1:8082
  4 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   368.31us  327.11us   8.11ms   92.48%
    Req/Sec   131.79k     6.77k  147.18k    77.50%
  Latency Distribution
     50%  295.00us
     75%  390.00us
     90%  595.00us
     99%    1.64ms
  4196014 requests in 8.06s, 352.14MB read
Requests/sec: 520718.02
Transfer/sec:     43.70MB
```


### Framework 2: Rocket (https://github.com/SergioBenitez/Rocket)

_Requests/sec:_ `115231.95` and _Latency:_ `412.75us`

```txt
Running 8s test @ http://127.0.0.1:8000
  4 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   412.75us    1.35ms 204.10ms   99.83%
    Req/Sec    29.20k     1.60k   31.79k    85.94%
  Latency Distribution
     50%  312.00us
     75%  551.00us
     90%  737.00us
     99%    0.99ms
  929792 requests in 8.07s, 129.46MB read
  Socket errors: connect 0, read 929766, write 0, timeout 0
Requests/sec: 115231.95
Transfer/sec:     16.04MB
```