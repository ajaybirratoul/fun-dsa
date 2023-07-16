# Bloom Filter

This is a Bloom filter implementation in C++. A Bloom filter is a space-efficient probabilistic data structure used to test whether an element is a member of a set. It provides an approximate containment check with a low probability of false positives. The implementation supports variable-size Bloom filters and utilizes double hashing.

## Usage

To use the Bloom filter, follow these steps:

1. Include the `BloomFilter` class in your C++ project.

2. Create an instance of the `BloomFilter` class, specifying the number of hash functions to use.

3. Initialize the Bloom filter by calling the `initialize` function, passing the expected number of elements and the desired false positive rate. This will calculate the optimal size of the Bloom filter and set it up for use.

4. Insert elements into the Bloom filter using the `insert` function.

5. Check if elements are potentially in the Bloom filter using the `contains` function. Note that the `contains` function may return false positives but will not return false negatives.

Here's an example of using the Bloom filter:

```cpp
BloomFilter bloomFilter(3);
bloomFilter.initialize(100, 0.05);
bloomFilter.insert("apple");
bloomFilter.insert("banana");
bloomFilter.insert("orange");

bool containsApple = bloomFilter.contains("apple");   // true
bool containsBanana = bloomFilter.contains("banana"); // true
bool containsGrape = bloomFilter.contains("grape");   // false
```

## Features

- Variable-size Bloom filter: The `initialize` function calculates the optimal size of the bitset based on the expected number of elements and the desired false positive rate.

- Double hashing: The implementation uses double hashing to generate unique hash values for the elements.
