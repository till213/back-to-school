# Backtracking

## General Pattern

```pascal
PROCEDURE Try;
BEGIN 
    intialize selection of candidates;
    REPEAT 
        select next; 
        IF acceptable THEN
            record it;
            IF solution incomplete THEN Try;
            IF not successful THEN cancel recording END END
        END
    UNTIL successful OR no more candidates
END Try
```

## With Depth Limit

```pascal
PROCEDURE Try(i: INTEGER); 
    V AR k: INTEGER;
BEGIN 
    k := 0;
    REPEAT 
        k := k+1; 
        select k-th candidate;
        IF acceptable THEN record it;
        IF i < n THEN 
            Try(i+1);
            IF not successful THEN cancel recording END
        END 
    END
UNTIL successful OR (k = m) END Try
```