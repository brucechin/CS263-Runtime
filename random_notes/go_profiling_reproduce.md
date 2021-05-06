# Go profiling tutorial reproduced

## Procedure
* havlak1.go: a single-file version of the original Hindt's program (repo/src/havlak/go)
    * removing all prints except the last one
    * adding code to use pprof for cpuprofile and memprofile
* havlakpro1.go: a single-file version of the original Hindt's program (repo/src/havlak/go_pro)
    * It is supposed to be an optimized version, doing No.1&2 of the following opts.
    * Its optimizations include
        * inline some short functions 
        * container iterating styles. avoid indirection
* havlak2.go
    * on top of havlakpro1.go, change the inefficient global *maps* to *lists*
* havlak3.go 
    * on top of havlak2, apply the memory cache trick mentioned in opt3 (next part)

## Notes about Optimization
* \[Time\]\[Data structure\] 
    * Diagnosed by cpuprofile hottest instr. in terms of flat time
    * What-to-do: in the DFS algorithm, modify the structure used to find node from *map* to *list*, because every node has UID.
* \[Mem\]\[Data structure\]
    * Diagnosed by memprofile the instr. that allocates memory most
    * What-to-do: modify all global maps (used to preserve references to nodes) to slices (i.e. lists), again, because the node has UID.
* \[Mem\]\[GC\]
    * Diagnosed by memprofile: in iterative call to FindLoops, every time it allocs new bookkeep storage. It can be reused, not wait until GC.
    * What-to-do: use cache structure to reuse the memory as stated in the tutorial

## How to reproduce
./all_prof.sh

Expected results: see logs under the log/ dir. Expect in terms of performance, havlak1 < havlakpro1 < havlak2 

Here are the results of one run collected.

User | System | Real | Mem | Testcase
- | - | - | - | -
37.85 | 0.64 | 20.90 | 462484kB | ./havlak1
24.04 | 0.43 | 15.09 | 308048kB | ./havlakpro1
8.29 | 0.18 | 5.34 | 194860kB | ./havlak2
4.70 | 0.09 | 3.60 | 189552kB | ./havlak3


The data structure opt. contributes 3x speedup, and the cache trick to reduce GC contributes another 1.5-2x. 

## TODO
A few issues need to be checked, but not very big issues. 
* Observe user time is shorter than real cpu time. There must be some multithreading done by the compiler. However, passing `-p 1` to `go build` does not change results. How to turn off it?
* (Since this is irrelevant to profiling, not implemented) Reproduce the final tricks in the tutorial, according to their results, it will give another 3~4x speedup. 
* Why the test results vary each run, around +/- 3% in terms of time. Frequency scaling?  

## Memo
run `./all_prof.sh 1` will give detailed cpu profile results under log/, and can be interpreted using pprof.