# How to benchmark in Go

Write an implementation of `wc` and then benchmark that to see if i can make it go faster

## Running

```sh
go test -bench=. ./wc
```

## Output Overview

```sh
goos: darwin
goarch: arm64
pkg: github.com/griggsca91/gobenchmarkexample/wc
BenchmarkCountLines-8              10000            112378 ns/op
BenchmarkCountWords-8               1509            807683 ns/op
PASS
ok      github.com/griggsca91/gobenchmarkexample/wc     2.615s
....
```

1st Column is the benchmark test that was ran
2nd Column is the number of times that the benchmark loop ran
3rd Column is the amount of time each loop took to run

### Thoughts

That's pretty high, lets see what we can do.  

## Running a specific benchmark

## Comparing Benchmarks

## Odd findings

### Sinks

## References

<https://codingchallenges.fyi/challenges/challenge-wc>
<https://gobyexample.com/testing-and-benchmarking>
<https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go>
<https://pkg.go.dev/testing#hdr-Benchmarks>
