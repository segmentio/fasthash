package jody

import (
	"reflect"
	"unsafe"
)

const (
	shift    = 11
	constant = 0x1f3d5b79

	// Init64 is what 64 bits hash values should be initialized with.
	Init64 = uint64(0)
)

var mask64 = [...]uint64{
	0x0000000000000000,
	0x00000000000000ff,
	0x000000000000ffff,
	0x0000000000ffffff,
	0x00000000ffffffff,
	0x000000ffffffffff,
	0x0000ffffffffffff,
	0x00ffffffffffffff,
	0xffffffffffffffff,
}

func accumulate(data, acc, cons uint64) uint64 {
	acc = acc + data + cons
	acc = (acc<<shift | acc>>(64-shift)) ^ data
	acc = ((acc<<shift | acc>>(64-shift)) ^ cons) + data
	return acc
}

// HashString64 returns the hash of s.
func HashString64(s string) uint64 {
	return AddString64(Init64, s)
}

// HashBytes64 returns the hash of u.
func HashBytes64(b []byte) uint64 {
	return AddBytes64(Init64, b)
}

// HashUint64 returns the hash of u.
func HashUint64(u uint64) uint64 {
	return AddUint64(Init64, u)
}

// AddString64 adds the hash of s to the precomputed hash value h.
func AddString64(h uint64, s string) uint64 {
	/*
		This is an implementation of the jody hashing algorithm as found here:

		- https://github.com/jbruchon/jodyhash

		It was revisited a bit and only the 64 bit version is available for now,
		but it's indeed really fast:

		- BenchmarkHash64/jody/algorithm-4   100000000   11.8 ns/op   3050.83 MB/s   0 B/op   0 allocs/op

		Here's what the reference implementation looks like:

		hash_t hash = start_hash;
		hash_t element;
		hash_t partial_salt;
		size_t len;

		len = count / sizeof(hash_t);
		for (; len > 0; len--) {
			element = *data;
			hash += element;
			hash += constant;
			hash = (hash << shift) | hash >> (sizeof(hash_t) * 8 - shift);
			hash ^= element;
			hash = (hash << shift) | hash >> (sizeof(hash_t) * 8 - shift);
			hash ^= constant;
			hash += element;
			data++;
		}

		len = count & (sizeof(hash_t) - 1);
		if (len) {
			partial_salt = constant & tail_mask[len];
			element = *data & tail_mask[len];
			hash += element;
			hash += partial_salt;
			hash = (hash << shift) | hash >> (sizeof(hash_t) * 8 - shift);
			hash ^= element;
			hash = (hash << shift) | hash >> (sizeof(hash_t) * 8 - shift);
			hash ^= partial_salt;
			hash += element;
		}

		return hash;
	*/

	p := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	dataLen := len(s)

	for dataLen > 8 {
		v := *(*uint64)(p)
		h = accumulate(v, h, constant)
		p = unsafe.Pointer(uintptr(p) + 8)
		dataLen -= 8
	}

	if dataLen == 8 {
		v := *(*uint64)(p)
		return accumulate(v, h, constant)
	}

	if dataLen == 0 {
		return h
	}

	// now we have a remainder 1 <= r.Len <= 7

	off := uintptr(0)
	v := uint64(0)
	if 0 != dataLen&4 {
		v += uint64(*(*uint32)(unsafe.Pointer(uintptr(p) + off)))
		off += 4
	}

	if 0 != dataLen&2 {
		v += uint64(*(*uint16)(unsafe.Pointer(uintptr(p) + off)))
		off += 2
	}

	if 0 != dataLen&1 {
		v += uint64(*(*uint8)(unsafe.Pointer(uintptr(p) + off)))
	}

	return accumulate(v, h, constant&mask64[dataLen])
}

// AddBytes64 adds the hash of b to the precomputed hash value h.
func AddBytes64(h uint64, b []byte) uint64 {
	return AddString64(h, *(*string)(unsafe.Pointer(&b)))
}

// AddUint64 adds the hash value of the 8 bytes of u to h.
func AddUint64(h uint64, u uint64) uint64 {
	return accumulate(u, h, constant)
}
