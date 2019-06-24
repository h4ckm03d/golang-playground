# POINTER benchmark

## Description

Benchmark methodology is simple. Use struct `Person` and `Student` as base struct with additional unrelated data from `https://mholt.github.io/json-to-go/` just to make the struct bigger.
The function to compare just passing `Person` as parameter both with and without pointer then create instance `Student` to return it as pointer or not.

## Result

```bash
goos: linux
goarch: amd64
pkg: github.com/h4ckm03d/golang-playground/4-pointer
BenchmarkA-4            100000000               22.7 ns/op             0 B/op          0 allocs/op
BenchmarkB-4            100000000               13.1 ns/op             0 B/op          0 allocs/op
BenchmarkC-4            2000000000               0.27 ns/op            0 B/op          0 allocs/op
BenchmarkAParalel-4     300000000                5.66 ns/op            0 B/op          0 allocs/op
BenchmarkBParalel-4     500000000                2.93 ns/op            0 B/op          0 allocs/op
BenchmarkCParalel-4     2000000000               0.38 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/h4ckm03d/golang-playground/4-pointer 9.041s
```