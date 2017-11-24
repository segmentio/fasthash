package fnv1

import (
	"hash/fnv"
	"testing"

	"github.com/segmentio/fasthash"
)

func TestHash32(t *testing.T) {
	fasthash.TestHashString32(t, "fnv1", fasthash.HashString32(fnv.New32), HashString32)
	fasthash.TestHashUint32(t, "fnv1", fasthash.HashUint32(fnv.New32), HashUint32)
}

func BenchmarkHash32(b *testing.B) {
	fasthash.BenchmarkHashString32(b, "fnv1", fasthash.HashString32(fnv.New32a), HashString32)
}
