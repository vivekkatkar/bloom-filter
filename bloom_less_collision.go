package main

import (
	"hash/fnv"

	"github.com/spaolacci/murmur3"
)

type LessCollisionBloomFilter struct {
	filter []uint8
	size   int
}

func NewLessCollisionBloomFilter(size int) *LessCollisionBloomFilter {
	return &LessCollisionBloomFilter{
		filter: make([]uint8, size),
		size:   size,
	}
}

func hash1(key string, size int) int {
	h := murmur3.New64()
	h.Write([]byte(key))
	return int(h.Sum64() % uint64(size))
}

func hash2(key string, size int) int {
	h := fnv.New64()
	h.Write([]byte(key))
	return int(h.Sum64() % uint64(size))
}

func hash3(key string, size int) int {
	h := murmur3.New32()
	h.Write([]byte(key))
	return int(uint64(h.Sum32()) % uint64(size))
}

func (b *LessCollisionBloomFilter) Add(key string) {
	hashes := []int{
		hash1(key, b.size),
		hash2(key, b.size),
		hash3(key, b.size),
	}

	for _, hv := range hashes {
		arrayIdx := hv / 8
		bitIdx := hv % 8
		b.filter[arrayIdx] |= (1 << bitIdx)
	}
}

func (b *LessCollisionBloomFilter) Exists(key string) bool {
	hashes := []int{
		hash1(key, b.size),
		hash2(key, b.size),
		hash3(key, b.size),
	}

	for _, hv := range hashes {
		arrayIdx := hv / 8
		bitIdx := hv % 8
		if (b.filter[arrayIdx] & (1 << bitIdx)) == 0 {
			return false
		}
	}
	return true
}
