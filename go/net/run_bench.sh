#!/bin/bash

echo "[Go][Benchmark] http server"

# http service case
PORT=8080
go run socket.go $PORT & echo $! > tmp && gopid=$(tail -n 1
tmp)
wrk --latency -t4 -c200 -d8s "http://127.0.0.1:${PORT}"

echo

kill $gopid