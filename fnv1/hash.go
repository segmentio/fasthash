package fnv1

const (
	// FNV-1
	offset64 = uint64(14695981039346656037)
	prime64  = uint64(1099511628211)

	// Init64 is what 64 bits hash values should be initialized with.
	Init64 = offset64
)

// HashString64 returns the hash of s.
func HashString64(s string) uint64 {
	return AddString64(Init64, s)
}

// HashUint64 returns the hash of u.
func HashUint64(u uint64) uint64 {
	return AddUint64(Init64, u)
}

// AddString64 adds the hash of s to the precomputed hash value h.
func AddString64(h uint64, s string) uint64 {
	for len(s) >= 8 {
		h = (h * prime64) ^ uint64(s[0])
		h = (h * prime64) ^ uint64(s[1])
		h = (h * prime64) ^ uint64(s[2])
		h = (h * prime64) ^ uint64(s[3])
		h = (h * prime64) ^ uint64(s[4])
		h = (h * prime64) ^ uint64(s[5])
		h = (h * prime64) ^ uint64(s[6])
		h = (h * prime64) ^ uint64(s[7])
		s = s[8:]
	}

	if len(s) >= 4 {
		h = (h * prime64) ^ uint64(s[0])
		h = (h * prime64) ^ uint64(s[1])
		h = (h * prime64) ^ uint64(s[2])
		h = (h * prime64) ^ uint64(s[3])
		s = s[4:]
	}

	if len(s) >= 2 {
		h = (h * prime64) ^ uint64(s[0])
		h = (h * prime64) ^ uint64(s[1])
		s = s[2:]
	}

	if len(s) > 0 {
		h = (h * prime64) ^ uint64(s[0])
	}

	return h
}

// AddUint64 adds the hash value of the 8 bytes of u to h.
func AddUint64(h uint64, u uint64) uint64 {
	h = (h * prime64) ^ ((u >> 56) & 0xFF)
	h = (h * prime64) ^ ((u >> 48) & 0xFF)
	h = (h * prime64) ^ ((u >> 40) & 0xFF)
	h = (h * prime64) ^ ((u >> 32) & 0xFF)
	h = (h * prime64) ^ ((u >> 24) & 0xFF)
	h = (h * prime64) ^ ((u >> 16) & 0xFF)
	h = (h * prime64) ^ ((u >> 8) & 0xFF)
	h = (h * prime64) ^ ((u >> 0) & 0xFF)
	return h
}
