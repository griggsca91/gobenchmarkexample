# How to benchmark in Go

Write an implementation of `wc` and then benchmark that to see if i can make it go faster

## Running

```sh
go test -bench=. -benchmem ./wc
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

## Get CPU Profile

```sh
go test -bench=. -benchmem ./wc -cpuprofile=cpuprofile.out
go tool pprof cpuprofile.out
```

use `top` to show the top functions
use `list <function>` to get a more detailed version

```sh
go tool pprof cpuprof.out
File: wc.test
Type: cpu
Time: Aug 29, 2024 at 1:39pm (EDT)
Duration: 2.52s, Total samples = 2.15s (85.23%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 2.12s, 98.60% of 2.15s total
Dropped 24 nodes (cum <= 0.01s)
      flat  flat%   sum%        cum   cum%
     0.99s 46.05% 46.05%      1.02s 47.44%  github.com/griggsca91/gobenchmarkexample/wc.CountLines (inline)
     0.53s 24.65% 70.70%      1.10s 51.16%  github.com/griggsca91/gobenchmarkexample/wc.CountWords (inline)
     0.53s 24.65% 95.35%      0.56s 26.05%  github.com/griggsca91/gobenchmarkexample/wc.isWhitespace (inline)
     0.07s  3.26% 98.60%      0.07s  3.26%  runtime.asyncPreempt
         0     0% 98.60%      1.02s 47.44%  github.com/griggsca91/gobenchmarkexample/wc.BenchmarkCountLines
         0     0% 98.60%      1.10s 51.16%  github.com/griggsca91/gobenchmarkexample/wc.BenchmarkCountWords
         0     0% 98.60%      0.02s  0.93%  runtime.systemstack
         0     0% 98.60%      2.12s 98.60%  testing.(*B).launch
         0     0% 98.60%      2.12s 98.60%  testing.(*B).runN
(pprof) list CountLines
Total: 2.15s
ROUTINE ======================== github.com/griggsca91/gobenchmarkexample/wc.BenchmarkCountLines in /Users/chris/Documents/go-benchmark/wc/wc_test.go
         0      1.02s (flat, cum) 47.44% of Total
         .          .     83:func BenchmarkCountLines(t *testing.B) {
         .          .     84:   var result int
         .          .     85:   for i := 0; i < t.N; i++ {
         .      1.02s     86:           result = CountLines(testContent)
         .          .     87:   }
         .          .     88:   sink = result
         .          .     89:}
         .          .     90:
         .          .     91:func BenchmarkCountWords(t *testing.B) {
ROUTINE ======================== github.com/griggsca91/gobenchmarkexample/wc.CountLines in /Users/chris/Documents/go-benchmark/wc/wc.go
     990ms      1.02s (flat, cum) 47.44% of Total
         .          .      7:func CountLines(content []byte) int {
         .          .      8:   var count int
      40ms       40ms      9:   for i := 0; i < len(content); i++ {
     950ms      980ms     10:           if content[i] == '\n' {
         .          .     11:                   count++
         .          .     12:           }
         .          .     13:   }
         .          .     14:   return count
         .          .     15:}
(pprof)
```

## Running a specific benchmark

## Comparing Benchmarks

## Odd findings

### Sinks

## References

<https://codingchallenges.fyi/challenges/challenge-wc>
<https://gobyexample.com/testing-and-benchmarking>
<https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go>
<https://pkg.go.dev/testing#hdr-Benchmarks>
