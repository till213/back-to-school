package main

import (
	"math/rand"
	"fmt"
	 "strconv"
	 "os"
	 "errors"
)

func partition(arr []int, p, r int) int {
	x := arr[p]
	i := p
	j := r
	for {
		for arr[j] > x {
			j--
		}
		for arr[i] < x {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			return j
		}
	}
}

func random(min, max int) int {
	return int(rand.Float64() * (float64(max) - float64(min) + 1) + (float64(min)))
}

func randomisedPartition(arr []int, p, r int) int {
	i := random(p, r)
	arr[p] , arr[i] = arr[i] , arr[p]
	return partition(arr, p, r)
}

// QuickSelect selects the i-th smallest element from arr,
// searching in the range [p, r]
// Randomized-Select, Introduction to Algorithms, ch. 10.2
func QuickSelect(arr []int, p, r, i int) int {
	if p == r {
		return arr[p]
	}
	q := randomisedPartition(arr, p, r)
	k := q - p + 1
	if i <= k {
		return QuickSelect(arr, p, q, i)
	} else {
		return QuickSelect(arr, q + 1, r, i - k)
	}
}

func main() {
	arr := []int{5, 2, 6, 7, 3}
	var n int
	var err error
	if len(os.Args) > 1 {
		n, err = strconv.Atoi(os.Args[1])
	} else {
		err = errors.New("No arg")
	}
	if err == nil {
		if n > len(arr) {
			err = errors.New("Argument too large")
			fmt.Println("Parse error: ", err)
		} else {
			r := QuickSelect(arr, 0, len(arr) - 1, n)
			fmt.Printf("%v-th smallest element: %v\n", n, r)
		}
	} else {
		fmt.Println("Parse error: ", err)
	}
}
