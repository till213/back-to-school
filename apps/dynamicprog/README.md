# Dynamic Programming

## Recipe

### Backtracking

1. Recursively try every possible next "move" from the current situation, until one succeeds.
2. If we get stuck, then we "backtrack" (removing the last step)
3. Without any further ado this quickly results in bad exponential running times, such as O(2<sup>n</sup>)
4. However space is usually O(n), that is the space used by the stack (recursion).

Example recursion tree for the [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/) problem:
![Recursion tree in the Climbing Stairs problem](https://leetcode.com/problems/climbing-stairs/Figures/70_Climbing_Stairs_rt.jpg)

### Dynamic Programming - Top Down

1. In the recursive approach we often need to calculate the same situation ("same state + inputs").
2. To prevent the repeated re-calculation of the same sub-problems we use "memoisation", a "dynamic table" used to store previous results (it may be sufficient to store only the last n results, e.g. Fibonacci sequence only requires the previous two results)
3. The steps are the same as in the backtracking approach, except that we store the latest result(s)
4. Time complexity can often be reduced to polynomial time, e.g. O(n&#x00B2;) for the [Jump Game](https://leetcode.com/problems/jump-game/)

### Dynamic Programming - Bottom Up

1. We try to come up with an iterative solution (no recursion)
2. The recursion is usually eliminated by trying to reverse the order of the steps from the top-down approach.
3. E.g. in the [Jump Game](https://leetcode.com/problems/jump-game/) we note that we can only ever move to the right
4. This means that if we start from the right of the array, every time we will query a position to our right, that position has already be determined as being GOOD or BAD.
5. There follows that we don't need to recurse anymore, as we will always hit the memo table.