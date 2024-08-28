# How to benchmark in Go

I'm going to benchmark the [strings.Builder](https://pkg.go.dev/strings#Builder) against normal concatenation and then try to make an optimized version of it (though potentially unsafe)

## Running

```sh
go test -bench=.
```

## Output Overview

```sh
goos: darwin
goarch: arm64
pkg: github.com/griggsca91/stringbuilder-benchmark
BenchmarkBuildString_Builder/1Write_NoGrow-8            35288278                29.25 ns/op           48 B/op          1 allocs/op
BenchmarkBuildString_Builder/3Write_NoGrow-8            10506189               115.5 ns/op           336 B/op          3 allocs/op
BenchmarkBuildString_Builder/3Write_Grow-8              29229705                42.11 ns/op          112 B/op          1 allocs/op
BenchmarkBuildString_WriteString/1Write_NoGrow-8        41453700                29.38 ns/op           48 B/op          1 allocs/op
....
```

1st Column is the benchmark test that was ran
2nd Column is the number of times that the benchmark loop ran
3rd Column is the amount of time each loop took to run
4th Column is the number of bytes allocated in each operation

## Running a specific benchmark

## Comparing Benchmarks

## Odd findings

### Sinks

## References

<https://gobyexample.com/testing-and-benchmarking>
<https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go>
<https://pkg.go.dev/testing#hdr-Benchmarks>
