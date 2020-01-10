# Continuous Median

## Definition

The median of a set containing odd number of integer is the middle element after sorting. For a set of even numbers it is the average of the two middle numbers

## Calculation

For a random stream of numbers we can continuously keep track of the median as follows.

Given the following example random number stream:
```
3, 6, 7, 5, 1, 2, 4
```
Sorting this stream results in:
```
1, 2, 3, 4, 5, 6, 7
```
and the median is the middle element 4.

## Continuous Calculation

 But what about if we keep adding new numbers? The trick is to use two heaps: a *min-heap* which keeps track of the "upper half" of the sorted numbers and a *max-heap* which keeps track of the "lower half", with the invariant:
 ```
 INVARIANT:
 0 <= size(max-heap) - size(min-heap) <= 1
 ```
(the *max-heap* is at most one element larger than the *min-heap*)

So with the above sorted stream:
```
1, 2, 3, 4, 5, 6, 7
**********  #######
Max-heap(4) Min-heap(5)
```
The number in brackets indicates the minimum respectively the maximum value of the corresponding heap. In this case the maximum element of the max-heap is our median (4).

If we add a new number, say 8, then we get:

```
1, 2, 3, 4, 5, 6, 7, 8
**********  ##########
Max-heap(4) Min-heap(5)
```

Of course we do not sort the actual number stream, but use the implicit sorting of the heap data structure. We added 8 to the *min-heap* because:

* The *min-heap* was smaller than the *max-heap*
* The number 8 was larger than the minimum element

The median is now the average of the two middle elements 4 and 5: (4+5)/2 = 4.5

Let's add another number, say 9. Both heaps have the same lenght now, so in order to keep the invariant valid we need to first *push over* one element from the *min-heap* to the *max-heap*, specifically the minimum element which is 5:

```
1, 2, 3, 4, 5, 6, 7, 8
*************  #######
Max-heap(5)    Min-heap(5)
```

Of course this (temporarily) invalidates our invariant, but which is quickly amended by actually adding the new number 9:

```
1, 2, 3, 4, 5, 6, 7, 8, 9
*************  ##########
Max-heap(5)    Min-heap(5)
```

The new median is now simply the maximum element of the *max-heap* again, which is number 5.

What if we add a smaller number, say 0? It belongs into the *max-heap*, as it is smaller that the minimum value 6 of the *min-heap*. 

However the *max-heap* is already one element larger than the *min-heap*, so we need again to "push over" one element, this time from the *max-heap* to the *min-heap*. We push over the maximum value 5:

```
1, 2, 3, 4, 5, 6, 7, 8, 9
**********  #############
Max-heap(4) Min-heap(5)
```

Again the invariant is (temporarily) invalidated, but we amend this by adding the new number 0:

```
0, 1, 2, 3, 4, 5, 6, 7, 8, 9
*************  #############
Max-heap(4)    Min-heap(5)
```

As the heaps are of equal size again the new median is now (4+5)/2 = 4.5.
