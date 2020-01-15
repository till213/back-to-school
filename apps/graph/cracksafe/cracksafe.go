package main

import (
	"fmt"
	"strings"
	"strconv"
)

var seen map[string]bool

// Returns a minimal string that contains the code of length n
// with k characters. There are k^n possible codes, e.g. for 
// k=2 and n=3 there are 8 possible codes in total:
//   000, 001, 010, 011, 100, 101, 110, 111.
//
// The safe opens as soon at is "sees" that the last three entered
// digits match.
//
// The trick here is to generate k^(n-1) nodes, each node having
// k outgoing (directed) edges. We call "a node plus and edge"
// a "complete edge", which represents a given code.
//
// In the above example (k=2, n=3) there are 4 nodes:
//   00, 01, 10, 11
// Each node has two outgoing edges: 0 and 1. So for "starting"
// node 00 the "complete edges" are 000 and 001, which both
// represent two possible codes.
//
// Any directed (strongly connected) graph where each node has
// equal in-degree and out-degree has an Euler Circle. That Euler
// Circle ("visit each edge exactly once") happens to be our
// minimal code string.
// With the given graph we can find our Euler Circle with
// a simple Depth First Search: by tracking the visited "complete
// edges" (instead of nodes) we basically follow a "Hierholzer Walk".
// We just have to make sure to append our current code ("edge")
// in a post-order fashion.
//
// https://de.wikipedia.org/wiki/Algorithmus_von_Hierholzer
func crackSafe(n, k int) string {
	var ans strings.Builder
	
	if n == 1 && k == 1 {
		return "0"
	}

	seen = make(map[string]bool)
	start := strings.Repeat("0", n-1)
	dfs(start, k, &ans)
	// "Post-order" visit of the start node
	ans.WriteString(start)
	return ans.String()
}

func dfs(node string, k int, ans *strings.Builder) {
	for x := 0; x < k; x++ {
		completeEdge := fmt.Sprintf("%s%d", node, x)
		if !seen[completeEdge] {
			seen[completeEdge] = true
			dfs(completeEdge[1:], k, ans)
			// "Post-order" visit
			ans.WriteString(strconv.Itoa(x))
		}
	}
}

func main() {
	k := 2
	n := 2
	code := crackSafe(n, k)
	fmt.Printf("Code for k=%d and n=%d: %s\n", k, n, code)

	k = 2
	n = 3
	code = crackSafe(n, k)
	fmt.Printf("Code for k=%d and n=%d: %s\n", k, n, code)
}
