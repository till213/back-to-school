# Value Swap

## Task

Implement a value swap, that is x :=: y.

## Solution

x := x xor y; y := x xor y; x := x xor y;

### Complexity

Three XOR operations.