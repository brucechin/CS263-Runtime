#!/bin/bash

echo "[Go][Benchmark] compress"
go test -bench . -benchmem 
echo