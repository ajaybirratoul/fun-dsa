# Van Emde Boas Tree (VEB Tree)

This is a Go implementation of the Van Emde Boas Tree (VEB Tree), a specialized data structure that provides efficient search, insertion, and deletion operations for elements in a universe of size 2^k for some integer k.

## Usage

To use the VEB Tree, follow these steps:

1. Include the `veb_tree.go` file in your Go project.

2. Create an instance of the `VEBTree` struct using the `newVEBTree` function, specifying the size of the universe (2^k).

3. Perform various operations on the VEB Tree such as:
   - `insert`: to insert an element into the VEB Tree.
   - `delete`: to delete an element from the VEB Tree.
   - `member`: to check if an element is present in the VEB Tree.
   - `minimum`: to get the minimum element in the VEB Tree.
   - `maximum`: to get the maximum element in the VEB Tree.
   - `successor`: to get the successor of an element in the VEB Tree.
   - `predecessor`: to get the predecessor of an element in the VEB Tree.

Here's an example of using the VEB Tree:

```go
package main

import (
	"fmt"
)

func main() {
	v := newVEBTree(16)

	v.insert(5)
	v.insert(10)
	v.insert(2)

	fmt.Println("Minimum:", v.minimum())
	fmt.Println("Maximum:", v.maximum())

	fmt.Println("Successor of 5:", v.successor(5))
	fmt.Println("Predecessor of 5:", v.predecessor(5))

	v.delete(10)

	fmt.Println("Minimum:", v.minimum())
	fmt.Println("Maximum:", v.maximum())
}
```

## Features

- Efficient search, insertion, and deletion: The VEB Tree provides efficient operations for elements in a universe of size 2^k.

- Minimum and maximum retrieval: The VEB Tree allows retrieving the minimum and maximum elements quickly.

- Successor and predecessor computation: The VEB Tree supports calculating the successor and predecessor of elements efficiently.
