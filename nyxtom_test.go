package main

import (
	"testing"

	"github.com/nyxtom/mphf"
)

func Benchmark_nyxtomNoBuild(b *testing.B) {
	m := make(map[string]int)
	for k, v := range words {
		m[v] = k
	}
	table := mphf.Create(m)
	length := len(words)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		offset := i % length
		p := table.Lookup(words[offset])
		if p == nil {
			b.Fatal("missing key")
		}
		if *p != offset {
			b.Fatal("bad result index")
		}
	}
}
