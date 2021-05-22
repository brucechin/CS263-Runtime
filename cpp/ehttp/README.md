# ehttp 
![C/C++ CI](https://github.com/hongliuliao/ehttp/workflows/C/C++%20CI/badge.svg?branch=master)
[![Build Status](https://travis-ci.org/hongliuliao/ehttp.svg?branch=master)](https://travis-ci.org/hongliuliao/ehttp)
[![codecov.io](http://codecov.io/github/hongliuliao/ehttp/coverage.svg?branch=master)](http://codecov.io/github/hongliuliao/ehttp?branch=master)


## Feature
* Base on linux epoll
* Multi-thread model

## Build && Test
```
 make && make test && ./output/test/hello_server 3456
 wrk --latency -t4 -c200 -d8s http://127.0.0.1:3456
```

