# Comparison with other sort algorithms 

Although heapsort has the same time bounds as merge sort, it requires only Θ(1) auxiliary space instead of merge sort's Θ(n). On typical modern architectures, efficient quicksort implementations generally outperform mergesort for sorting RAM-based arrays.[citation needed] On the other hand, merge sort is a stable sort and is more efficient at handling slow-to-access sequential media. Merge sort is often the best choice for sorting a linked list: in this situation it is relatively easy to implement a merge sort in such a way that it requires only Θ(1) extra space, and the slow random-access performance of a linked list makes some other algorithms (such as quicksort) perform poorly, and others (such as heapsort) completely impossible.

[Wikipedia](https://en.wikipedia.org/wiki/Merge_sort)

# Max k Distance from Sorted

You are given an unsorted array of size n. However, you are also told that where each element is in the unsorted array right now is at most k-1 positions away from its final correctly sorted position. You are told that 1 < k < n.

We can sort in O(n log k)

```C++
#include <iostream>
#include <vector>
#include <queue>
using namespace std;
 
// Function to sort a K-Sorted Array
void sortKSortedArray(vector<int> &A, int k)
{
    // create an empty min heap using std::priority_queue
    // use std::greater as the comparison function for min-heap
    priority_queue<int, std::vector<int>, std::greater<int>> pq;
 
    int n = A.size();
 
    // insert first k+1 elements in the heap
    for (int i = 0; i <= k; i++) {
        pq.push(A[i]);
    }
 
    int index = 0;
 
    // do for remaining elements of the array
    for (int i = k + 1; i < n; i++)
    {
        // pop top element from min-heap and assign it to
        // next available array index
        A[index++] = pq.top();
        pq.pop();
 
        // push next array element into min-heap
        pq.push(A[i]);
    }
 
    // pop all remaining elements from the min heap and assign
    // it to next available array index
    while (!pq.empty())
    {
        A[index++] = pq.top();
        pq.pop();
    }
}
```

## Useful Resources

* [Sort a K-Sorted Array](https://www.techiedelight.com/sort-k-sorted-array/)
