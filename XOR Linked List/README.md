# XOR Linked List

This is a Kotlin implementation of the XOR Linked List, a memory-efficient data structure that uses XOR operations on addresses to represent linked nodes. The XOR Linked List allows efficient traversal and saves space by storing the XOR of the addresses of the previous and next nodes in each node.

## Implementation Details

The XOR Linked List consists of two main classes:

1. `Node`: Represents a single node in the XOR Linked List. Each node contains an integer value and an XOR of the addresses of its previous and next nodes.

2. `XORLinkedList`: Manages the list and provides methods to add elements and retrieve elements based on their indices.

## Usage

To use the XOR Linked List, follow these steps:

1. Include the `Node` and `XORLinkedList` classes in your Kotlin project.

2. Create an instance of the `XORLinkedList` class.

3. Use the `add` method to add elements to the list.

4. Use the `get` method to retrieve elements based on their indices.

Here's an example of using the XOR Linked List:

```kotlin
fun main() {
    val xorLinkedList = XORLinkedList()
    xorLinkedList.add(10)
    xorLinkedList.add(20)
    xorLinkedList.add(30)

    println(xorLinkedList.get(0)) // Output: 10
    println(xorLinkedList.get(1)) // Output: 20
    println(xorLinkedList.get(2)) // Output: 30
    println(xorLinkedList.get(3)) // Output: null
}
```

## Features

- Memory-efficient: The XOR Linked List uses XOR operations on addresses to save memory by not storing explicit pointers.

- Efficient traversal: The XOR Linked List allows traversing the list efficiently using the XOR of node addresses.

- Add and get elements: The `add` method adds elements to the list, and the `get` method retrieves elements based on their indices.

## Limitations

The current implementation uses Kotlin's hash code mechanism to simulate XOR operations with addresses. While this works for demonstration purposes, it's not a perfect XOR implementation and may not work correctly in all scenarios. For real-world applications, low-level bit manipulation should be used to handle XOR operations with addresses accurately.
