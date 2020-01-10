# Trees

## Segment Tree

Segment tree or segtree is a basically a binary tree used for storing the intervals or segments. Each node in the segment tree represents an interval.

Since a segtree is a binary tree, we can use a simple linear array to represent the segment tree. In almost any segtree problem we need think about what we need to store in the segment tree?.

For example, if we want to find the sum of all the elements in an array from index left to right, then at each node (except leaf nodes) we will store the sum of its children nodes. If we want to find the minimum of all the elements in an array from index left to right, then at each node (except leaf nodes) we will store the minimum of its children nodes.

![Segment Tree](https://he-s3.s3.amazonaws.com/media/uploads/a0c7f90.jpg)

![Array representation](https://he-s3.s3.amazonaws.com/media/uploads/aad673e.jpg)

### Implementation

The root of the segtree will represent the whole array A[0:N-1]. Then we will break the interval or segment into half and the two children of the root will represent the A[0:(N-1) / 2] and A[(N-1) / 2 + 1:(N-1)].

So in each step we will divide the interval into half and the two children will represent the two halves. So the height of the segment tree will be log N. There are N leaves representing the N elements of the array. The number of internal nodes is N-1. So total number of nodes are 2*N - 1.

## Useful Resources

* [Morris Tree Traversal (Space O(1)) (YouTube)](https://www.youtube.com/watch?v=wGXB9OWhPTg)

* [Segment Tree (HackerEarth)](https://www.hackerearth.com/practice/notes/segment-tree-and-lazy-propagation/)

* [Various Trees (Growing with the Web)](https://www.growingwiththeweb.com/p/explore.html?t=Data%20structure)
