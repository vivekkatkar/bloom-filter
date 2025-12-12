package main

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

var hasher = murmur3.New64()

func hash(key string, size int) int {
	hasher.Write([]byte(key))
	hash := hasher.Sum64()
	hasher.Reset()

	return int(hash % uint64(size))
}

type BloomFilter struct {
	filter []uint8
	size   int
}

func NewOptimizedBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		filter: make([]uint8, size),
		size:   size,
	}
}

func (b *BloomFilter) add(key string) {
	hashValue := hash(key, b.size)
	array_idx := hashValue / 8
	bit_idx := hashValue % 8

	b.filter[array_idx] |= (1 << bit_idx)
}

func (b *BloomFilter) exists(key string) bool {
	hashValue := hash(key, b.size)
	array_idx := hashValue / 8
	bit_idx := hashValue % 8

	return (b.filter[array_idx] & (1 << bit_idx)) != 0
}

func (b *BloomFilter) print() {
	print("[")
	for _, val := range b.filter {
		print(val, " ")
	}
	print("]")
}

func optimized_bloom_test() {
	bloom := NewBasicBloomFilter(16)
	keys := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	for _, key := range keys {
		bloom.add(key)
	}

	for _, key := range keys {
		fmt.Println(key+" exists: ", bloom.exists(key))
	}

	bloom.print()
}
