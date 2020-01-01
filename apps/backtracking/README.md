# Backtracking

## General Pattern (Single Solution)

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

## General Pattern (All Solutions)

```pascal
PROCEDURE Try(i: INTEGER); VAR k: INTEGER;        
BEGIN
    FOR k := 0 TO n-1 DO
        select k th candidate;
        IF acceptable THEN record it;
        IF i < n THEN 
            Try(i+1)
        ELSE 
            output solution
        END;
        cancel recording END
    END 
END Try
```