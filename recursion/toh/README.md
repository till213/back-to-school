# Tower of Hanoi

## Task

Given a tower of n disks (set D), smallest on top, largest on bottom, held in place by a peg (source) peg S. Intermediate peg I and target peg T.

The task is to transfer all disks from S via I to T, without ever placing a larger disk on a smaller one.

## Solution

### Recursive

Divide et impera: be D' all disks of D except the lowest (largest) one. D' is possibly empty.

1. Transfer D' to intermediate 
2. Move largest disk to target peg T
3. Tranfer D' on top of the largest disk at T

How to solve step 1? Enter recursion: simply apply the same steps for the n-1 disks, until n = 0.

## Complexity

An iterative solution exists (assuming "cyclic pegs", taking "odd/even" towers into account) which solves the problem in:

&theta;(2<sup>n</sup> - 1) (exponential)