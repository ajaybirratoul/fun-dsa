# Quadtree
This is a C++ implementation of a basic Quadtree, a data structure used for spatial partitioning in 2D space. The Quadtree allows efficient storing and retrieval of points in a 2D space, making it useful for various applications like collision detection, range queries, and spatial indexing.

## How the Quadtree Works

The Quadtree is a recursive data structure that divides a 2D space into four quadrants. Each node in the Quadtree represents a rectangular region, and if the number of points in a node exceeds a certain threshold (CAPACITY), the node is subdivided into four children, each representing one of the quadrants. This process continues until each leaf node contains at most CAPACITY points.

## Usage

To use the Quadtree, follow these steps:

1. Create a new C++ project or add the `main.cpp` file containing the Quadtree implementation to your existing project.

2. Create an instance of the `Quadtree` class, specifying the width and height of the 2D space you want to partition.

3. Use the `insert` method to add points to the Quadtree.

Here's an example of how to use the Quadtree:

```cpp
int main() {
    Quadtree quadtree(800, 600);

    quadtree.insert(Point(100, 100));
    quadtree.insert(Point(200, 200));
    quadtree.insert(Point(300, 300));
    quadtree.insert(Point(400, 400));

    return 0;
}
```

## Features

- Spatial partitioning: The Quadtree efficiently partitions a 2D space into quadrants, making it easy to store and retrieve points in the space.

- Dynamic resizing: The Quadtree can dynamically resize as more points are added to maintain efficient spatial partitioning.
