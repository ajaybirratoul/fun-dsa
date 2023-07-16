#include <iostream>
#include <vector>
#include <bitset>
#include <functional>
#include <random>

class BloomFilter {
private:
    std::vector<bool> bitset;
    std::vector<std::function<size_t(const std::string&)>> hashFunctions;
    size_t numHashes;

public:
    BloomFilter(size_t numHashes) : numHashes(numHashes) {}

    void insert(const std::string& key) {
        if (bitset.empty()) {
            throw std::logic_error("Bloom filter not initialized");
        }

        for (const auto& hashFunc : hashFunctions) {
            size_t index = hashFunc(key) % bitset.size();
            bitset[index] = true;
        }
    }

    bool contains(const std::string& key) const {
        if (bitset.empty()) {
            throw std::logic_error("Bloom filter not initialized");
        }

        for (const auto& hashFunc : hashFunctions) {
            size_t index = hashFunc(key) % bitset.size();
            if (!bitset[index])
                return false;
        }
        return true;
    }

    void initialize(size_t numElements, double falsePositiveRate) {
        size_t bitsetSize = calculateOptimalSize(numElements, falsePositiveRate);
        bitset.resize(bitsetSize, false);

        hashFunctions.clear();
        hashFunctions.reserve(numHashes);
        for (size_t i = 0; i < numHashes; ++i) {
            hashFunctions.emplace_back(generateHashFunction(bitsetSize));
        }
    }

private:
    size_t calculateOptimalSize(size_t numElements, double falsePositiveRate) const {
        double ln2Squared = std::log(2) * std::log(2);
        return static_cast<size_t>(-((numElements * std::log(falsePositiveRate)) / ln2Squared));
    }

    std::function<size_t(const std::string&)> generateHashFunction(size_t bitsetSize) const {
        std::random_device rd;
        std::mt19937_64 gen(rd());
        std::uniform_int_distribution<size_t> distribution(0, bitsetSize - 1);
        size_t a = distribution(gen);
        size_t b = distribution(gen);

        return [a, b, bitsetSize](const std::string& str) {
            size_t hash1 = std::hash<std::string>{}(str);
            size_t hash2 = a * hash1 + b;
            return hash2 % bitsetSize;
        };
    }
};

int main() {
    BloomFilter bloomFilter(3);
    bloomFilter.initialize(100, 0.05);
    bloomFilter.insert("apple");
    bloomFilter.insert("banana");
    bloomFilter.insert("orange");

    std::cout << "Contains apple? " << std::boolalpha << bloomFilter.contains("apple") << std::endl;
    std::cout << "Contains banana? " << std::boolalpha << bloomFilter.contains("banana") << std::endl;
    std::cout << "Contains orange? " << std::boolalpha << bloomFilter.contains("orange") << std::endl;
    std::cout << "Contains grape? " << std::boolalpha << bloomFilter.contains("grape") << std::endl;

    return 0;
}
