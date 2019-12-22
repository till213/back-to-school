# Data Structures

## Common Dictionary Operations

|        | Ordered Array | Linear List | Balanced Tree |       Hash Table |
| ------ | :-----------: | ----------- | ------------- | ---------------: |
| find   |   O(log n)    | O(n)        | O(log n)      | O(1)&#x00B9; |
| next   |     O(1)      | O(1)        | O(log n)      |             O(n) |
| insert |     O(n)      | O(1)        | O(log n)      |             O(1)&#x00B9; |
| delete |     O(n)      | O(1)        | O(log n)      |             O(1)&#x00B2; |

&#x00B9; On the average, but not necessarily in the worst case

&#x00B2; Deletions are possible, but may degrade performace