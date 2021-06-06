#!/bin/bash
test_name=ehttp

echo "[Cpp][Benchmark] ${test_name}"
echo "Compiling..."
make && make test
echo 
echo "Compilatio done."
echo 
echo "run command 'wrk --latency -t4 -c200 -d8s http://127.0.0.1:3456' from another terminal"
echo "After wrk report is generated, please manually stop this running program"
./output/test/hello_server 3456

echo