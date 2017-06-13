package fnv1a

const (
	// FNV-1a
	offset64 = uint64(14695981039346656037)
	prime64  = uint64(1099511628211)

	// Init64 is the what Hash64 values should be initialized with.
	Init64 = offset64
)

// Hash64String returns the FNV-1a hash of s.
func Hash64String(s string) uint64 {
	return Add64String(Init64, s)
}

// Add64String adds the FNV-1a hash of s to the precomputed hash value h.
func Add64String(h uint64, s string) uint64 {
	/*
		This is an unrolled version of this algorithm:

		for _, c := range s {
			h = (h ^ uint64(c)) * prime64
		}

		It seems to be ~1.5x faster than the simple loop in BenchmarkHash64:

		- BenchmarkHash64/hash_function-4   30000000   56.1 ns/op   642.15 MB/s   0 B/op   0 allocs/op
		- BenchmarkHash64/hash_function-4   50000000   38.6 ns/op   932.35 MB/s   0 B/op   0 allocs/op

	*/

	for i, n := 0, len(s); i != n; {
		switch n - i {
		case 1:
			h = (h ^ uint64(s[i])) * prime64
			i += 1

		case 2:
			h = (h ^ uint64(s[i])) * prime64
			h = (h ^ uint64(s[i+1])) * prime64
			i += 2

		case 3:
			h = (h ^ uint64(s[i])) * prime64
			h = (h ^ uint64(s[i+1])) * prime64
			h = (h ^ uint64(s[i+2])) * prime64
			i += 3

		case 4:
			h = (h ^ uint64(s[i])) * prime64
			h = (h ^ uint64(s[i+1])) * prime64
			h = (h ^ uint64(s[i+2])) * prime64
			h = (h ^ uint64(s[i+3])) * prime64
			i += 4

		case 5:
			h = (h ^ uint64(s[i])) * prime64
			h = (h ^ uint64(s[i+1])) * prime64
			h = (h ^ uint64(s[i+2])) * prime64
			h = (h ^ uint64(s[i+3])) * prime64
			h = (h ^ uint64(s[i+4])) * prime64
			i += 5

		case 6:
			h = (h ^ uint64(s[i])) * prime64
			h = (h ^ uint64(s[i+1])) * prime64
			h = (h ^ uint64(s[i+2])) * prime64
			h = (h ^ uint64(s[i+3])) * prime64
			h = (h ^ uint64(s[i+4])) * prime64
			h = (h ^ uint64(s[i+5])) * prime64
			i += 6

		case 7:
			h = (h ^ uint64(s[i])) * prime64
			h = (h ^ uint64(s[i+1])) * prime64
			h = (h ^ uint64(s[i+2])) * prime64
			h = (h ^ uint64(s[i+3])) * prime64
			h = (h ^ uint64(s[i+4])) * prime64
			h = (h ^ uint64(s[i+5])) * prime64
			h = (h ^ uint64(s[i+6])) * prime64
			i += 7

		default:
			h = (h ^ uint64(s[i])) * prime64
			h = (h ^ uint64(s[i+1])) * prime64
			h = (h ^ uint64(s[i+2])) * prime64
			h = (h ^ uint64(s[i+3])) * prime64
			h = (h ^ uint64(s[i+4])) * prime64
			h = (h ^ uint64(s[i+5])) * prime64
			h = (h ^ uint64(s[i+6])) * prime64
			h = (h ^ uint64(s[i+7])) * prime64
			i += 8
		}
	}
	return h
}
