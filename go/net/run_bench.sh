#!/bin/bash

echo "[Go][Benchmark] http server"

# http service case
PORT=8082

go run socket.go $PORT 

echo
