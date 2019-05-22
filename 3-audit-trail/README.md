## Experiment with audit trail

## Purpose:

Find the changes of 2 same struct


### Benchmark

```    
goos: linux
goarch: amd64
pkg: github.com/h4ckm03d/golang-playground/3-audit-trail
BenchmarkCompare-4       2000000               773 ns/op             896 B/op         17 allocs/op
BenchmarkCmpDiff-4         50000             29135 ns/op           10522 B/op        176 allocs/op
BenchmarkLibra-4         1000000              1018 ns/op             656 B/op         10 allocs/op
```