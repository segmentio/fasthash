package fnv1a

import (
	"hash/fnv"
	"testing"

	"github.com/segmentio/fasthash"
	"github.com/segmentio/fasthash/fasthashtest"
)

func TestHash64(t *testing.T) {
	fasthashtest.TestHashString64(t, "fnv1a", fasthash.HashString64(fnv.New64a), HashString64)
	fasthashtest.TestHashBytes64(t, "fnv1a", fasthash.HashBytes64(fnv.New64a), HashBytes64)
	fasthashtest.TestHashUint64(t, "fnv1a", fasthash.HashUint64(fnv.New64a), HashUint64)
}

func BenchmarkHash64(b *testing.B) {
	fasthashtest.BenchmarkHashString64(b, "fnv1a", fasthash.HashString64(fnv.New64a), HashString64)
}
