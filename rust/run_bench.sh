#!/bin/bash

for testcase in $@
cd $testcase
./run_bench.sh
cd ..
done


