# Go Profiling

## Enable profile in code
```
# use http server for profile
import _ "net/http/pprof"
```


## Install Analyze Tool and Depenencies
```
go get -u github.com/google/pprof

# Mac only
brew install graphviz
```

## Analyze
```
go tool pprof  http://localhost:6060/debug/pprof/heap
```

## Build docker image
```
docker build . -t afgo/basic-app:v0.0.2 -t afgo/basic-app:latest
```

## Reference
* [Profiling Go programs with pprof](https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/)
* [Go: The Complete Guide to Profiling Your Code](https://hackernoon.com/go-the-complete-guide-to-profiling-your-code-h51r3waz)
* [Reserve compute resources](https://kubernetes.io/docs/tasks/administer-cluster/reserve-compute-resources/)
* [Node pressure eviction](https://kubernetes.io/docs/concepts/scheduling-eviction/node-pressure-eviction/)
## Notes
### Flat Vs Cumulative
Flat means only by this function, while cum (cumulative) means by this function and functions called down the stack.


