package fnv1a

import (
	"hash/fnv"
	"testing"

	"github.com/segmentio/fasthash"
)

func TestHash64(t *testing.T) {
	fasthash.TestHashString64(t, "fnv1a", fasthash.HashString64(fnv.New64a), HashString64)
	fasthash.TestHashUint64(t, "fnv1a", fasthash.HashUint64(fnv.New64a), HashUint64)
}

func BenchmarkHash64(b *testing.B) {
	fasthash.BenchmarkHashString64(b, "fnv1a", fasthash.HashString64(fnv.New64a), HashString64)
}
