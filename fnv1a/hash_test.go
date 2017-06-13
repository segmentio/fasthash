package fnv1a

import (
	"hash/fnv"
	"testing"
)

func TestHash64(t *testing.T) {
	t.Run("ensures the hash function behaves like the standard fnv1-a", func(t *testing.T) {
		for _, s := range [...]string{"", "A", "Hello World!", "DAB45194-42CC-4106-AB9F-2447FA4D35C2"} {
			t.Run(s, func(t *testing.T) {
				h := fnv.New64a()
				h.Write([]byte(s))

				sum1 := h.Sum64()
				sum2 := HashString64(s)

				if sum1 != sum2 {
					t.Errorf("invalid hash, expected %d but got %d", sum1, sum2)
				}
			})
		}
	})
}

func BenchmarkHash64(b *testing.B) {
	tests := []struct {
		scenario string
		function func(string) uint64
	}{
		{
			scenario: "standard hash function",
			function: func(s string) uint64 {
				h := fnv.New64a()
				h.Write([]byte(s))
				return h.Sum64()
			},
		},
		{
			scenario: "hash function",
			function: HashString64,
		},
	}

	const uuid = "DAB45194-42CC-4106-AB9F-2447FA4D35C2"

	for _, test := range tests {
		b.Run(test.scenario, func(b *testing.B) {
			for i := 0; i != b.N; i++ {
				test.function(uuid)
			}
			b.SetBytes(int64(len(uuid)))
		})
	}
}
