# funk-bench
Thought I'd test the performance of funk (github.com/thoas/go-funk). 

This is that.

## Results

```
➜  funk-bench git:(master) ✗ go test -bench .
testing: warning: no tests to run
BenchmarkContains/Idiomatic-8               1000           1818405 ns/op
BenchmarkContains/funk-8                       2         580595289 ns/op
PASS
ok      github.com/christian-blades-cb/funk-bench       3.868s
```

## How to run?

Clone and run the following from the repo root: `go get -t; go test -bench .`
