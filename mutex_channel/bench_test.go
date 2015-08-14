package main

import (
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	m := NewMapMutex()

	for i := 0; i < 10; i++ {
		m.Write()
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Read()
	}
}

func BenchmarkChannel(b *testing.B) {
	m := NewMapChannel()

	for i := 0; i < 10; i++ {
		m.Write()
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Read()
	}
}
