package fasthashtest

import (
	"fmt"
	"strings"
	"testing"
)

// TestHashString64 is the implementation of a test suite to verify the
// behavior of a hashing algorithm.
func TestHashString64(t *testing.T, name string, reference func(string) uint64, algorithm func(string) uint64) {
	t.Run(name, func(t *testing.T) {
		for _, s := range [...]string{"", "A", "Hello World!", "DAB45194-42CC-4106-AB9F-2447FA4D35C2", "你好吗"} {
			t.Run(s, func(t *testing.T) {
				if reference == nil {
					algorithm(s)
				} else {
					sum1 := reference(s)
					sum2 := algorithm(s)

					if sum1 != sum2 {
						t.Errorf("invalid hash, expected %x but got %x", sum1, sum2)
					}
				}
			})
		}
	})
}

// TestHashBytes64 is the implementation of a test suite to verify the
// behavior of a hashing algorithm.
func TestHashBytes64(t *testing.T, name string, reference func([]byte) uint64, algorithm func([]byte) uint64) {
	t.Run(name, func(t *testing.T) {
		for _, s := range [...]string{"", "A", "Hello World!", "DAB45194-42CC-4106-AB9F-2447FA4D35C2"} {
			t.Run(s, func(t *testing.T) {
				b := []byte(s)
				if reference == nil {
					algorithm(b)
				} else {
					sum1 := reference(b)
					sum2 := algorithm(b)

					if sum1 != sum2 {
						t.Errorf("invalid hash, expected %x but got %x", sum1, sum2)
					}
				}
			})
		}
	})
}

// TestHashUint64 is the implementation of a test suite to verify the
// behavior of a hashing algorithm.
func TestHashUint64(t *testing.T, name string, reference func(uint64) uint64, algorithm func(uint64) uint64) {
	t.Run(name, func(t *testing.T) {
		if reference == nil {
			algorithm(42)
		} else {
			sum1 := reference(42)
			sum2 := algorithm(42)

			if sum1 != sum2 {
				t.Errorf("invalid hash, expected %x but got %x", sum1, sum2)
			}
		}
	})
}

// BenchmarkHashString64 is the implementation of a benchmark suite to compare
// the CPU and memory efficiency of a hashing algorithm against a reference
// implementation.
func BenchmarkHashString64(b *testing.B, name string, reference func(string) uint64, algorithm func(string) uint64) {
	b.Run(name, func(b *testing.B) {
		if reference != nil {
			b.Run("reference", func(b *testing.B) { benchmark(b, reference) })
		}
		b.Run("optimized", func(b *testing.B) { benchmark(b, algorithm) })
	})
}

var benchmarkStrings = [...]string{
	"asdf",
	"hello world",
	"DAB45194-42CC-4106-AB9F-2447FA4D35C2",
	strings.Repeat("1234567890", 100),
}

func benchmark(b *testing.B, hash func(string) uint64) {
	for _, s := range benchmarkStrings {
		b.Run(fmt.Sprintf("strlen=%d", len(s)), func(b *testing.B) {
			for i := 0; i != b.N; i++ {
				hash(s)
			}
			b.SetBytes(int64(len(s)))
		})
	}
}
