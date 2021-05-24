#!/bin/bash

echo "[Go][Benchmark] sorting"
go test -bench . -benchmem 
echo