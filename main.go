package main

import (
	"fmt"
	"hash"
	"time"

	"github.com/google/uuid"
	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	filter []uint8
	size   int32
}

var mHasher hash.Hash32

func init() {
	mHasher = murmur3.New32WithSeed(uint32(time.Now().Unix()))
}

func murmurhash(key string, size int32) int32 {
	mHasher.Write([]byte(key))
	result := mHasher.Sum32() % uint32(size)
	mHasher.Reset()
	return int32(result)
}

func NewBloomFilter(size int32) *BloomFilter {
	return &BloomFilter{
		filter: make([]uint8, size),
		size:   size,
	}
}

func (b *BloomFilter) Add(key string) {
	idx := murmurhash(key, b.size)
	aIdx := idx / 8
	bIdx := idx % 8
	b.filter[aIdx] = b.filter[aIdx] | (1 << bIdx)
}

func (b *BloomFilter) Exists(key string) bool {
	idx := murmurhash(key, b.size)
	aIdx := idx / 8
	bIdx := idx % 8
	return b.filter[aIdx]&(1<<bIdx) > 0
}

func main() {
	dataset := make([]string, 0)
	dataset_exists := make(map[string]bool)
	dataset_notexists := make(map[string]bool)

	for i := 0; i < 500; i++ {
		u := uuid.New()
		dataset = append(dataset, u.String())
		dataset_exists[u.String()] = true
	}

	for i := 0; i < 500; i++ {
		u := uuid.New()
		dataset = append(dataset, u.String())
		dataset_notexists[u.String()] = false
	}

	bloom := NewBloomFilter(1000)

	for key := range dataset_exists {
		bloom.Add(key)
	}

	falsePositives := 0
	for _, Key := range dataset {
		exists := bloom.Exists(Key)
		if exists {
			if _, ok := dataset_notexists[Key]; ok {
				falsePositives++
			}
		}
	}

	fmt.Println(100 * float64(falsePositives) / float64(len(dataset)))
}
