#!/bin/bash

echo "[Go][Benchmark] Dijkstra SSSP"
go test -bench . -benchmem 
echo