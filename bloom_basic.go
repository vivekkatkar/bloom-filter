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
	filter []bool
	size   int
}

func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		filter: make([]bool, size),
		size:   size,
	}
}

func (b *BloomFilter) add(key string) {
	var idx = hash(key, b.size)
	b.filter[idx] = true

	fmt.Println("Wrote "+key+" at : ", idx)
}

func (b *BloomFilter) exists(key string) bool {
	var idx = hash(key, b.size)
	return b.filter[idx]
}

func (b *BloomFilter) print() {
	print("[")
	for _, val := range b.filter {
		print(val, " ")
	}
	print("]")
}

func main() {
	bloom := NewBloomFilter(16)
	keys := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	for _, key := range keys {
		bloom.add(key)
	}

	for _, key := range keys {
		fmt.Println(key+" exists: ", bloom.exists(key))
	}

	bloom.print()
}
