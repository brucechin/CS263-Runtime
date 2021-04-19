
## Some more sample programs
From awesome-c++:
1. standard library
2. async event, semaphore, volatile var
3. data struc: bitset, Btree, kdtree....
4. image processing. I think only c++ has this type of app? 
5. serializtion, de-. json? 
6. malloc, alloca 
7. math, physics, hpc kernels
8. primitive algorithms, map,reduce,scan,sort,
9. web
10. VM? container? (I feel it's too complicated)
https://github.com/fffaraz/awesome-cpp#reflection

## Things we can do
- Automate profiling & performance analysis : scripts or projects to automatically run, profile, collect metrics, present, maybe visualize, maybe generate input tests ranging different sizes
    - for Go, based on an example- workflow in https://blog.golang.org/profiling-go-programs, try to automate the common part and run for new programs.
    - find 5 additional programs to do profiling on
    - for other languages, find some similar workflow? compare their profiling support?
- find large open source projects
    - test with various inputs
    - automate and visualize the performance under various input sizes
    - this can be combined with the previous? previous topic is develop workflow for program, this is develop workflow for inputs

- profile and extend the compiler optimizations 
    - i think we can pass this one...

- Investigate and empirically evaluate concurrency (parallel/threading and distribution/messaging) mechanisms extensively 
    - this topic is huge... I am only interested in some specific features, like coroutine in c++, but I have little experience. I think we can also skip this...


## Ideas
### Go profiling 
https://docs.google.com/document/d/1Lvf5QGwuB5yyoBa4OAOfGut8OpzbG7-DlIoJiAprvLo/edit
- metrics :
    - basic: time, memory usage
    - GC cycles? GC performance? (can be tuned by flags)
- practice
    - go profiling tutorial, and a bunch of references in it 
        - basic: aggregate the execution time by functions
        - can dump call graph, each node attached with time and memory
        - can collect traces , memory usage and time for each source line (? execution snapshot)
        - **Overall, after checking this tutorial, I feel that go profiler is a nicely integrated tool, many good functionalities to explore.**


- focus, topics
    - try to find a gradual top-down profiling workflow to find performance bottleneck? 
    - explore some go language features, maybe concurrency? and profile their performance against naive version.
    - maybe we can also study how compiler-time flags affect the performance?

### Rust profiling metrics, practice, targets
- metrics
- practice
- interesting topics

### C++ profiling 
- metrics
  - basically time, memory. 
  - for openmp/parallel programming, things like multithread scale-out is important
- practice
  - basic ones acquired by timing APIs in C++ or linux process supervision command for memory usage.
  - VS tools do hotspot analysis  at granularity of source lines. I know of some reference materials focusing on tuning performance of c++, openMP, llvm. 
- interesting topics
  - new c++20 features?
  - system-level insights, to profile complex applications ( database? graph processing? games? ML jobs? ) and analyze the entire project, what functions/ops are bottleneck? this may be very useful.
  - in terms of sample programs, awesome c++ https://github.com/fffaraz/awesome-cpp
