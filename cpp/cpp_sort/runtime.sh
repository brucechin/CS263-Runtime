#!/bin/bash
cmake .
make clean
make -j16
# Variables for setting timeout on test calls.
timeoutOccuredExitCode=124
timeoutLength=100

# Construct the output file path.
mkdir -p outputs
fileName=outputs/result
currentTime=$(date "+%s")
outputFileName=$fileName-$currentTime.csv

# Run the script with timeout.
function runScript {
    ./sorting_benchmarks.out $1 $2 data/output.txt $3
}

# Number of elements and algorithms arrays.
files=(data/rand.txt)
elementCount=(10000 40000 160000 640000 2560000 10240000)
algorithms=(qs)

# Loop over the elements and algorithms and run the tests for every combination.


for fileName in "${files[@]}"
do
    echo "-------------------- $fileName --------------------" | tee -a $outputFileName
    for numberOfElements in "${elementCount[@]}"
    do
        for algorithm in "${algorithms[@]}"
        do
            for count in {1..3}
            do
                runScript $numberOfElements $fileName $algorithm | tee -a $outputFileName
            done
        done
    done
done
