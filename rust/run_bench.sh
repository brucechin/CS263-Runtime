#!/bin/bash

for testcase in $@
do
cd $testcase
source run_bench.sh
cd ..
done


