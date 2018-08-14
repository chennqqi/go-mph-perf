package main

import (
	"testing"
)

func Benchmark_map(b *testing.B) {
	m := make(map[string]int)
	for k, v := range words {
		m[v] = k
	}
	length := len(words)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		offset := i % length
		_, exist := m[words[offset]]
		if !exist {
			b.Fatal("miss key")
		}
	}
}
