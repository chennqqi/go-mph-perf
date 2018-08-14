package main

import (
	"testing"

	"github.com/cespare/mph"
)

func Benchmark_cespareNoBuild(b *testing.B) {
	t := mph.Build(words)
	length := len(words)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		offset := i % length
		n, ok := t.Lookup(words[offset])
		if !ok {
			b.Fatal("missing key")
		}
		if n != uint32(offset) {
			b.Fatal("bad result index")
		}
	}
}
