package main

import (
	"bytes"
	"strings"
	"testing"
)

var someBytes = []byte("some bytes sdljlk jsklj3lkjlk djlkjw")

var sinkS string

func benchmarkBuilder(b *testing.B, f func(b *testing.B, numWrite int, grow bool)) {
	b.Run("1Write_NoGrow", func(b *testing.B) {
		b.ReportAllocs()
		f(b, 1, false)
	})
	b.Run("3Write_NoGrow", func(b *testing.B) {
		b.ReportAllocs()
		f(b, 3, false)
	})
	b.Run("3Write_Grow", func(b *testing.B) {
		b.ReportAllocs()
		f(b, 3, true)
	})
}

func BenchmarkBuildString_Builder(b *testing.B) {
	benchmarkBuilder(b, func(b *testing.B, numWrite int, grow bool) {
		for i := 0; i < b.N; i++ {
			var buf strings.Builder
			if grow {
				buf.Grow(len(someBytes) * numWrite)
			}
			for i := 0; i < numWrite; i++ {
				buf.Write(someBytes)
			}
			sinkS = buf.String()
		}
	})
}

func BenchmarkBuildString_WriteString(b *testing.B) {
	someString := string(someBytes)
	benchmarkBuilder(b, func(b *testing.B, numWrite int, grow bool) {
		for i := 0; i < b.N; i++ {
			var buf strings.Builder
			if grow {
				buf.Grow(len(someString) * numWrite)
			}
			for i := 0; i < numWrite; i++ {
				buf.WriteString(someString)
			}
			sinkS = buf.String()
		}
	})
}

func BenchmarkBuildString_ByteBuffer(b *testing.B) {
	benchmarkBuilder(b, func(b *testing.B, numWrite int, grow bool) {
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			if grow {
				buf.Grow(len(someBytes) * numWrite)
			}
			for i := 0; i < numWrite; i++ {
				buf.Write(someBytes)
			}
			sinkS = buf.String()
		}
	})
}

func BenchmarkBuildString_ConcatString(b *testing.B) {
	someString := string(someBytes)
	benchmarkBuilder(b, func(b *testing.B, numWrite int, grow bool) {
		for i := 0; i < b.N; i++ {
			var buf string
			for i := 0; i < numWrite; i++ {
				buf = buf + someString
			}
			sinkS = buf
		}
	})
}

func BenchmarkBuildString_ConcatBytes(b *testing.B) {
	benchmarkBuilder(b, func(b *testing.B, numWrite int, grow bool) {
		for i := 0; i < b.N; i++ {
			var buf []byte
			if grow {
				buf = make([]byte, 0, len(someBytes)*numWrite)
			}
			for i := 0; i < numWrite; i++ {
				buf = append(buf, someBytes...)
			}
			sinkS = string(buf)
		}
	})
}
