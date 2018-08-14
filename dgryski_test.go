package main

import (
	//"github.com/dgryski/go-boomphf"
	"testing"

	"github.com/dgryski/go-mph"
)

func Benchmark_dgryskiNoBuild(b *testing.B) {
	table := mph.New(words)
	length := len(words)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		offset := i % length
		n := table.Query(words[offset])
		if n < 0 {
			b.Fatal("missing key")
		}
		if int(n) != offset {
			b.Fatal("bad result index")
		}
	}
}
