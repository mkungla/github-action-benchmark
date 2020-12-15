package fib

import (
	"testing"
)

func BenchmarkFib(b *testing.B) {
	b.Run("Fib10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var _ = Fib(10)
		}
	})
	b.Run("Fib20", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var _ = Fib(20)
		}
	})

}
