package main 
import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

var hasher_obj = murmur3.New64()

func get_hash(key string, size int) int {
	hasher_obj.Write([]byte(key))
	hash := hasher_obj.Sum64()
	hasher_obj.Reset()

	return int(hash % uint64(size))
}

type BloomFilterBasic struct {
	filter []bool
	size   int
}

func NewBasicBloomFilter(size int) *BloomFilterBasic {
	return &BloomFilterBasic{
		filter: make([]bool, size),
		size:   size,
	}
}

func (b *BloomFilterBasic) add(key string) {
	var idx = get_hash(key, b.size)
	b.filter[idx] = true

	// fmt.Println("Wrote "+key+" at : ", idx)
}

func (b *BloomFilterBasic) exists(key string) bool {
	var idx = get_hash(key, b.size)
	return b.filter[idx]
}

func (b *BloomFilterBasic) print() {
	print("[")
	for _, val := range b.filter {
		print(val, " ")
	}
	print("]")
}

func basic_bloom_test() {
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
