# Skip List

This is a Skip List implementation in Rust. Skip List is a data structure that provides an ordered collection with efficient search, insert, and delete operations. It uses multiple linked lists with different skip distances to achieve a balance between efficiency and simplicity.

## Usage

To use the Skip List, follow these steps:

1. Include the Skip List implementation in your Rust project.

2. Create an instance of the `SkipList` struct using the `new` function.

3. Insert elements into the Skip List using the `insert` method.

4. Check if elements are present in the Skip List using the `contains` method.

Here's an example of using the Skip List:

```rust
let mut skip_list = SkipList::new();
skip_list.insert(5);
skip_list.insert(10);
skip_list.insert(2);

let contains_5 = skip_list.contains(5);   // true
let contains_7 = skip_list.contains(7);   // false
```

## Features

- Efficient search: The Skip List provides efficient search operations with an average time complexity of O(log n).

- Dynamic structure: The Skip List dynamically adjusts its height to maintain a balance between search efficiency and memory usage.

- Ordered collection: The Skip List keeps the elements in a sorted order, allowing efficient range queries and ordered iteration.
