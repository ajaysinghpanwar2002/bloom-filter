# Bloom Filter Implementation in Go

## Overview
This repository contains a Go implementation of a Bloom filter, an efficient data structure for probabilistic set membership testing. It's designed to quickly check if an element is possibly in a set, with a certain probability of false positives but no false negatives.

## Bloom Filter
A Bloom filter is a space-efficient probabilistic data structure that is used to test whether an element is a member of a set. False positive matches are possible, but false negatives are not. In other words, a query returns either "possibly in set" or "definitely not in set". Bloom filters are used to reduce the need for costly disk or network operations by eliminating non-existent keys.

## Algorithms Used
- **MurmurHash3**: A non-cryptographic hash function suitable for general hash-based lookup. It's used here to generate hash values for the elements.
- **Bit Array**: The core of the Bloom filter's data structure, where elements are marked by setting bits at specific indices determined by the hash functions.

## Approach
The implementation creates a Bloom filter with a specified size and uses MurmurHash3 to generate indices for elements to be added. Each element's presence is encoded in a bit array, allowing for efficient queries with minimal space usage. The main components of this implementation include:
- Initialization of a MurmurHash3 hasher with a random seed.
- A `BloomFilter` struct that holds the bit array and its size.
- Functions to add elements to the filter (`Add`) and to check for their existence (`Exists`).

## Outcomes
The provided code demonstrates the use of the Bloom filter with a dataset of UUIDs, half of which are added to the filter and the other half used to test for false positives. The output is the percentage of false positives observed, which is an important metric for evaluating the effectiveness and configuration of a Bloom filter.

## Usage
To use this Bloom filter implementation in your project, include the `.go` file in your project directory. You can initialize a new Bloom filter with a specified size and use the `Add` and `Exists` functions to work with elements.

```go
bloom := NewBloomFilter(1000) // Initialize Bloom filter with 1000 bits
bloom.Add("your_element")     // Add an element
exists := bloom.Exists("your_element") // Check if an element exists
fmt.Println(exists) // true (possibly in set) or false (definitely not in set)
