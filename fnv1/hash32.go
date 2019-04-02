package fnv1

import "unsafe"

const (
	// FNV-1
	offset32 = uint32(2166136261)
	prime32  = uint32(16777619)

	// Init32 is what 32 bits hash values should be initialized with.
	Init32 = offset32
)

// HashString32 returns the hash of s.
func HashString32(s string) uint32 {
	return AddString32(Init32, s)
}

// HashBytes32 returns the hash of u.
func HashBytes32(b []byte) uint32 {
	return AddBytes32(Init32, b)
}

// HashUint32 returns the hash of u.
func HashUint32(u uint32) uint32 {
	return AddUint32(Init32, u)
}

// AddString32 adds the hash of s to the precomputed hash value h.
func AddString32(h uint32, s string) uint32 {
	i := 0
	n := (len(s) / 8) * 8

	for i != n {
		h = (h * prime32) ^ uint32(s[i])
		h = (h * prime32) ^ uint32(s[i+1])
		h = (h * prime32) ^ uint32(s[i+2])
		h = (h * prime32) ^ uint32(s[i+3])
		h = (h * prime32) ^ uint32(s[i+4])
		h = (h * prime32) ^ uint32(s[i+5])
		h = (h * prime32) ^ uint32(s[i+6])
		h = (h * prime32) ^ uint32(s[i+7])

		i += 8
	}

	for _, c := range s[i:] {
		h = (h * prime32) ^ uint32(c)
	}

	return h
}

// AddBytes32 adds the hash of b to the precomputed hash value h.
func AddBytes32(h uint32, b []byte) uint32 {
	return AddString32(h, *(*string)(unsafe.Pointer(&b)))
}

// AddUint32 adds the hash value of the 8 bytes of u to h.
func AddUint32(h, u uint32) uint32 {
	h = (h * prime32) ^ ((u >> 24) & 0xFF)
	h = (h * prime32) ^ ((u >> 16) & 0xFF)
	h = (h * prime32) ^ ((u >> 8) & 0xFF)
	h = (h * prime32) ^ ((u >> 0) & 0xFF)
	return h
}
