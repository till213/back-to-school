# Population Count

## Task

Count the number of 1s in a bit array of size n.

## Solution

### Inspect Every Bit

Simply count the bits which are set, in &theta;(n)

### Skip the Zeros

c := 0
while w != 0 do  { c++; w := intersect (w, w - 1) }

For every set 1 one iteration. O(n)

### Logarithmic Bit Sum

* Use bitmasks 01010101, 00110011, 00001111 etc.
* For all masks i:
* Mask the input word w with mask i (w<sub>even</sub>)
* Right-shift w by 2<sup>i</sup> bits
* Mask again shifted w with mask i (w<sub>odd</sub>)
* w := (w<sub>even</sub> + w<sub>odd</sub>)

&theta;(log(n))

### Lookup (Fastest)

Pre-calculate population count for, say, 8bit words and look them up. For words a multiple of 8bits shift them accordingly, do n lookups, sum up the counts.