package fasthash

import (
	"encoding/binary"
	"hash"
	"testing"
)

// HashString64 makes a hashing function from the hashing algorithm returned by f.
func HashString64(f func() hash.Hash64) func(string) uint64 {
	return func(s string) uint64 {
		h := f()
		h.Write([]byte(s))
		return h.Sum64()
	}
}

// HashUint64 makes a hashing function from the hashing algorithm return by f.
func HashUint64(f func() hash.Hash64) func(uint64) uint64 {
	return func(u uint64) uint64 {
		b := [8]byte{}
		binary.BigEndian.PutUint64(b[:], u)
		h := f()
		h.Write(b[:])
		return h.Sum64()
	}
}

// TestHashString64 is the implementation of a test suite to verify the
// behavior of a hashing algorithm.
func TestHashString64(t *testing.T, name string, reference func(string) uint64, algorithm func(string) uint64) {
	t.Run(name, func(t *testing.T) {
		for _, s := range [...]string{"", "A", "Hello World!", "DAB45194-42CC-4106-AB9F-2447FA4D35C2"} {
			t.Run(s, func(t *testing.T) {
				sum1 := reference(s)
				sum2 := algorithm(s)

				if sum1 != sum2 {
					t.Errorf("invalid hash, expected %d but got %d", sum1, sum2)
				}
			})
		}
	})
}

// TestHashUint64 is the implementation of a test suite to verify the
// behavior of a hashing algorithm.
func TestHashUint64(t *testing.T, name string, reference func(uint64) uint64, algorithm func(uint64) uint64) {
	t.Run(name, func(t *testing.T) {
		sum1 := reference(42)
		sum2 := algorithm(42)

		if sum1 != sum2 {
			t.Errorf("invalid hash, expected %d but got %d", sum1, sum2)
		}
	})
}

// BenchmarkHashString64 is the implementation of a benchmark suite to compare
// the CPU and memory efficiency of a hashing algorithm against a reference
// implementation.
func BenchmarkHashString64(b *testing.B, name string, reference func(string) uint64, algorithm func(string) uint64) {
	b.Run(name, func(b *testing.B) {
		b.Run("reference", func(b *testing.B) { benchmark(b, reference) })
		b.Run("algorithm", func(b *testing.B) { benchmark(b, algorithm) })
	})
}

func benchmark(b *testing.B, hash func(string) uint64) {
	const uuid = "DAB45194-42CC-4106-AB9F-2447FA4D35C2"

	for i := 0; i != b.N; i++ {
		hash(uuid)
	}

	b.SetBytes(int64(len(uuid)))
}
