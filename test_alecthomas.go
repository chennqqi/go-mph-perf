package main

import (
	"testing"

	"github.com/alecthomas/mph"
)

func intToByte(v int) []byte {
	u := uint32(v)
	r := make([]byte, 4)
	r[0] = byte((u) >> 24 & 0xFF)
	r[1] = byte((u >> 16) & 0xFF)
	r[2] = byte((u >> 8) & 0xFF)
	r[3] = byte(u & 0xFF)
	return r
}

func Benchmark_alecthomasNoBuild(b *testing.B) {
	t := mph.Builder()
	keys := make([][]byte, len(words))
	for k, v := range words {
		kbyte := intToByte(k)
		keys[k] = kbyte
		t.Add(kbyte, []byte(v))
	}
	h, err := t.Build()
	if err != nil {
		b.Fatal("error:", err)
		return
	}
	length := len(words)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		offset := i % length
		v := h.Get(keys[offset])
		if v == nil {
			b.Fatal("missing key")
		}
	}
}
