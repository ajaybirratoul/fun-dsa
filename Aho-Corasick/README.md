# Aho-Corasick Algorithm

This is a Rust implementation of the Aho-Corasick algorithm, a string-searching algorithm that efficiently finds multiple patterns in a given text. The Aho-Corasick algorithm constructs a trie-like data structure to search for multiple patterns simultaneously.

## Usage

To use the Aho-Corasick algorithm, follow these steps:

1. Create a new Rust project or add the `main.rs` file containing the implementation to your existing project.

2. Implement the `AhoCorasick`, `TrieNode`, and related methods as shown in the `main.rs` file.

3. Create an instance of the `AhoCorasick` struct and insert the keywords (patterns) you want to search for using the `insert` method.

4. Build the failure links in the trie using the `build_failure_links` method.

5. Use the `search` method to find occurrences of the keywords in the given text.

Here's an example of how to use the Aho-Corasick algorithm:

```rust
fn main() {
    let mut ac = AhoCorasick::new();

    // Add keywords to search for
    ac.insert("he");
    ac.insert("she");
    ac.insert("his");
    ac.insert("hers");

    // Build failure links
    ac.build_failure_links();

    let text = "ushers";

    // Search for keywords in the text
    let found_words = ac.search(text);
    println!("Found words: {:?}", found_words);
}
```

## Features

- Efficient string searching: The Aho-Corasick algorithm efficiently searches for multiple patterns in a given text.

- Trie data structure: The algorithm constructs a trie-like data structure to store the patterns for efficient searching.
