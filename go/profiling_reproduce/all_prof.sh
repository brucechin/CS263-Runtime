#!/bin/sh

mkdir -p log
LOGFILE=log/log
echo "Writing to $LOGFILE"

detailed=${1:-0}

for version in 3 2 pro1 1
do

fname="havlak${version}"

# go compile
go build "${fname}.go"
echo "Done compiling ${fname}"

# go run and supervise with gtime
echo "${fname} "
./xtime "./${fname}" >>  $LOGFILE 2>&1

echo "Done profiling dryrun"

# # run with detailed profiler

# # memprofile is commented. If use memory profile, 
# # the program will only run FindLoop for once instead of 50 originall, so conflict with time profile
# ./${fname} -cpuprofile="log/${fname}.prof" # --memprofile="log/${fname}.mprof" 

# echo "Done generating detailed cpu and mem profiles"

done