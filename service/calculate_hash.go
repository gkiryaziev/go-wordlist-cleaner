package service

import (
	"hash/fnv"
)

// GetHashFvn return uint32 fnv hash
func GetHashFvn(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// GetHashFvn64 return uint64 fnv hash
func GetHashFvn64(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
