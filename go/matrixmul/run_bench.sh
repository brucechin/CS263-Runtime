#!/bin/bash

echo "[Go][Benchmark] matrixmul"
go test -bench . -benchmem 
echo