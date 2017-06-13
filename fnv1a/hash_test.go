package fnv1a

import (
	"encoding/binary"
	"hash/fnv"
	"testing"

	"github.com/segmentio/fasthash"
)

func TestHash64(t *testing.T) {
	fasthash.TestHashString64(t, "fnv1a", fasthash.HashString64(fnv.New64a), HashString64)

	t.Run("ensures that hashing a 8-byte integration yields the same hash value than the equivalent string", func(t *testing.T) {
		u := uint64(1234567890)
		b := [8]byte{}

		binary.BigEndian.PutUint64(b[:], u)

		h1 := AddUint64(Init64, u)
		h2 := AddString64(Init64, string(b[:]))

		if h1 != h2 {
			t.Errorf("bad hash value: expected %X but found %X", h1, h2)
		}
	})
}

func BenchmarkHash64(b *testing.B) {
	fasthash.BenchmarkHashString64(b, "fnv1a", fasthash.HashString64(fnv.New64a), HashString64)
}
