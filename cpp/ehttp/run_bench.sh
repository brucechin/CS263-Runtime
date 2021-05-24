#!/bin/bash
test_name=ehttp

echo "[Cpp][Benchmark] ${test_name}"

make && make test && ./output/test/hello_server 3456
wrk --latency -t4 -c200 -d8s http://127.0.0.1:3456

echo