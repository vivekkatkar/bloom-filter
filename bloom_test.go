package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/uuid"
)

func generate_string() string {
	u := uuid.New()
	str := u.String()
	return str
}

const N = 1000
const max_size = 10000
const steps = 100

// how false positive rate decrease with increaing bloom filter size
func TestBasicBloomFilter(t *testing.T) {

	println("Testing False positive rate in basic bloom filter")

	dataset := []string{}
	dataset_map := make(map[string]bool)

	for i := 0; i < N; i++ {
		key := generate_string()
		dataset = append(dataset, key)
		dataset_map[key] = true
	}

	for sz := 100; sz <= max_size; sz += steps {
		bloom := NewBasicBloomFilter(sz)

		// Insert half of the keys
		for i := 0; i < N/2; i++ {
			bloom.add(dataset[i])
		}

		var false_positive = 0
		for i := N / 2; i < N; i++ {
			if bloom.exists(dataset[i]) {
				false_positive++
			}
		}

		false_positive_rate := float64(false_positive) / float64(N/2) * 100
		rounded := math.Round(false_positive_rate*100) / 100

		fmt.Printf("Size: %d  False positive rate: %.4f\n", sz, rounded)
	}
}

func TestOptimizedBloomFilter(t *testing.T) {

	println("Testing False positive rate in optimized bloom filter")

	dataset := []string{}
	dataset_map := make(map[string]bool)

	for i := 0; i < N; i++ {
		key := generate_string()
		dataset = append(dataset, key)
		dataset_map[key] = true
	}

	for sz := 100; sz <= max_size; sz += steps {
		bloom := NewOptimizedBloomFilter(sz)

		// Insert half of the keys
		for i := 0; i < N/2; i++ {
			bloom.add(dataset[i])
		}

		var false_positive = 0
		for i := N / 2; i < N; i++ {
			if bloom.exists(dataset[i]) {
				false_positive++
			}
		}

		false_positive_rate := float64(false_positive) / float64(N/2) * 100
		rounded := math.Round(false_positive_rate*100) / 100

		fmt.Printf("Size: %d  False positive rate: %.4f\n", sz, rounded)
	}
}

func TestLessCollisionBloomFilter(t *testing.T) {

	println("Testing False positive rate in less collision bloom filter")

	dataset := []string{}
	dataset_map := make(map[string]bool)

	for i := 0; i < N; i++ {
		key := generate_string()
		dataset = append(dataset, key)
		dataset_map[key] = true
	}

	for sz := 100; sz <= max_size; sz += steps {
		bloom := NewLessCollisionBloomFilter(sz)

		// Insert half of the keys
		for i := 0; i < N/2; i++ {
			bloom.Add(dataset[i]) // note: Add() instead of add()
		}

		var false_positive = 0
		for i := N / 2; i < N; i++ {
			if bloom.Exists(dataset[i]) { // note: Exists() instead of exists()
				false_positive++
			}
		}

		false_positive_rate := float64(false_positive) / float64(N/2) * 100
		rounded := math.Round(false_positive_rate*100) / 100

		fmt.Printf("Size: %d  False positive rate: %.4f\n", sz, rounded)
	}
}
